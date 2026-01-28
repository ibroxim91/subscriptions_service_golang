package handlers

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func TestLoginHandler(t *testing.T) {
    gin.SetMode(gin.TestMode)

    router := gin.Default()
    authHandler := NewAuthHandler()
    router.POST("/login", authHandler.Login)

    t.Run("valid credentials", func(t *testing.T) {
        body := LoginRequest{Username: "admin", Password: "password"}
        jsonBody, _ := json.Marshal(body)

        req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
        req.Header.Set("Content-Type", "application/json")
        w := httptest.NewRecorder()

        router.ServeHTTP(w, req)

        assert.Equal(t, http.StatusOK, w.Code)

        var resp TokenResponse
        err := json.Unmarshal(w.Body.Bytes(), &resp)
        assert.NoError(t, err)
        assert.Equal(t, "test-token", resp.Token)
    })

    t.Run("invalid credentials", func(t *testing.T) {
        body := LoginRequest{Username: "admin", Password: "wrong"}
        jsonBody, _ := json.Marshal(body)

        req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
        req.Header.Set("Content-Type", "application/json")
        w := httptest.NewRecorder()

        router.ServeHTTP(w, req)

        assert.Equal(t, http.StatusUnauthorized, w.Code)

        var resp ErrorResponse
        err := json.Unmarshal(w.Body.Bytes(), &resp)
        assert.NoError(t, err)
        assert.Equal(t, "invalid credentials", resp.Error)
    })
}
