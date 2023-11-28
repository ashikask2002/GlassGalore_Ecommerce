package middleware

import (
	"GlassGalore/pkg/helper"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AdminAuthMiddleware(c *gin.Context) {

	accessToken := c.Request.Header.Get("Authorization")
	fmt.Println("access_token", accessToken)
	_, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("1234"), nil
	})

	if err != nil {
		fmt.Println("it happpens here")
		// The access token is invalid.
		c.AbortWithStatusJSON(401, gin.H{
			"message": "token error",
			"err":     err.Error(),
		})
		return
	}

	c.Next()
}

func CreateNewAccessTokenAdmin() (string, error) {
	claims := &helper.AuthcustomClaims{
		Role: "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newAccessToken, err := token.SignedString([]byte("1234"))
	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}
