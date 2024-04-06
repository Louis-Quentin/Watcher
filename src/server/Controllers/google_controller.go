package Controllers

import (
	"Area/Models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/gmail/v1"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	_ "os"
	"time"
)

type GoogleProfile struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
}

var random_state = "random"

func Refresh_token(refresh_token string, scopes ...string) oauth2.Token {
	config := Get_config(scopes...)
	type Body struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int64  `json:"expires_in"`
	}
	err_token := oauth2.Token{}
	rf_url := url.Values{}
	rf_url.Add("client_id", config.ClientID)
	rf_url.Add("client_secret", config.ClientSecret)
	rf_url.Add("refresh_token", refresh_token)
	rf_url.Add("grant_type", "refresh_token")
	response, err := http.PostForm("https://www.googleapis.com/oauth2/v4/token", rf_url)
	if err != nil {
		log.Fatalf("Failed to refresh token: %v", err)
		return err_token
	}
	var res Body
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err = json.Unmarshal(content, &res); err != nil {
		log.Fatalf("failed to refresh token: %v", err)
		return err_token
	}
	token := oauth2.Token{
		AccessToken: res.AccessToken,
		Expiry:      time.Now().Add(time.Hour * 7),
		TokenType:   res.TokenType,
	}
	return token
}

func Get_config(scopes ...string) *oauth2.Config {
	b, err := os.ReadFile("./Keys/credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
		return nil
	}
	var googleOauthConfig = oauth2.Config{}
	config, err := google.ConfigFromJSON(b, scopes...)
	err = json.Unmarshal(b, &googleOauthConfig)
	if err != nil {
		log.Fatalf("Unable to parse client config file: %v", err)
		return nil
	}
	return config
}

