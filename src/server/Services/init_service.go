package Services

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type Area struct {
}

type Service struct {
	Areas map[string]func(r *gin.Context, token *oauth2.Token)
}

func Initialize_services(r *gin.Context) map[string]*Service {
	var services = map[string]*Service{
		"Gmail":    Init_gmail_service_object(r),
		"Calendar": Init_calendar_service_object(r),
		"Youtube":  Init_Youtube_service_object(r),
		"Gdoc":     Init_gdoc_service_object(r),
		"Gdrive":   Init_gdrive_service_object(r),
		"Gsheet":   Init_gsheet_service_object(r),
		"Gform":    Init_gform_service_object(r),
		"Discord":  Init_discord_service_object(r),
		"Gslides":  Init_gslides_service_object(r),
	}
	return services
}
