package jwt

import (
	"net/http"
	"strings"

	"aioc/internal/api/response"
	"aioc/pkg/e"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const TOKEN_TIMEOUT = 60 * 10

var JwtSecret = []byte("infrawaves")

type Claims struct {
	Id       int    `json:"id"`
	Username string `json:"name"`
	Email    string `json:"email"`
	Auth     bool   `json:"auth"`
	jwt.StandardClaims
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			response.Resp(c, http.StatusUnauthorized, e.INVALID_TOKEN, nil)
			c.Abort()
			return
		}
		token := strings.Split(authorization, " ")
		if len(token) < 2 {
			response.Resp(c, http.StatusUnauthorized, e.INVALID_TOKEN, nil)
			c.Abort()
			return
		}
		info, err := ParseToken(token[1])
		if err != nil {
			var code int
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT

			default:
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			}
			response.Resp(c, http.StatusUnauthorized, code, nil)
			c.Abort()
			return
		}
		c.Set("account", info)
		c.Next()
	}
}

func ParseToken(token string) (claims *Claims, err error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})

	if tokenClaims != nil {
		var ok bool
		claims, ok = tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return
}
