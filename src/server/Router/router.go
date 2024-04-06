package Router

import (
	"Area/Controllers"
	"Area/HandleLifeCycle"
	"Area/Models"
	_ "Area/Models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

type AboutStruct struct {
	Str []byte `json:"value"`
}

func Handle_home_request(context *gin.Context) {
	var msg string
	if err := context.ShouldBind(&msg); err != nil {
		context.JSON(400, gin.H{"Error": "Wrong home request format"})
	} else {
		var html_index = `<html>
<body>
		<a href="/google_login">Google Log in<a/>
</body>
</html>`
		context.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html_index))
	}
}

func Handle_start_life_cycle(context *gin.Context) {
	go HandleLifeCycle.Set_life_cycle(context)
}

func getLocalAddress() string {
	con, error := net.Dial("udp", "8.8.8.8:80")
	if error != nil {
		log.Fatal(error)
	}
	defer con.Close()

	localAddress := con.LocalAddr().(*net.UDPAddr)

	return localAddress.IP.String()
}

func Handle_mobile_google_login(r *gin.Context) {
	var body struct {
		GoogleAccessToken string
		Mail              string
	}
	if r.Bind(&body) != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
	}
	println("mobile token = ", body.GoogleAccessToken)
	println("mobile mail = ", body.Mail)
	//var user Models.User
	db := r.MustGet("gorm").(gorm.DB)

	r.SetCookie("email", body.Mail, 3600*24*30, "", "", false, false)

	var user Models.User
	res := db.Where(Models.User{Email: body.Mail}).First(&user)
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		println("not found")
	}
	if res.RowsAffected > 0 {
		user.Google_access_token = body.GoogleAccessToken
		res = db.Model(&user).Updates(&user)
		if res.Error != nil {
			println("update failed")
		}
	} else {
		db.FirstOrCreate(Controllers.Build_user("", body.Mail, body.GoogleAccessToken, ""), Models.User{Email: body.Mail})
	}
	/*if err := db.Where(Models.User{Email: body.Mail}).
		Assign(Models.User{Email: body.Mail, Google_refresh_token: "", Google_access_token: body.GoogleAccessToken, Password: ""}).
		FirstOrCreate(Controllers.Build_user("", body.Mail, body.GoogleAccessToken, "")).Error; err != nil {
		r.Next()
		return
	}*/
	println("success")
	r.JSON(http.StatusOK, gin.H{
		"Google login": "success",
	})

}

func Handle_about_request(context *gin.Context) {

	//println("IN ABOUT REQUEST")
	email_string, err := context.Cookie("email")

	if err != nil {
		context.Redirect(http.StatusTemporaryRedirect, "/google_login")
		return
	}

	var user Models.User
	db := context.MustGet("gorm").(gorm.DB)
	db.Model(&user).Preload("Services.Areas").First(&user, "email= ?", email_string)
	//println("db loaded and user too")
	if user.ID == 0 {
		context.Redirect(http.StatusTemporaryRedirect, "/google_login")
		return
	}
	res, err := json.Marshal(&user)
	if err != nil {
		log.Fatalf("le marshall a fail = %v", err)
	}
	//println("json res = ", string(res))

	//println(res)
	//println(string(res))
	context.JSON(http.StatusOK, string(res))
	//println("END of about request")
}

func About(context *gin.Context) {
	curTime := time.Now().Unix()
	curTimeStr := strconv.Itoa(int(curTime))
	ipString := getLocalAddress()

	about := `{
		"client":{
		"host":"` + ipString + `"
		},
		"server":{
		"current_time":"` + curTimeStr + `",
		"services":[
			{
				"name":"Gcalendar",
				"actions":[
					{
					"name":"get_last_calendar",
					"description":"Get last google calendar of google user"
					}
				],
				"reactions":[
					{
					"name":"create_calendar",
					"description":"A google calendar event is created"
					}
				]
			},
			{
				"name":"Discord",
				"actions":[
					{
					"name":"get_discord_name",
					"description":"Get tag of discord user with ID"
					},
					{
					"name":"get_bot",
					"description":"Know is the user given is a bot or no"
					}
				],
				"reactions":[
					{
					"name":"send_message",
					"description":"Send a discord message"
					}
				]
			},
			{
				"name":"Gdoc",
				"reactions":[
					{
					"name":"create_doc",
					"description":"A google doc is created"
					}
				]
			},
			{
				"name":"Gdrive",
				"actions":[
					{
					"name":"get_last_drive",
					"description":"Get last google drive of google user"
					}
				],
				"reactions":[
					{
					"name":"create_drive",
					"description":"A google drive is created"
					}
				]
			},
			{
				"name":"Gform",
				"reactions":[
					{
					"name":"create_form",
					"description":"A google form is created"
					}
				]
			},
			{
				"name":"Gmail",
				"actions":[
					{
					"name":"get_last_mail",
					"description":"Get last google mail of google user"
					}
				],
				"reactions":[
					{
					"name":"send_mail",
					"description":"A google mail is sended"
					}
				]
			},
			{
				"name":"Gsheet",
				"reactions":[
					{
					"name":"create_sheet",
					"description":"A google sheet is created"
					}
				]
			},
			{
				"name":"Gslides",
				"reactions":[
					{
					"name":"create_slides",
					"description":"A google slides is created"
					}
				]
			},
			{
				"name":"Youtube",
				"actions":[
					{
					"name":"get_channel_nb_user",
					"description":"Get number of subscribers of chanel given"
					},
					{
					"name":"get_channel_nb_videos",
					"description":"Get number of videos of chanel given"
					},
					{
					"name":"get_channel_nb_views",
					"description":"Get number of views of chanel given"
					}
				],
				"reactions":[
					{
					"name":"subscribe",
					"description":"The user is subscribed to the channel given"
					},
					{
					"name":"like_video",
					"description":"The user will like a video from given url"
					}
				]
			}
		]
		}
	}`
	var aboutJSON map[string]interface{}
	err := json.Unmarshal([]byte(about), &aboutJSON)
	if err != nil {
		log.Fatalf("Erreur lors du décodage JSON : ", err)
	}
	context.JSON(http.StatusOK, aboutJSON)
}
