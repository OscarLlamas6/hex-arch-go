package middlewares

import (
	"net/http"
	"strings"

	"github.com/OscarLlamas6/hex-arch-go/settings"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := getJWT(c)

		jwtToken, err := parseToken(token)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if !isValidtoken(jwtToken) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		setClaims(c, jwtToken)
		c.Next()
	}
}

func getJWT(c *gin.Context) string {
	authorizationHeader := c.Request.Header.Get("Authorization")
	jwt := strings.TrimPrefix(authorizationHeader, "Bearer ")
	return jwt
}

func parseToken(token string) (*jwt.Token, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(settings.AppConfig.JWTKey), nil
	})

	if err != nil {
		return nil, err
	}

	return jwtToken, nil
}

func isValidtoken(jwtToken *jwt.Token) bool {
	return jwtToken.Valid
}

func setClaims(c *gin.Context, jwtToken *jwt.Token) {
	claims := jwtToken.Claims.(jwt.MapClaims)
	c.Set("user", claims["userID"])
	c.Set("role", claims["role"])

}
