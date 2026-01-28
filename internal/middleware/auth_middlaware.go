package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

// AuthMiddleware checks for Bearer token in Authorization header
func AuthMiddleware(required bool) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")

        // If token is required
        if required {
            if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
                c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
                c.Abort()
                return
            }

            token := strings.TrimPrefix(authHeader, "Bearer ")
            // Test token
            if token != "test-token" {
                c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
                c.Abort()
                return
            }
        }

        // If token is optional or valid, continue
        c.Next()
    }
}
