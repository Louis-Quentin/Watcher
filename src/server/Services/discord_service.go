package Services

import (
	"Area/Models"
	reaction "Area/Reaction"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
	"strconv"
	"log"
	"net/http"
)

var lastName string
var lastBot string

func Initialize_discord_service(r *gin.Context, token *oauth2.Token) *discordgo.Session {
	discord, err := discordgo.New("Bot " + "MTA3Njg3NTYxMDIzNzk2NDI5OA.GrOfhp.oFFOvutB4tThf2mt731Qp5tTMI48qHoo1YmGRA")

	err = discord.Open()
	if err != nil {
		log.Printf("Can't open discord session: %v", err)
		return nil
	}
	return discord
}

func Get_discord_name(s *discordgo.Session, r *gin.Context, token *oauth2.Token) string {
	email_string, err := r.Cookie("email")

	if err != nil {
		r.Redirect(http.StatusTemporaryRedirect, "/google_login")
	}
	var user Models.User
	db := r.MustGet("gorm").(gorm.DB)
	db.Where(Models.User{Email: email_string}).Preload("Services.Areas").First(&user)
	ID := user.Services[7].Areas[0].FieldValue_1

	userDiscord, err2 := s.User(ID)
	if err2 != nil {
		log.Printf("Error getting user info: %v", err2)
	}
	
	if userDiscord.Username != lastName && lastName != " "{
		fmt.Println(string(userDiscord.Username))
		lastName = string(userDiscord.Username)
		return string(userDiscord.Username)
	} else {
		fmt.Println("On sort de discord")
		return " "
	}

	return ""

}

func Get_bot(s *discordgo.Session, r *gin.Context, token *oauth2.Token) string {
	email_string, err := r.Cookie("email")
	fmt.Println(email_string)

	if err != nil {
		r.Redirect(http.StatusTemporaryRedirect, "/google_login")
	}
	var user Models.User
	db := r.MustGet("gorm").(gorm.DB)
	db.Where(Models.User{Email: email_string}).Preload("Services.Areas").First(&user)
	ID := user.Services[7].Areas[1].FieldValue_1

	userDiscord, err2 := s.User(ID)
	if err2 != nil {
		log.Printf("Error getting user info: %v", err2)
	}

	if userDiscord.Username != lastBot && lastBot != " " {
		fmt.Println(userDiscord.Bot)
		lastBot = string(userDiscord.Username)
		return strconv.FormatBool(userDiscord.Bot)
	} else {
		fmt.Println("On sort de discord 2")
		return " "
	}
	return " "
}

func Discord_1_area(r *gin.Context, token *oauth2.Token) {
	discord_service := Initialize_discord_service(r, token)
	Get_discord_name(discord_service, r, token)
}

func Discord_2_area(r *gin.Context, token *oauth2.Token) {
	discord_service := Initialize_discord_service(r, token)
	Get_bot(discord_service, r, token)
}

func Init_discord_service_object(r *gin.Context) *Service {
	s := new(Service)
	s.Areas = map[string]func(r *gin.Context, token *oauth2.Token){
		"Discord_1": Discord_1_area,
		"Discord_2": Discord_2_area,
		"Discord_3": reaction.Discord_3_area,
	}
	return s
}
