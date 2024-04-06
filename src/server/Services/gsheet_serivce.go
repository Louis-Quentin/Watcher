package Services

import (
	"Area/Models"
	"context"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/gmail/v1"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	sheets "google.golang.org/api/sheets/v4"
)

func Initialize_gsheet_service(r *gin.Context, token *oauth2.Token) *sheets.Service {
	ctx := context.Background()
	b, err := os.ReadFile("./Keys/credentials.json")
	if err != nil {
		log.Printf("Unable to read client secret file: %v", err)
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
	}

	config, err := google.ConfigFromJSON(b, scopes...)
	if err != nil {
		log.Printf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config, token)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Printf("Unable to create google sheets service: %v", err)
	}
	return srv
}

func Create_sheet(srv *sheets.Service, r *gin.Context, token *oauth2.Token) *sheets.Spreadsheet {
	email_string, err := r.Cookie("email")

	if err != nil {
		r.Redirect(http.StatusTemporaryRedirect, "/google_login")
	}
	var user Models.User
	db := r.MustGet("gorm").(gorm.DB)
	db.Where(Models.User{Email: email_string}).Preload("Services.Areas").First(&user)
	infos := user.Services[5].Areas[0].FieldValue_1

	sheet := &sheets.Spreadsheet{
		Properties: &sheets.SpreadsheetProperties{
			Title: infos,
		},
	}
	sheet, err2 := srv.Spreadsheets.Create(sheet).Do()
	if err2 != nil {
		log.Printf("Failed to create sheet: %v", err)
	}
	return sheet
}

func Gsheet_1_area(r *gin.Context, token *oauth2.Token) {
	println("ON EST DANS L'EXECUTION DE L'ACTION DE L'AREA !!")

	gsheet_service := Initialize_gsheet_service(r, token)
	Create_sheet(gsheet_service, r, token)
}

func Init_gsheet_service_object(r *gin.Context) *Service {
	s := new(Service)
	s.Areas = map[string]func(r *gin.Context, token *oauth2.Token){
		"Gsheet_1": Gsheet_1_area,
	}
	return s
}
