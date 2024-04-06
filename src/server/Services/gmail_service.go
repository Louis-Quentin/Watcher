package Services

import (
	action "Area/Action"
	"Area/Models"
	_ "Area/Models"
	reaction "Area/Reaction"
	"context"
	_ "encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	gmail "google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func getClient(config *oauth2.Config, token *oauth2.Token) *http.Client {
	return config.Client(context.Background(), token)
}

func Initialize_gmail_service(r *gin.Context, token *oauth2.Token) *gmail.Service {
	//println("INITIALIZE GMAIL SERVICE")
	//fmt.Println(token.AccessToken)
	ctx := context.Background()
	b, err := os.ReadFile("./Keys/credentials.json")
	if err != nil {
		log.Printf("Unable to read client secret file: %v", err)
	}
	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b,
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
		"https://mail.google.com/",
		"https://www.googleapis.com/auth/gmail.addons.current.action.compose",
		"https://www.googleapis.com/auth/gmail.addons.current.message.action",
		"https://www.googleapis.com/auth/gmail.addons.current.message.metadata",
		"https://www.googleapis.com/auth/gmail.compose",
		"https://www.googleapis.com/auth/gmail.addons.current.message.readonly",
		"https://www.googleapis.com/auth/gmail.insert",
		"https://www.googleapis.com/auth/gmail.labels",
		"https://www.googleapis.com/auth/gmail.metadata",
		"https://www.googleapis.com/auth/gmail.modify",
		"https://www.googleapis.com/auth/gmail.readonly",
		"https://www.googleapis.com/auth/gmail.send",
		"https://www.googleapis.com/auth/gmail.settings.basic",
		"https://www.googleapis.com/auth/gmail.settings.sharing",
	)
	if err != nil {
		log.Printf("Unable to parse client secret file to config: %v", err)
	}
	//println("config", config.Scopes, config.Endpoint.AuthURL)
	client := getClient(config, token)
	//println("get client done", client)

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Printf("Unable to retrieve Gmail client: %v", err)
	}
	return srv
}

func findAreaIndexByName(user Models.User, areaName string) int {
    for _, service := range user.Services {
        for aIndex, area := range service.Areas {
            if area.Name == areaName {
                return aIndex
            }
        }
    }
    return -1
}

