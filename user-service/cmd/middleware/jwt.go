package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"strings"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(g *gin.Context) {
		header := g.GetHeader("Authorization")
		if header == "" {
			g.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			g.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(header, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			g.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			g.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		g.Set("user_id", claims["user_id"])
		g.Set("role", claims["role"])

		g.Next()
	}
}
