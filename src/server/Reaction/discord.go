package Reaction

import (
	"Area/Models"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
	"net/http"
)

func SendDiscordMessage(r *gin.Context, token *oauth2.Token) {

	email_string, err := r.Cookie("email")

	if err != nil {
		r.Redirect(http.StatusTemporaryRedirect, "/google_login")
		return
	}

	var user Models.User
	db := r.MustGet("gorm").(gorm.DB)
	db.Where(Models.User{Email: email_string}).Preload("Services.Areas").First(&user)
	msg := user.Services[0].Areas[5].FieldValue_1
	println("field value = ", msg)

	message := struct {
		Content string `json:"content"`
	}{
		Content: msg,
	}

	messageJSON, err := json.Marshal(message)

	if err != nil {
		return
	}

	http.Post("https://discordapp.com/api/webhooks/1072470767524130846/1TuT9HUV5h_i8yuDgRNrx2LgrAQhCP_yGbCB2oDVxiIFkN_H1JludtU-3D_QmtV4rOnx", "application/json", bytes.NewBuffer(messageJSON))
}

func Discord_3_area(r *gin.Context, token *oauth2.Token) {
	SendDiscordMessage(r, token)
}
