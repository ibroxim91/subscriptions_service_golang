package handlers

import (
	"net/http"
	"strconv"
	"time"

	"subscriptions_service_golang/internal/models"
	"subscriptions_service_golang/internal/services"
	"subscriptions_service_golang/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SubscriptionHandler struct {
	service services.SubscriptionService
}

func NewSubscriptionHandler(service services.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{service: service}
}

// CreateSubscription godoc
// @Summary Create a new subscription
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription body models.Subscription true "Subscription object"
// @Success 201 {object} models.Subscription
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscriptions [post]
func (h *SubscriptionHandler) Create(c *gin.Context) {

	var req models.Subscription
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Log.Error("Failed to bind JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sub, err := h.service.Create(req)
	if err != nil {
		logger.Log.Error("Failed to create subscription", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, sub)
}

// GetSubscriptionByID godoc
// @Summary Get subscription by ID
// @Tags subscriptions
// @Produce json
// @Param id path int true "Subscription ID"
// @Success 200 {object} models.Subscription
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /subscriptions/{id} [get]
func (h *SubscriptionHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Log.Error("Failed to parse id", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	sub, err := h.service.GetByID(uint(id))
	if err != nil {
		logger.Log.Error("Failed to get subscription", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sub)
}

// ListSubscriptions godoc
// @Summary List subscriptions with optional filters
// @Tags subscriptions
// @Produce json
// @Param user_id query string false "Filter by user ID"
// @Param service_name query string false "Filter by service name"
// @Param from query string false "Filter from date (YYYY-MM-DD)"
// @Param to query string false "Filter to date (YYYY-MM-DD)"
// @Success 200 {array} models.Subscription
// @Failure 500 {object} models.ErrorResponse
// @Router /subscriptions [get]
func (h *SubscriptionHandler) List(c *gin.Context) {
	userID := c.Query("user_id")
	serviceName := c.Query("service_name")
	logger.Log.Info("Failed to bind JSON")
	var fromTime, toTime *time.Time
	if from := c.Query("from"); from != "" {
		t, err := time.Parse("2006-01-02", from)
		if err == nil {
			fromTime = &t
		}
	}
	if to := c.Query("to"); to != "" {
		t, err := time.Parse("2006-01-02", to)
		if err == nil {
			toTime = &t
		}
	}

	subs, err := h.service.List(userID, serviceName, fromTime, toTime)
	if err != nil {
		logger.Log.Error("Failed to list subscriptions", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subs)
}

// UpdateSubscription godoc
// @Summary Update subscription by ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path int true "Subscription ID"
// @Param subscription body models.Subscription true "Updated subscription object"
// @Success 200 {object} models.Subscription
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /subscriptions/{id} [put]
func (h *SubscriptionHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Log.Error("Failed to parse id", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req models.Subscription
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Log.Error("Failed to bind JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.ID = uint(id)

	sub, err := h.service.Update(req)
	if err != nil {
		logger.Log.Error("Failed to update subscription", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sub)
}

// DeleteSubscription godoc
// @Summary Delete subscription by ID
// @Tags subscriptions
// @Produce json
// @Param id path int true "Subscription ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /subscriptions/{id} [delete]
func (h *SubscriptionHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Log.Error("Failed to parse id", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.service.Delete(uint(id)); err != nil {
		logger.Log.Error("Failed to delete subscription", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

// GetTotalPrice godoc
// @Summary Calculate total price of subscriptions
// @Tags subscriptions
// @Produce json
// @Param user_id query string false "Filter by user ID"
// @Param service_name query string false "Filter by service name"
// @Param from query string false "Filter from date (YYYY-MM-DD)"
// @Param to query string false "Filter to date (YYYY-MM-DD)"
// @Success 200 {object} map[string]int
// @Failure 500 {object} models.ErrorResponse
// @Router /subscriptions/total [get]
func (h *SubscriptionHandler) TotalPrice(c *gin.Context) {
	userID := c.Query("user_id")
	serviceName := c.Query("service_name")

	var fromTime, toTime *time.Time
	if from := c.Query("from"); from != "" {
		t, err := time.Parse("2006-01-02", from)
		if err == nil {
			fromTime = &t
		}
	}
	if to := c.Query("to"); to != "" {
		t, err := time.Parse("2006-01-02", to)
		if err == nil {
			toTime = &t
		}
	}

	total, err := h.service.TotalPrice(userID, serviceName, fromTime, toTime)
	if err != nil {
		logger.Log.Error("Failed to calculate total price", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total_price": total})
}
