package handlers

import (
    "net/http"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"
    "subscriptions_service_golang/internal/models"
    "subscriptions_service_golang/internal/services"
)

type SubscriptionHandler struct {
    service services.SubscriptionService
}

func NewSubscriptionHandler(service services.SubscriptionService) *SubscriptionHandler {
    return &SubscriptionHandler{service: service}
}

// POST /subscriptions
func (h *SubscriptionHandler) Create(c *gin.Context) {
    var req models.Subscription
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    sub, err := h.service.Create(req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, sub)
}

// GET /subscriptions/:id
func (h *SubscriptionHandler) GetByID(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }
    sub, err := h.service.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, sub)
}

// GET /subscriptions
func (h *SubscriptionHandler) List(c *gin.Context) {
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

    subs, err := h.service.List(userID, serviceName, fromTime, toTime)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, subs)
}

// PUT /subscriptions/:id
func (h *SubscriptionHandler) Update(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    var req models.Subscription
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    req.ID = uint(id)

    sub, err := h.service.Update(req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, sub)
}

// DELETE /subscriptions/:id
func (h *SubscriptionHandler) Delete(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }
    if err := h.service.Delete(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

// GET /subscriptions/total
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
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"total_price": total})
}
