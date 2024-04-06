package main

import (
	controller "Area/Controllers"
	"Area/Database"
	"Area/Middleware"
	_ "Area/Models"
	router "Area/Router"
	_ "fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	var db Database.Database
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:8081"
		},
		MaxAge: 12 * time.Hour,
	}))
	err := db.Init_database()
	if err != nil {
		log.Fatal(err)
	}

	apply_routes(r, &db)
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func apply_routes(r *gin.Engine, db *Database.Database) {
	println("START APPLY ROUTES")
	r.Use(Middleware.Link_api_to_db(db))
	r.GET("/", router.Handle_home_request)
	r.POST("/signup", controller.Signup)
	r.POST("/login", controller.Login)
	//r.GET("/google_callback", controller.Google_callback)
	//r.GET("/google_login", controller.Google_login)
	//r.POST("/update_area" /*Middleware.Require_auth,*/ /*,Middleware.Require_google_auth*/, controller.Update_area)
	//r.GET("/start_back", router.Handle_start_life_cycle)
	//r.GET("/about", router.Handle_about_request)
	//r.GET("/about.json", router.About)
	//r.POST("/mobile_google_login", router.Handle_mobile_google_login)
	//r.GET("/gmail_init", service.Initialize_gmail_service)
}