func run_reaction(mail []string, context *gin.Context, token *oauth2.Token) {
	println("size of mail = ", len(mail))
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	area_type := reg.ReplaceAllString(mail[0], "")
	println("area type = ", area_type)
	email_string, _ := context.Cookie("email")
	s := context.MustGet("services").(map[string]*Service)
	var user Models.User
	db := context.MustGet("gorm").(gorm.DB)
	db.Where(Models.User{Email: email_string}).Preload("Services.Areas").First(&user)
	//println("size of type == ", len(area_type))
	//println("len of discord == ", len("DISCORD"))
	if area_type == "DISCORD" && len(mail) >= 2 {
		//println("gmail - discord reaction")
		//println("area state = ", user.Services[0].Areas[5].Active)
		id := findAreaIndexByName(user, "Discord_1")
		if user.Services[0].Areas[id].Active == true {
			//println("discord service true")
			user.Services[0].Areas[5].FieldValue_1 = mail[1]
			db.Save(&user.Services)
			db.Save(&user.Services[0].Areas)
			s["Discord"].Areas["Discord_3"](context, token)
		}
	} else if area_type == "GMAIL" && len(mail) >= 4 {
		id := findAreaIndexByName(user, "Gmail_2")
		if user.Services[0].Areas[id].Active == true {
			user.Services[0].Areas[1].FieldValue_1 = mail[1]
			user.Services[0].Areas[1].FieldValue_2 = mail[2]
			user.Services[0].Areas[1].FieldValue_3 = mail[3]
			db.Save(&user.Services)
			db.Save(&user.Services[0].Areas)
			s["Gmail"].Areas["Gmail_2"](context, token)
		}
	} else if area_type == "GDOC" && len(mail) >= 2 {
		id := findAreaIndexByName(user, "Gdoc_1")
		if user.Services[0].Areas[id].Active == true {
			user.Services[3].Areas[0].FieldValue_1 = mail[1]
			db.Save(&user.Services)
			db.Save(&user.Services[3].Areas)
			s["Gdoc"].Areas["Gdoc_1"](context, token)
		}
	} else if area_type == "GFORM" && len(mail) >= 3 {
		id := findAreaIndexByName(user, "Gform_1")
		if user.Services[0].Areas[id].Active == true {
			user.Services[6].Areas[0].FieldValue_1 = mail[1]
			user.Services[6].Areas[0].FieldValue_2 = mail[2]
			db.Save(&user.Services)
			db.Save(&user.Services[6].Areas)
			s["Gform"].Areas["Gform_1"](context, token)
		}
	} else if area_type == "CALENDAR" && len(mail) >= 3 {
		id := findAreaIndexByName(user, "Calendar_2")
		if user.Services[0].Areas[id].Active == true {
			user.Services[1].Areas[0].FieldValue_1 = mail[1]
			user.Services[1].Areas[0].FieldValue_2 = mail[2]
			db.Save(&user.Services)
			db.Save(&user.Services[1].Areas)
			s["Calendar"].Areas["Calendar_2"](context, token)
		}
	} else if area_type == "GDRIVE" && len(mail) >= 2 {
		id := findAreaIndexByName(user, "Gdrive_2")
		if user.Services[0].Areas[id].Active == true {
			user.Services[4].Areas[0].FieldValue_1 = mail[1]
			db.Save(&user.Services)
			db.Save(&user.Services[4].Areas)
			s["Gdrive"].Areas["Gdrive_2"](context, token)
		}
	} else if area_type == "GSHEET" && len(mail) >= 2 {
		id := findAreaIndexByName(user, "Gsheet_1")
		if user.Services[0].Areas[id].Active == true {
			user.Services[5].Areas[0].FieldValue_1 = mail[1]
			db.Save(&user.Services)
			db.Save(&user.Services[5].Areas)
			s["Gsheet"].Areas["Gsheet_1"](context, token)
		}
	} else if area_type == "GSLIDES" && len(mail) >= 2 {
		id := findAreaIndexByName(user, "Gslides_1")
		if user.Services[0].Areas[id].Active == true {
			user.Services[8].Areas[0].FieldValue_1 = mail[1]
			db.Save(&user.Services)
			db.Save(&user.Services[8].Areas)
			s["Gslides"].Areas["Gslides_1"](context, token)
		}
	} else if area_type == "YOUTUBE" && len(mail) >= 2 {
		id := findAreaIndexByName(user, "Youtube_4")
		if user.Services[0].Areas[id].Active == true {
			user.Services[2].Areas[3].FieldValue_1 = mail[1]
			db.Save(&user.Services)
			db.Save(&user.Services[2].Areas)
			s["Youtube"].Areas["Youtube_4"](context, token)
		}
	} else if area_type == "YOUTUBELIKE" && len(mail) >= 2 {
		println("first if")
		id := findAreaIndexByName(user, "Youtube_5")
		if user.Services[0].Areas[id].Active == true {
			user.Services[2].Areas[4].FieldValue_1 = mail[1]
			db.Save(&user.Services)
			db.Save(&user.Services[2].Areas)
			s["Youtube"].Areas["Youtube_5"](context, token)
		}
	} else if area_type == "YOUTUBEUSER" && len(mail) >= 2 {
		// println("yt area state = ", user.Services[0].Areas[11].Active)
		id := findAreaIndexByName(user, "Youtube_1")
		if user.Services[0].Areas[id].Active == true {
			println("111111111111111111111111111111111111")
			user.Services[2].Areas[0].FieldValue_1 = mail[1]
			db.Save(&user.Services)
			db.Save(&user.Services[2].Areas)
			println("going to yt reaction")
			s["Youtube"].Areas["Youtube_1"](context, token)
		}
	} else if area_type == "YOUTUBEVIDEO" && len(mail) >= 2 {
		id := findAreaIndexByName(user, "Youtube_2")
		if user.Services[0].Areas[id].Active == true {
			println("2222222222222222222222222222222222222222222222")
			user.Services[2].Areas[1].FieldValue_1 = mail[1]
			db.Save(&user.Services)
			db.Save(&user.Services[2].Areas)
			s["Youtube"].Areas["Youtube_2"](context, token)
		}
	} else if area_type == "YOUTUBEVIEW" && len(mail) >= 2 {
		id := findAreaIndexByName(user, "Youtube_3")
		if user.Services[0].Areas[id].Active == true {
				println("333333333333333333333333333333333333")
			user.Services[2].Areas[2].FieldValue_1 = mail[1]
			db.Save(&user.Services)
			db.Save(&user.Services[2].Areas)
			s["Youtube"].Areas["Youtube_3"](context, token)
		}
	} else {
		println("no area found to run")
		return
	}
}
func Gmail_1_area(r *gin.Context, token *oauth2.Token) {
	//println("ON EST DANS L'EXECUTION DE L'ACTION DE L'AREA !!")

	gmail_service := Initialize_gmail_service(r, token)
	mail := action.Get_last_mail(gmail_service, r)
	if mail == nil {
		return
	}
	//println("mail get")
	//println(mail.Header)
	//println(mail.Id)
	//println(mail.InternalDate)
	//println(mail.HistoryId)
	cleaned_mail := action.Parse_mail(mail)
	if cleaned_mail == "" {
		println("empty body -> no area to check")
		return
	}
	parsed_mail := strings.Split(cleaned_mail, "\n")
	/*for _, part := range parsed_mail {
		fmt.Printf("%s\n", part)
	}*/
	run_reaction(parsed_mail, r, token)
}

func Init_gmail_service_object(r *gin.Context) *Service {
	s := new(Service)
	s.Areas = map[string]func(r *gin.Context, token *oauth2.Token){
		"Gmail_1":    Gmail_1_area,
		"Gmail_2":    reaction.Gmail_2_area,
		"Gdoc_1":     Gdoc_1_area,
		"Gform_1":    Gform_1_area,
		"Calendar_2": Calendar_2_area,
		"Discord_3":  reaction.Discord_3_area,
		"Gdrive_2":   Gdrive_2_area,
		"Gsheet_1":   Gsheet_1_area,
		"Gslides_1":  Gslides_1_area,
		"Youtube_4":  Youtube_4_area,
		"Youtube_5":  Youtube_5_area,
		"Youtube_1":  Youtube_1_area,
		"Youtube_2":  Youtube_2_area,
		"Youtube_3":  Youtube_3_area,
	}
	return s
}
