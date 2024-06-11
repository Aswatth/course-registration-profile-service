package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateAuthorization(context *gin.Context) {
	//Get authorization token from cookies
	token_string, err := context.Cookie("Authorization")

	current_user := strings.Split(context.Request.URL.Path, "/")[1]

	if err != nil {
		context.AbortWithStatus(http.StatusUnauthorized)
	}

	// Parse the token
	token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		context.AbortWithStatus(http.StatusUnauthorized)
	}

	// Read parsed token
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		//Check for expired token
		if float64(time.Now().Unix()) > claims["expiry"].(float64) {
			context.AbortWithStatus(http.StatusUnauthorized)
		} else {
			//Check if user type is admin
			if claims["user_type"] == current_user {
				context.Next()
			} else {
				context.AbortWithStatus(http.StatusUnauthorized)
			}
		}
	} else {
		context.AbortWithStatus(http.StatusUnauthorized)
	}

}
