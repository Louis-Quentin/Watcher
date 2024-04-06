package Controllers

import (
	"Area/Models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Update_area(r *gin.Context) {
	//println("ICI ON VA PARSER L'AREA (ACTIVER/DÉSACTIVER)")

	email_string, err := r.Cookie("email")

	if err != nil {
		r.Redirect(http.StatusTemporaryRedirect, "/google_login")
		return
	}

	var user Models.User
	db := r.MustGet("gorm").(gorm.DB)
	db.Model(&user).Preload("Services.Areas").First(&user, "email= ?", email_string)

	if user.ID == 0 {
		r.JSON(http.StatusBadRequest, gin.H{
			"error": "Account not found",
		})
		r.Redirect(http.StatusTemporaryRedirect, "/google_login")
		return
	}

	var body struct {
		Service      string
		Area         string
		State        bool
		FieldValue_1 string
		FieldValue_2 string
		FieldValue_3 string
	}

	err = json.NewDecoder(r.Request.Body).Decode(&body)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid body",
		})
	}

	println("handle area")
	println(body.Area)
	//println(body.FieldValue_1)
	//println(body.FieldValue_2)
	//println(body.FieldValue_3)
	println(body.State)
	println(body.Service)
	println("handle area end")
	for i, service := range user.Services {
		//println("service name = ", service.Name)
		if service.Name == body.Service {
			//println("service name found = ", service.Name)
			if body.State == true {
				//println("set services to true")
				user.Services[i].Active = body.State
			}
			for j, area := range service.Areas {
				if area.Name == body.Area {
					//println("area name found = ", area.Name)
					user.Services[i].Areas[j].Active = body.State
					user.Services[i].Areas[j].FieldValue_1 = body.FieldValue_1
					user.Services[i].Areas[j].FieldValue_2 = body.FieldValue_2
					user.Services[i].Areas[j].FieldValue_3 = body.FieldValue_3
					db.Save(&user.Services)
					db.Save(&service.Areas)
					//println("user area updated in db")
				}
			}
		}
	}
}
