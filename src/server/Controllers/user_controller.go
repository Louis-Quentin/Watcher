package Controllers

import (
	"Area/Models"
	"encoding/json"
	_ "fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	_ "os"
)

func Signup(r *gin.Context) {
	println("register BACK FUNC")
	var body struct {
		Email    string
		Password string
	}
	err := json.NewDecoder(r.Request.Body).Decode(&body)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid body",
		})
	}
	/*hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash the password",
		})
		return
	}*/
	//println("the NEW BODY EMAIL IS: ", body.Email)
	//println("THE NEW BODY PASSWORD IS: ", body.Password)
	/*fmt.Printf("Email = %s, Password = %s\n", body.Email, body.Password)
	 */
	r.SetCookie("email", body.Email, 3600*24*30, "", "", false, false)
	db := r.MustGet("gorm").(gorm.DB)
	var user Models.User
	res := db.Where(Models.User{Email: body.Email}).First(&user)
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		println("not found")
	}
	if res.RowsAffected > 0 {
		res = db.Model(&user).Updates(Models.User{Password: body.Password})
		if res.Error != nil {
			println("update failed")
		}
	} else {
		db.FirstOrCreate(Build_user("", body.Email, "", body.Password), Models.User{Email: body.Email})
	}
	//Debug.Print_all_users(r)
	/*if err := db.Where(Models.User{Email: body.Email}).
		Assign(Models.User{Email: body.Email, Google_refresh_token: "", Google_access_token: "", Password: body.Password}).
		FirstOrCreate(Build_user("", body.Email, "", body.Password)).Error; err != nil {
		r.Next()
		println("failed to build user with email = ", body.Email)
		return
	}*/
	r.Redirect(http.StatusTemporaryRedirect, "http://localhost:8081/")
}

func Login(r *gin.Context) {
	println("LOGIN BACK FUNC")
	var body struct {
		Email    string
		Password string
	}

	err := json.NewDecoder(r.Request.Body).Decode(&body)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid body",
		})
	}
	var user Models.User
	db := r.MustGet("gorm").(gorm.DB)
	db.First(&user, "email= ?", body.Email)

	if user.ID == 0 {
		r.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email",
		})
		return
	}
	if user.Password != body.Password {
		r.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password",
		})
		return
	}
	/*err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Password",
		})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	token_string, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}*/
	//r.SetSameSite(http.SameSiteLaxMode)
	//r.SetCookie("Authorization", token_string, 3600*24*30, "", "", false, false)

	r.SetCookie("email", body.Email, 3600*24*30, "", "", false, false)
	r.Redirect(http.StatusTemporaryRedirect, "http://localhost:8081/")
}
