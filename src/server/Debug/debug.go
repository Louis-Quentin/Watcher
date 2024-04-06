package Debug

import (
	"Area/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Print_all_users(r *gin.Context) {
	db := r.MustGet("gorm").(gorm.DB)
	println("DEBUG MOD START\n")
	var users []Models.User
	result := db.Model(&users).Preload("Services.Areas").Find(&users)
	fmt.Printf("there is %d users in db", result.RowsAffected)
	for _, user := range users {
		println(user.Email)
		println(user.Google_refresh_token)
		for _, service := range user.Services {
			println(service.Name)
			println(service.Active)
			for _, area := range service.Areas {
				println(area.Name)
				println(area.Active)
			}
		}
	}
	println("DEBUG MOD FINISHED\n")
}
