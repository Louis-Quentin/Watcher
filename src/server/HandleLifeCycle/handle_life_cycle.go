package HandleLifeCycle

import (
	"Area/Controllers"
	"Area/Models"
	Service "Area/Services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/gmail/v1"
	"gorm.io/gorm"
	"log"
	"time"
)

func run_life_cycle(c *gin.Context) {
	//println("Run life cycle")
	s := c.MustGet("services").(map[string]*Service.Service)
	db := c.MustGet("gorm").(gorm.DB)
	var users []Models.User
	db.Find(&users).Preload("Services.Areas").Find(&users)
	//fmt.Println("NB users found in db: ", len(users))

	for i, elem := range users {
		fmt.Printf("user %d: using email: %s\n", i, elem.Email)
		token := oauth2.Token{
			AccessToken: elem.Google_access_token,
		}
		if elem.Google_access_token == "" {
			token = Controllers.Refresh_token(elem.Google_refresh_token,
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
				"https://www.googleapis.com/auth/youtubepartner-channel-audit")
			if token.AccessToken == "" {
				fmt.Printf("Can't connect this user to google\n")
				continue
			}
		}
		//println("refresh token ok, now check available services")
		for _, service := range elem.Services {
			if service.Active == true {
				//println("service found = ", service.Name)
				for _, area := range service.Areas {
					if area.Active == true && area.Action == true {
						//println("area found = ", area.Name)
						//call area
						//println("calling area func")
						s[service.Name].Areas[area.Name](c, &token)
					}
				}
			}
		}
	}
}

func Set_life_cycle(c *gin.Context) {
	life_cycle := gocron.NewScheduler(time.UTC)
	_, err := life_cycle.Every(3).Seconds().Do(run_life_cycle, c)
	if err != nil {
		log.Fatalf("could not start life cycle: %v", err)
		return
	}
	life_cycle.StartBlocking()
}
