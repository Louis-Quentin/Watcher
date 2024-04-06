package Middleware

import (
	"Area/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
)

/*func Require_google_auth(r *gin.Context) {
	//faire comme l'auth jwt
	//check si c vide, si  oui se login, sinon refresh le token et le stocker (on va stocker le refresh et le access)

	user := r.MustGet("user").(Models.User)
	if user.Google_access_token == "" {
		r.Redirect(http.StatusTemporaryRedirect, "http://localhost:8080/google_login")
		return
	}

	//utiliser le user get pour checker si le google token est valide sinon rediriger sur le google login
}*/

func Require_auth(r *gin.Context) {

	token_string, err := r.Cookie("Authorization")

	if err != nil {
		r.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			r.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var user Models.User
		db := r.MustGet("gorm").(gorm.DB)
		db.First(&user, claims["sub"])
		if user.ID == 0 {
			r.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		r.Set("user", user)
		r.Next()
	} else {
		r.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
