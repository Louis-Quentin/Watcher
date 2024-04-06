package Services

import (
	"Area/Models"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	drive "google.golang.org/api/drive/v3"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func Initialize_gdrive_service(r *gin.Context, token *oauth2.Token) *drive.Service {
	ctx := context.Background()

	b, err := os.ReadFile("./Keys/credentials.json")
	if err != nil {
		log.Printf("Unable to read client secret file: %v", err)
		return nil
	}

	scopes := []string{
		"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email",
		gmail.MailGoogleComScope, gmail.GmailModifyScope, gmail.GmailComposeScope, gmail.GmailSendScope,
		calendar.CalendarEventsScope, calendar.CalendarScope,
		"https://www.googleapis.com/auth/drive",
		"https://www.googleapis.com/auth/drive.appdata",
		"https://www.googleapis.com/auth/drive.metadata",
		"https://www.googleapis.com/auth/drive.metadata.readonly",
		"https://www.googleapis.com/auth/drive.photos.readonly",
		"https://www.googleapis.com/auth/drive.readonly",
		"https://www.googleapis.com/auth/drive.file",
		"https://www.googleapis.com/auth/drive.scripts",
		"https://www.googleapis.com/auth/drive.readonly",
		"https://www.googleapis.com/auth/drive.appdata",
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
	}

	config, err := google.ConfigFromJSON(b, scopes...)
	if err != nil {
		log.Printf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config, token)

	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Printf("Unable to retrieve Drive client: %v", err)
	}
	return srv
}

func Create_drive(srv *drive.Service, r *gin.Context, token *oauth2.Token) *drive.File {
	email_string, err := r.Cookie("email")

	if err != nil {
		r.Redirect(http.StatusTemporaryRedirect, "/google_login")
	}
	var user Models.User
	db := r.MustGet("gorm").(gorm.DB)
	db.Where(Models.User{Email: email_string}).Preload("Services.Areas").First(&user)
	infos := user.Services[4].Areas[0].FieldValue_1

	doc := &drive.File{
		Name:     infos,
		MimeType: "application/vnd.google-apps.document",
	}
	doc, err2 := srv.Files.Create(doc).Do()
	if err2 != nil {
		log.Printf("Failed to create file: %v", err)
	}
	return doc
}

func Get_last_drive(srv *drive.Service, context *gin.Context, token *oauth2.Token) *drive.File {
	r, err := srv.Files.List().OrderBy("modifiedTime desc").PageSize(1).Do()
	if err != nil {
		log.Printf("Unable to retrieve files: %v", err)
	}

	var lastFile *drive.File
	if len(r.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		lastFile = r.Files[0]
	}
	lastFileInfos, _ := srv.Files.Get(lastFile.Id).Fields("*").Do()

	if lastFileInfos.Description != "" {
		log.Println("On stop le drive")
		return nil
	} else {
		tmp := &drive.File{
			Description: "description",
		}
		updateFile, err := srv.Files.Update(lastFile.Id, tmp).Do()
		if err != nil {
			log.Printf("Unable to update drive description: %v", err)
		}
		log.Println("On à modifier le drive")
		return updateFile
	}
}

func Gdrive_1_area(r *gin.Context, token *oauth2.Token) {
	gdrive_service := Initialize_gdrive_service(r, token)
	drive := Get_last_drive(gdrive_service, r, token)
	if drive == nil {
		return
	}
	email_string, _ := r.Cookie("email")
	s := r.MustGet("services").(map[string]*Service)
	var user Models.User
	db := r.MustGet("gorm").(gorm.DB)
	db.Where(Models.User{Email: email_string}).Preload("Services.Areas").First(&user)
	if user.Services[4].Areas[1].Active == true {
		user.Services[0].Areas[1].FieldValue_1 = email_string
		user.Services[0].Areas[1].FieldValue_2 = "new google drive document received"
		user.Services[0].Areas[1].FieldValue_3 = "Let's check this new shared file !!"
		db.Save(&user.Services)
		db.Save(&user.Services[0].Areas)
		s["Gmail"].Areas["Gmail_2"](r, token)
	}
}

func Gdrive_2_area(r *gin.Context, token *oauth2.Token) {
	println("ON EST DANS L'EXECUTION DE gdrive 2!!")

	gdrive_service := Initialize_gdrive_service(r, token)
	Create_drive(gdrive_service, r, token)
}

func Init_gdrive_service_object(r *gin.Context) *Service {
	s := new(Service)
	s.Areas = map[string]func(r *gin.Context, token *oauth2.Token){
		"Gdrive_1": Gdrive_1_area,
		"Gdrive_2": Gdrive_2_area,
	}
	return s
}
