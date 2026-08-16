package middleware

import (
	"Authentication/internaal/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	ctxUserIDKey = "auth.userId"
	ctxRoleKey   = "auth.role"
)

func AutRequired(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := strings.TrimSpace(c.GetHeader("Authorization"))
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Missing Authorization token",
			})
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid Authorization format",
			})
			return
		}
		scheme := strings.TrimSpace(parts[0])
		tokenString := strings.TrimSpace(parts[1])

		if !strings.EqualFold(scheme, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authentication Scheme must be Bearer",
			})
			return
		}

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Missing Token here",
			})
			return
		}
		claims, err := auth.ParseToken(jwtSecret, tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or Expired toten",
			})
			return
		}

		c.Set(ctxUserIDKey, claims.Subject)
		c.Set(ctxRoleKey, claims.Role)

		c.Next()
	}
}

func GetUserID(c *gin.Context) (string, bool) {
	res, ok := c.Get(ctxUserIDKey)
	if !ok {
		return "", false
	}
	userID, ok := res.(string)
	return userID, ok
}

func GetRole(c *gin.Context) (string, bool) {
	res, ok := c.Get(ctxRoleKey)
	if !ok {
		return " ", false
	}
	role, ok := res.(string)
	return role, ok
}
