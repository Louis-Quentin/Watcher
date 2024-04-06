package Middleware

import (
	"Area/Controllers"
	"Area/Database"
	"Area/Services"
	"github.com/gin-gonic/gin"
)

/*func Add_header(c *gin.Context) {
	c.Writer.Header().Set("Location", Controllers.Set_Login_location())
}*/

func Link_api_to_db(db *Database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("gorm", *db.DB)
		c.Set("services", Services.Initialize_services(c))
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Location", Controllers.Set_Login_location())
		c.Writer.Header().Set("Test-Header", "Test-Value")
		/*c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")*/
		c.Next()
	}
}
