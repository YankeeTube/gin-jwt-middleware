package authenticate

import (
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TokenAuthMiddleware(c *gin.Context) {
	privateKey := c.MustGet("PRIVKEY")
	token, err := c.Request.Cookie("access-token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  "Authentication failed",
		})
		c.Abort()
		return
	}

	tokenString := token.Value
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Token is None",
		})
		c.Abort()
		return
	}

	claims := &jwt.MapClaims{}

	_, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})
	fmt.Println(err)

	if err == jwt.ErrSignatureInvalid {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Invalid Signature",
		})
		c.Abort()
	} else if v, _ := err.(*jwt.ValidationError); v != nil && v.Errors == jwt.ValidationErrorExpired {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Token is Expired",
		})
		c.Abort()
	} else if err != nil && err.Error() == rsa.ErrVerification.Error() {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Verification failed",
		})
		c.Abort()
	} else if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Authentication failed",
		})
		c.Abort()
	} else {
		c.Next()
	}
}
