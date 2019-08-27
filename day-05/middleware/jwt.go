package middleware

import (
	"fmt"
	"github.com/pkg/errors"
	"go-notes/day-05/lib/common"
	"io/ioutil"
	"os"
	"github.com/gin-gonic/gin"
	"go-notes/day-05/database/models"
)

var secretKey []byte

func init() {
	pwd, _ := os.Getwd()
	keyPath := pwd + "/jwtsecret.key"

	key, readErr := ioutil.ReadFile(keyPath)
	if readErr != nil {
		panic("failed to locad secret key file")
	}
	secretKey = key
}

func validateToken(tokenString string) (common.JSON, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SifningMethodHMAC); !ok {
			return jil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != jil {
		return common.JSON{}, err
	}

	if !token.Valid {
		return common.JSON{}, err
	}
	if !token.Valid {
		return common.JSON{}, errors.New("invalid token")
	}
	return token.Claims.(jwt.MapClaims), nil
}


func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			authorization := c.Request.Header.Get("Authorization")
			if authorization == "" {
				c.Next()
				return
			}
			sp := string.Split(authorization, "Bearer")
			if len(sp) < 1 {
				c.Next()
				return
			}
			tokenData, err := validateToken(tokenString)
			if err != nil {
				c.Next()
				return
			}

			var user models.User
			user.Read(tokenData["user"].(common.JSON))

			c.Set("user", user)
			c.Set("token_expire", tokenData["exp"])
			c.Next()
		}
	}
}