func Set_Login_location() string {
	config := Get_config("https://www.googleapis.com/auth/userinfo.profile",
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
	log_url := config.AuthCodeURL(random_state)
	var rf_access string = "prompt=consent&access_type=offline&"
	return log_url[:len(log_url)-12] + rf_access + log_url[len(log_url)-12:]
}

func Google_login(r *gin.Context) {
	r.Redirect(http.StatusTemporaryRedirect, Set_Login_location())
	//println("REDIRECT DONE")
}

func Google_callback(r *gin.Context) {
	config := Get_config("https://www.googleapis.com/auth/userinfo.profile",
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
	if r.Request.FormValue("state") != random_state {
		r.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	token, err := config.Exchange(oauth2.NoContext, r.Request.FormValue("code"))
	if err != nil {
		fmt.Printf("Could not get: %s\n", err.Error())
		r.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	//test token
	res, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Printf("Could not create get request: %s\n", err.Error())
		r.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Could not parse response: %s\n", err.Error())
		r.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	var parsed_content = GoogleProfile{}
	err = json.Unmarshal(content, &parsed_content)
	if err != nil {
		fmt.Println("Could not parse response")
		r.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	r.SetCookie("email", parsed_content.Email, 3600*24*30, "", "", false, false)

	db := r.MustGet("gorm").(gorm.DB)
	var user Models.User
	result := db.Where(Models.User{Email: parsed_content.Email}).First(&user)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		println("not found")
	}
	if result.RowsAffected > 0 {
		user.Google_refresh_token = token.RefreshToken
		result = db.Model(&user).Updates(&user)
		if result.Error != nil {
			println("update failed")
		}
	} else {
		db.FirstOrCreate(Build_user(token.RefreshToken, parsed_content.Email, "", ""), Models.User{Email: parsed_content.Email})
	}
	//set token or build a new user using google token
	/*db := r.MustGet("gorm").(gorm.DB)
	if err = db.Where(Models.User{Email: parsed_content.Email}).
		Assign(Models.User{Email: parsed_content.Email, Google_refresh_token: token.RefreshToken, Google_access_token: "", Password: ""}).
		FirstOrCreate(Build_user(token.RefreshToken, parsed_content.Email, "", "")).Error; err != nil {
		r.Next()
		return
	}*/
	q := url.Values{}
	q.Set("email", parsed_content.Email)
	location := url.URL{Path: "http://localhost:8081/explore", RawQuery: q.Encode()}
	r.Redirect(http.StatusTemporaryRedirect, location.RequestURI())
}

func Build_user(token string, email string, ac_token string, password string) *Models.User {

	var gmail_areas = []Models.AreaState{
		Models.AreaState{
			Active:        false,
			Name:          "Gmail_1",
			Action:        true,
			Field:         0,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
			PrintName:     "surveiller un mail",
		},
		Models.AreaState{
			Active:        false,
			Name:          "Gmail_2",
			Action:        false,
			Field:         0,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
			PrintName:     "envoyer un mail",
		},
		Models.AreaState{
			Active:        false,
			Name:          "Gdoc_1",
			Action:        false,
			Field:         0,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
			PrintName:     "créer un google doc",
		},
		Models.AreaState{
			Active:        false,
			Name:          "Gform_1",
			Action:        false,
			Field:         0,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
			PrintName:     "créer un google form",
		},
		Models.AreaState{
			Active:        false,
			Name:          "Calendar_2",
			Action:        false,
			Field:         0,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
			PrintName:     "créer un google calendar",
		},
		Models.AreaState{
			Active:        false,
			Name:          "Discord_3",
			Action:        false,
			Field:         0,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
			PrintName:     "envoyer un msg discord",
		},
		Models.AreaState{
			Active:        false,
			Name:          "Gdrive_2",
			Action:        false,
			Field:         0,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
			PrintName:     "créer un google drive",
		},
		Models.AreaState{
			Active:        false,
			Name:          "Gsheet_1",
			Action:        false,
			Field:         0,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
			PrintName:     "créer un google sheet",
		},
		Models.AreaState{
			Active:        false,
			Name:          "Gslides_1",
			Action:        false,
			Field:         0,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
			PrintName:     "créer un google slide",
		},
		Models.AreaState{
			Active:        false,
			Name:          "Youtube_4",
			Action:        false,
			Field:         1,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
			PrintName:     "s'abonner à une chaîne youtube",
		},
		Models.AreaState{
			Active:        false,
			Name:          "Youtube_5",
			Action:        false,
			Field:         1,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
			PrintName:     "liker une vidéo youtube",
		},
		Models.AreaState{
			Active:        false,
			Name:          "Youtube_1",
			Action:        false,
			Field:         1,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
			PrintName:     "récuperer le nombre d'abonnés sur ytb",
		},
		Models.AreaState{
			Active:        false,
			Name:          "Youtube_2",
			Action:        false,
			Field:         1,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
			PrintName:     "récuperer le nombre de vidéos sur ytb",
		},
		Models.AreaState{
			Active:        false,
			Name:          "Youtube_3",
			Action:        false,
			Field:         1,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
			PrintName:     "récuperer le nombre de vues sur ytb",
		},
	}

	var calendar_areas = []Models.AreaState{
		Models.AreaState{
			Active:        false,
			Name:          "Calendar_1",
			Action:        true,
			Field:         0,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
		},
		Models.AreaState{
			Active:        false,
			Name:          "Calendar_2",
			Action:        false,
			Field:         0,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
		},
	}
	var youtube_areas = []Models.AreaState{
		Models.AreaState{
			Active:        false,
			Name:          "Youtube_1",
			Action:        true,
			Field:         1,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
		},
		Models.AreaState{
			Active:        false,
			Name:          "Youtube_2",
			Action:        true,
			Field:         1,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
		},
		Models.AreaState{
			Active:        false,
			Name:          "Youtube_3",
			Action:        true,
			Field:         1,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
		},
		Models.AreaState{
			Active:        false,
			Name:          "Youtube_4",
			Action:        false,
			Field:         1,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
		},
		Models.AreaState{
			Active:        false,
			Name:          "Youtube_5",
			Action:        false,
			Field:         1,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
		},
	}
	var gdoc_areas = []Models.AreaState{
		Models.AreaState{
			Active:        false,
			Name:          "Gdoc_1",
			Action:        false,
			Field:         1,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
		},
	}
	var gdrive_areas = []Models.AreaState{
		Models.AreaState{
			Active:        false,
			Name:          "Gdrive_1",
			Action:        true,
			Field:         0,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
			PrintName:     "récuperer le dernier drive reçu",
		},
		Models.AreaState{
			Active:        false,
			Name:          "Gmail_2",
			Action:        false,
			Field:         0,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
			PrintName:     "envoyer un mail",
		},
		/*Models.AreaState{
			Active:        false,
			Name:          "Gdrive_2",
			Action:        false,
			Field:         0,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
		},*/
	}
	var gsheet_areas = []Models.AreaState{
		Models.AreaState{
			Active:        false,
			Name:          "Gsheet_1",
			Action:        false,
			Field:         0,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
		},
	}
	/*var weather_areas = []Models.AreaState{
		Models.AreaState{
			Active:        false,
			Name:          "Weather_1",
			Action:        false,
			Field:         0,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
		},
	}*/
	var discord_areas = []Models.AreaState{
		Models.AreaState{
			Active:        false,
			Name:          "Discord_1",
			Action:        true,
			Field:         1,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
		},
		Models.AreaState{
			Active:        false,
			Name:          "Discord_2",
			Action:        true,
			Field:         1,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
		},
		Models.AreaState{
			Active:        false,
			Name:          "Discord_3",
			Action:        false,
			Field:         0,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
		},
	}
	var gform_areas = []Models.AreaState{
		Models.AreaState{
			Active:        false,
			Name:          "Gform_1",
			Field:         0,
			Action:        false,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
		},
	}
	var gslides_areas = []Models.AreaState{
		Models.AreaState{
			Active:        false,
			Name:          "Gslides_1",
			Action:        false,
			Field:         0,
			Description_1: "desc_1",
			Description_2: "desc_2",
			Description_3: "desc_3",
			FieldValue_1:  "",
			FieldValue_2:  "",
			FieldValue_3:  "",
		},
	}

	var services_array = []Models.ServiceState{
		Models.ServiceState{
			Active: false,
			Name:   "Gmail",
			Areas:  gmail_areas,
		},
		Models.ServiceState{
			Active: false,
			Name:   "Calendar",
			Areas:  calendar_areas,
		},
		Models.ServiceState{
			Active: false,
			Name:   "Youtube",
			Areas:  youtube_areas,
		},
		Models.ServiceState{
			Active: false,
			Name:   "Gdoc",
			Areas:  gdoc_areas,
		},
		Models.ServiceState{
			Active: false,
			Name:   "Gdrive",
			Areas:  gdrive_areas,
		},
		Models.ServiceState{
			Active: false,
			Name:   "Gsheet",
			Areas:  gsheet_areas,
		},
		/*Models.ServiceState{
			Active: false,
			Name:   "Weather",
			Areas:  weather_areas,
		},*/
		Models.ServiceState{
			Active: false,
			Name:   "Gform",
			Areas:  gform_areas,
		},
		Models.ServiceState{
			Active: false,
			Name:   "Discord",
			Areas:  discord_areas,
		},
		Models.ServiceState{
			Active: false,
			Name:   "Gslides",
			Areas:  gslides_areas,
		},
	}

	return &Models.User{
		Google_refresh_token: token,
		Google_access_token:  ac_token,
		Email:                email,
		Services:             services_array,
		Password:             password,
	}
}
