package Services

import (
	"Area/Models"
	reaction "Area/Reaction"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	youtube "google.golang.org/api/youtube/v3"
	"gorm.io/gorm"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

var lastChannel string
var lastVideo string
var lastView string

func Initialize_youtube_service(r *gin.Context, token *oauth2.Token) *youtube.Service {
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

	srv, err := youtube.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Printf("Unable to create youtube service: %v", err)
	}
	return srv
}

func Get_channel_nb_user(srv *youtube.Service, r *gin.Context, token *oauth2.Token) string {
	email_string, err := r.Cookie("email")

	if err != nil {
		r.Redirect(http.StatusTemporaryRedirect, "/google_login")
	}
	var user Models.User
	db := r.MustGet("gorm").(gorm.DB)
	db.Where(Models.User{Email: email_string}).Preload("Services.Areas").First(&user)
	// println("ac token = ", token.AccessToken)
	// println("rf token = ", token.RefreshToken)

	err2 := db.Where(Models.User{Email: email_string}).Preload("Services.Areas").First(&user)
	if err2 != nil {
		log.Printf("Erreur de fdp: %v", err2)
	}
	channelName := user.Services[2].Areas[0].FieldValue_1

	//if channelName != lastChannel && lastChannel != " " {

	searchResponse, err := srv.Search.List([]string{"id"}).Q(channelName).Type("channel").MaxResults(1).Do()

	if err != nil {
		log.Printf("search for channel: %v", err)
	}

	if len(searchResponse.Items) == 0 {
		log.Printf("No channel found with name %s", channelName)
	}

	channelID := searchResponse.Items[0].Id.ChannelId

	channelResponse, err := srv.Channels.List([]string{"snippet", "statistics"}).Id(channelID).Do()

	if err != nil {
		log.Printf("Failed to get channel information: %v", err)
	}

	channel := channelResponse.Items[0]
	println("nb subscribed peoples = ", channel.Statistics.SubscriberCount)
	lastChannel = channelName
	return strconv.FormatUint(channel.Statistics.SubscriberCount, 10)
	/*} else {
		log.Println("On sort de YTB")
		return " "
	}*/

	//return ""
}

func Get_channel_nb_videos(srv *youtube.Service, r *gin.Context, token *oauth2.Token) string {
	email_string, err := r.Cookie("email")

	if err != nil {
		r.Redirect(http.StatusTemporaryRedirect, "/google_login")
	}
	var user Models.User
	db := r.MustGet("gorm").(gorm.DB)
	db.Where(Models.User{Email: email_string}).Preload("Services.Areas").First(&user)
	channelName := user.Services[2].Areas[1].FieldValue_1

	searchResponse, err := srv.Search.List([]string{"id"}).Q(channelName).Type("channel").MaxResults(1).Do()

	if err != nil {
		log.Printf("Failed to search for channel: %v", err)
	}

	if len(searchResponse.Items) == 0 {
		log.Printf("No channel found with name %s", channelName)
	}

	channelID := searchResponse.Items[0].Id.ChannelId

	//	if channelName != lastVideo && lastVideo != " " {

	channelResponse, err := srv.Channels.List([]string{"snippet", "statistics"}).Id(channelID).Do()

	if err != nil {
		log.Printf("Failed to get channel information: %v", err)
	}

	channel := channelResponse.Items[0]
	fmt.Println(channel.Statistics.VideoCount)
	lastVideo = channelName
	return strconv.FormatUint(channel.Statistics.VideoCount, 10)
	/*	} else {
			log.Println("On sort de YTB 2")
			return " "
		}
		return ""*/
}

func Get_channel_nb_views(srv *youtube.Service, r *gin.Context, token *oauth2.Token) string {

	email_string, err := r.Cookie("email")

	if err != nil {
		r.Redirect(http.StatusTemporaryRedirect, "/google_login")
	}
	var user Models.User
	db := r.MustGet("gorm").(gorm.DB)
	db.Where(Models.User{Email: email_string}).Preload("Services.Areas").First(&user)
	channelName := user.Services[2].Areas[2].FieldValue_1
	searchResponse, err := srv.Search.List([]string{"id"}).Q(channelName).Type("channel").MaxResults(1).Do()

	if err != nil {
		log.Printf("Failed to search for channel: %v", err)
	}

	if len(searchResponse.Items) == 0 {
		log.Printf("No channel found with name %s", channelName)
	}

	channelID := searchResponse.Items[0].Id.ChannelId

	channelResponse, err := srv.Channels.List([]string{"snippet", "statistics"}).Id(channelID).Do()

	if err != nil {
		log.Printf("Failed to get channel information: %v", err)
	}
	//if channelName != lastView && lastView != " " {
	channel := channelResponse.Items[0]
	fmt.Println(channel.Statistics.ViewCount)
	lastView = channelName
	return strconv.FormatUint(channel.Statistics.ViewCount, 10)
	/*} else {
		log.Println("On sort de ytb 3")
		return " "
	}
	return ""*/
}

