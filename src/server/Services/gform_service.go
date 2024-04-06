package Services

import (
	"Area/Models"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	forms "google.golang.org/api/forms/v1"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"gorm.io/gorm"
)

func Initialize_gform_serice(r *gin.Context, token *oauth2.Token) *forms.Service {
	fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	ctx := context.Background()

	b, err := os.ReadFile("./Keys/credentials.json")
	if err != nil {
		log.Printf("Unable to read clien secret file: %v", err)
	}

	scopes := []string{
		"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email",
		gmail.MailGoogleComScope, gmail.GmailModifyScope, gmail.GmailComposeScope, gmail.GmailSendScope,
		calendar.CalendarEventsScope, calendar.CalendarScope,
		"https://www.googleapis.com/auth/drive",
		"https://www.googleapis.com/auth/drive.file",
		"https://www.googleapis.com/auth/drive.readonly",
		"https://www.googleapis.com/auth/drive.appdata",
		"https://www.googleapis.com/auth/drive.metadata",
		"https://www.googleapis.com/auth/spreadsheets",
		"https://www.googleapis.com/auth/spreadsheets.readonly",
		"https://www.googleapis.com/auth/script.scriptapp",
		"https://www.googleapis.com/auth/script.external_request",
		"https://www.googleapis.com/auth/script.send_mail",
		"https://www.googleapis.com/auth/documents",
		"https://www.googleapis.com/auth/presentations",
		"https://www.googleapis.com/auth/forms",
		"https://www.googleapis.com/auth/youtube",
		"https://www.googleapis.com/auth/youtube.force-ssl",
		"https://www.googleapis.com/auth/youtube.readonly",
		"https://www.googleapis.com/auth/youtubepartner",
		"https://www.googleapis.com/auth/youtubepartner-channel-audit",
		"https://www.googleapis.com/auth/forms.body",
	}

	config, err := google.ConfigFromJSON(b, scopes...)
	if err != nil {
		log.Printf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config, token)

	srv, err := forms.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Printf("Unable to retrieve Form client: %v", err)
	}
	println("gform service spawned")
	return srv
}

func Create_form(srv *forms.Service, r *gin.Context, token *oauth2.Token) *forms.Form {
	//fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	email_string, err := r.Cookie("email")

	if err != nil {
		r.Redirect(http.StatusTemporaryRedirect, "/google_login")
	}
	var user Models.User
	db := r.MustGet("gorm").(gorm.DB)
	db.Where(Models.User{Email: email_string}).Preload("Services.Areas").First(&user)
	title := user.Services[6].Areas[0].FieldValue_1
	//title = "YESSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSS"
	if title == "" {
		println("title or description are empty -> failed")
		return nil
	}
	initForm := &forms.Form{
		Info: &forms.Info{
			Title: title,
		},
	}

	form, err := srv.Forms.Create(initForm).Do()
	if err != nil {
		log.Printf("Failed to create form: %v", err)
	}
	fmt.Println("FORM created")
	return form
}

func Gform_1_area(r *gin.Context, token *oauth2.Token) {
	//fmt.Println("GGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGj")
	gform_service := Initialize_gform_serice(r, token)
	//fmt.Println("allo")
	Create_form(gform_service, r, token)
}

func Init_gform_service_object(r *gin.Context) *Service {
	s := new(Service)
	s.Areas = map[string]func(r *gin.Context, token *oauth2.Token){
		"Gform_1": Gform_1_area,
	}
	return s
}
