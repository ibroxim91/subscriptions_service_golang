package handlers

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"
	"subscriptions_service_golang/pkg/logger"
    "subscriptions_service_golang/internal/models"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

type FakeSubscriptionService struct{}

func (s *FakeSubscriptionService) Create(sub models.Subscription) (*models.Subscription, error) {
    sub.ID = 1
    return &sub, nil
}
func (s *FakeSubscriptionService) GetByID(id uint) (*models.Subscription, error) {
    return &models.Subscription{ID: id, ServiceName: "Netflix", Price: 10000}, nil
}
func (s *FakeSubscriptionService) List(userID, serviceName string, from, to *time.Time) ([]models.Subscription, error) {
    return []models.Subscription{
        {ID: 1, ServiceName: "Netflix", Price: 10000},
        {ID: 2, ServiceName: "Spotify", Price: 5000},
    }, nil
}
func (s *FakeSubscriptionService) Update(sub models.Subscription) (*models.Subscription, error) {
    sub.ServiceName = "Updated"
    return &sub, nil
}
func (s *FakeSubscriptionService) Delete(id uint) error {
    return nil
}
func (s *FakeSubscriptionService) TotalPrice(userID, serviceName string, from, to *time.Time) (int, error) {
    return 15000, nil
}



func setupRouter() *gin.Engine {
    gin.SetMode(gin.TestMode)
    r := gin.Default()

    // logger init
    logger.Init()

    service := &FakeSubscriptionService{}
    handler := NewSubscriptionHandler(service)

    r.POST("/subscriptions", handler.Create)
    r.GET("/subscriptions/:id", handler.GetByID)
    r.GET("/subscriptions", handler.List)
    r.PUT("/subscriptions/:id", handler.Update)
    r.DELETE("/subscriptions/:id", handler.Delete)
    r.GET("/subscriptions/total", handler.TotalPrice)

    return r
}


func TestCreateSubscription(t *testing.T) {
    r := setupRouter()

    body := models.Subscription{ServiceName: "Netflix", Price: 10000}
    jsonBody, _ := json.Marshal(body)

    req, _ := http.NewRequest("POST", "/subscriptions", bytes.NewBuffer(jsonBody))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusCreated, w.Code)
    var resp models.Subscription
    json.Unmarshal(w.Body.Bytes(), &resp)
    assert.Equal(t, "Netflix", resp.ServiceName)
}

func TestGetByID(t *testing.T) {
    r := setupRouter()

    req, _ := http.NewRequest("GET", "/subscriptions/1", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    var resp models.Subscription
    json.Unmarshal(w.Body.Bytes(), &resp)
    assert.Equal(t, uint(1), resp.ID)
}

func TestListSubscriptions(t *testing.T) {
    r := setupRouter()

    req, _ := http.NewRequest("GET", "/subscriptions", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    var resp []models.Subscription
    json.Unmarshal(w.Body.Bytes(), &resp)
    assert.Len(t, resp, 2)
}

func TestUpdateSubscription(t *testing.T) {
    r := setupRouter()

    body := models.Subscription{ServiceName: "Netflix", Price: 10000}
    jsonBody, _ := json.Marshal(body)

    req, _ := http.NewRequest("PUT", "/subscriptions/1", bytes.NewBuffer(jsonBody))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    var resp models.Subscription
    json.Unmarshal(w.Body.Bytes(), &resp)
    assert.Equal(t, "Updated", resp.ServiceName)
}

func TestDeleteSubscription(t *testing.T) {
    r := setupRouter()

    req, _ := http.NewRequest("DELETE", "/subscriptions/1", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
}

func TestTotalPrice(t *testing.T) {
    r := setupRouter()

    req, _ := http.NewRequest("GET", "/subscriptions/total", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    var resp map[string]int
    json.Unmarshal(w.Body.Bytes(), &resp)
    assert.Equal(t, 15000, resp["total_price"])
}
