package Reaction

import (
	"Area/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
	"net/http"
	"net/smtp"
)

func Send_mail(r *gin.Context, token *oauth2.Token) {
	println("in send mail before cookie")
	email_string, err := r.Cookie("email")

	if err != nil {
		r.Redirect(http.StatusTemporaryRedirect, "/google_login")
		return
	}
	println("cookie get")
	var user Models.User
	db := r.MustGet("gorm").(gorm.DB)
	db.Where(Models.User{Email: email_string}).Preload("Services.Areas").First(&user)
	To := user.Services[0].Areas[1].FieldValue_1
	title := user.Services[0].Areas[1].FieldValue_2
	body := user.Services[0].Areas[1].FieldValue_3
	println("in send mail to = ", To)
	println("in send mail title = ", title)
	println("in send mail body = ", body)

	auth := smtp.PlainAuth("", "area.epitech.2@gmail.com", "tcpfnzufiluqxgia", "smtp.gmail.com")

	to := []string{To}

	msg := "Subject: " + title + "\n" + body

	err2 := smtp.SendMail("smtp.gmail.com:587", auth, "area.epitech.2@gmail.com", to, []byte(msg))

	if err2 != nil {
		fmt.Printf("%v", err)
	}

}

func Gmail_2_area(r *gin.Context, token *oauth2.Token) {
	println("in send mail func")
	Send_mail(r, token)
}
