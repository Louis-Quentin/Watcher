package Services

import (
	"Area/Models"
	"context"
	"google.golang.org/api/gmail/v1"
	"gorm.io/gorm"
	"net/http"

	// "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	calendar "google.golang.org/api/calendar/v3"
	option "google.golang.org/api/option"
	"log"
	"os"
	"time"
	// "io"
)

func Initialize_calendar_service(r *gin.Context, token *oauth2.Token) *calendar.Service {
	ctx := context.Background()
	b, err := os.ReadFile("./Keys/credentials.json")
	if err != nil {
		log.Printf("Unable to read client secret file: %v", err)
	}
	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/userinfo.profile",
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
		"https://www.googleapis.com/auth/youtubepartner-channel-audit")
	if err != nil {
		log.Printf("Unable to parse client secret file to config: %v", err)
	}
	println("config", config.Scopes, config.Endpoint.AuthURL)
	client := getClient(config, token)
	println("get client done", client)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))

	if err != nil {
		log.Printf("Unable to retrieve Calendar client: %v", err)
	}
	println("calendar service spawned")
	return srv
}

func Create_calendar_event(srv *calendar.Service, context *gin.Context, token *oauth2.Token) {
	email_string, err := context.Cookie("email")

	if err != nil {
		context.Redirect(http.StatusTemporaryRedirect, "/google_login")
		return
	}
	var user Models.User
	db := context.MustGet("gorm").(gorm.DB)
	db.Where(Models.User{Email: email_string}).Preload("Services.Areas").First(&user)
	summary := user.Services[1].Areas[0].FieldValue_1
	description := user.Services[1].Areas[0].FieldValue_2
	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1)

	event := &calendar.Event{
		Summary:     summary,
		Location:    "",
		Description: description,
		Start: &calendar.EventDateTime{
			DateTime: tomorrow.Format(time.RFC3339),
			TimeZone: "Europe/Paris",
		},
		End: &calendar.EventDateTime{
			DateTime: tomorrow.Add(1 * time.Hour).Format(time.RFC3339),
			TimeZone: "Europe/Paris",
		},
		Recurrence: []string{"RRULE:FREQ=DAILY;COUNT=1"},
		Attendees: []*calendar.EventAttendee{
			&calendar.EventAttendee{Email: email_string},
		},
	}
	calendarId := "primary"
	event, err2 := srv.Events.Insert(calendarId, event).Do()
	if err2 != nil {
		log.Fatalf("Unable to create event. %v\n", err2)
	}
}

func Get_last_event(srv *calendar.Service, r *gin.Context, token *oauth2.Token) *calendar.Event {
	events, err := srv.Events.List("primary").TimeMin(time.Now().Format(time.RFC3339)).Do()
	if err != nil {
		log.Printf("Unable to retreive any calendar event: %v", err)
	}

	var lastEvent *calendar.Event
	for _, event := range events.Items {
		if lastEvent == nil || event.Start.DateTime > lastEvent.Start.DateTime {
			lastEvent = event
		}
	}

	if lastEvent == nil {
		println("O event found")
	}
	if lastEvent.Description != "" {
		// log.Println("On stop le calendar")
		return nil
	}
	fmt.Println(lastEvent.Summary)
	lastEvent.Description = "Description"
	updateEvent, err := srv.Events.Update("primary", lastEvent.Id, lastEvent).Do()
	if err != nil {
		log.Printf("Unable to update calendar description: %v", err)
	}
	// log.Println("CALLLENNNDDARR")
	return updateEvent
}

func Calendar_1_area(r *gin.Context, token *oauth2.Token) {
	calendar_service := Initialize_calendar_service(r, token)
	Get_last_event(calendar_service, r, token)
}

func Calendar_2_area(r *gin.Context, token *oauth2.Token) {
	service := Initialize_calendar_service(r, token)
	Create_calendar_event(service, r, token)
}

func Init_calendar_service_object(r *gin.Context) *Service {
	s := new(Service)
	s.Areas = map[string]func(r *gin.Context, token *oauth2.Token){
		"Calendar_1": Calendar_1_area,
		"Calendar_2": Calendar_2_area,
	}
	return s
}