func Subscribe(srv *youtube.Service, r *gin.Context, token *oauth2.Token) {
	email_string, err := r.Cookie("email")

	if err != nil {
		r.Redirect(http.StatusTemporaryRedirect, "/google_login")
	}
	var user Models.User
	db := r.MustGet("gorm").(gorm.DB)
	db.Where(Models.User{Email: email_string}).Preload("Services.Areas").First(&user)
	channelName := user.Services[2].Areas[3].FieldValue_1
	searchResponse, err := srv.Search.List([]string{"id"}).Q(channelName).Type("channel").MaxResults(1).Do()

	if err != nil {
		log.Printf("Failed to search for channel: %v", err)
	}

	if len(searchResponse.Items) == 0 {
		log.Printf("No channel found with name %s", channelName)
	}

	channelID := searchResponse.Items[0].Id.ChannelId

	subscription := &youtube.Subscription{
		Snippet: &youtube.SubscriptionSnippet{
			ResourceId: &youtube.ResourceId{
				Kind:      "youtube#channel",
				ChannelId: channelID,
			},
		},
	}

	_, err = srv.Subscriptions.Insert([]string{"snippet"}, subscription).Do()
	if err != nil {
		log.Printf("Error subscribing to channel: %v", err)
	}
}

func Like(srv *youtube.Service, r *gin.Context, token *oauth2.Token) {
	email_string, err := r.Cookie("email")

	if err != nil {
		r.Redirect(http.StatusTemporaryRedirect, "/google_login")
	}
	var user Models.User
	db := r.MustGet("gorm").(gorm.DB)
	db.Where(Models.User{Email: email_string}).Preload("Services.Areas").First(&user)
	link := user.Services[2].Areas[4].FieldValue_1
	//link = "https://www.youtube.com/watch?v=io7-o5Yyk7Q"
	println("link get = ", link)

	videoURL, err := url.Parse(link)
	if err != nil {
		log.Printf("Can't parse link: %v", err)
	}
	videoID := videoURL.Query().Get("v")

	like := youtube.VideoRating{}
	like.Rating = "like"

	srv.Videos.Rate(videoID, "like").Do()

	fmt.Println("La vidéo est likée")
}

func Youtube_1_area(r *gin.Context, token *oauth2.Token) {
	//println("ON EST DANS L'EXECUTION DE L'ACTION DE L'AREA !!")

	email_string, err := r.Cookie("email")
	if err != nil {
		println("failed to retrieve email")
		r.Redirect(http.StatusTemporaryRedirect, "/google_login")
		return
	}
	println("cookie passed")
	youtube_service := Initialize_youtube_service(r, token)
	res := Get_channel_nb_user(youtube_service, r, token)
	var user Models.User
	db := r.MustGet("gorm").(gorm.DB)
	db.Where(Models.User{Email: email_string}).Preload("Services.Areas").First(&user)
	println("RES = ", res)
	user.Services[0].Areas[5].FieldValue_1 = res
	db.Save(&user.Services)
	db.Save(&user.Services[0].Areas)
	println("YOUTUBE 1 AREA going to send discord msg")
	reaction.Discord_3_area(r, token)
}

func Youtube_2_area(r *gin.Context, token *oauth2.Token) {
	email_string, err := r.Cookie("email")

	if err != nil {
		r.Redirect(http.StatusTemporaryRedirect, "/google_login")
		return
	}
	youtube_service := Initialize_youtube_service(r, token)
	res := Get_channel_nb_videos(youtube_service, r, token)
	var user Models.User
	db := r.MustGet("gorm").(gorm.DB)
	db.Where(Models.User{Email: email_string}).Preload("Services.Areas").First(&user)
	user.Services[0].Areas[5].FieldValue_1 = res
	db.Save(&user.Services[0].Areas)
	db.Save(&user.Services)
	reaction.Discord_3_area(r, token)
}

func Youtube_3_area(r *gin.Context, token *oauth2.Token) {
	email_string, err := r.Cookie("email")

	if err != nil {
		r.Redirect(http.StatusTemporaryRedirect, "/google_login")
		return
	}
	println("in yt view area after cookie get")
	youtube_service := Initialize_youtube_service(r, token)
	res := Get_channel_nb_views(youtube_service, r, token)
	var user Models.User
	db := r.MustGet("gorm").(gorm.DB)
	db.Where(Models.User{Email: email_string}).Preload("Services.Areas").First(&user)
	user.Services[0].Areas[5].FieldValue_1 = res
	db.Save(&user.Services)
	db.Save(&user.Services[0].Areas)
	reaction.Discord_3_area(r, token)
}

func Youtube_4_area(r *gin.Context, token *oauth2.Token) {
	println("ON EST DANS L'EXECUTION DE youtube 4!!")

	youtube_service := Initialize_youtube_service(r, token)
	Subscribe(youtube_service, r, token)
}

func Youtube_5_area(r *gin.Context, token *oauth2.Token) {
	println("ON EST DANS L'EXECUTION DE youtube 5!!")
	youtube_service := Initialize_youtube_service(r, token)

	Like(youtube_service, r, token)
}

func Init_Youtube_service_object(r *gin.Context) *Service {
	s := new(Service)
	s.Areas = map[string]func(r *gin.Context, token *oauth2.Token){
		"Youtube_1": Youtube_1_area,
		"Youtube_2": Youtube_2_area,
		"Youtube_3": Youtube_3_area,
		"Youtube_4": Youtube_4_area,
		"Youtube_5": Youtube_5_area,
	}
	return s
}
