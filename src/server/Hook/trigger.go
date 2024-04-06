package Trigger

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"fmt"
)

func Google_webhook(r *gin.Context) {
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	body, _ := ioutil.ReadAll(r.Request.Body)
	fmt.Println(body)
}