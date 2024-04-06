package Action

import (
	"log"
	"strings"
	"fmt"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	gmail "google.golang.org/api/gmail/v1"
)

func Get_last_mail(srv *gmail.Service, context *gin.Context) *gmail.Message {
	user, _ := context.Cookie("email")
	r, err := srv.Users.Messages.List(user).MaxResults(1).Do()
	if err != nil {
		log.Printf("Unable to retrieve messages: %v", err)
		return nil
	}

	if len(r.Messages) == 0 {
		log.Printf("Can't retreive any messages: %v", err)
		return nil
	}

	lastMessage := r.Messages[0]

	readIDs := make([]string, 0)
	req, err := srv.Users.Messages.List(user).Q("in:inbox is:read").Fields("nextPageToken,messages/id").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve messages: %v", err)
	}
	for _, message := range req.Messages {
		readIDs = append(readIDs, message.Id)
	}

	
	lastMessageOpened := false
	for _, readID := range readIDs {
		if lastMessage.Id == readID {
			lastMessageOpened = true
			break
		}
	}
	if lastMessageOpened {
		fmt.Println("The last message has been opened.")
		return nil
	} else {
		fmt.Println("The last message has not been opened.")
		modifiedMessage := &gmail.ModifyMessageRequest{
			RemoveLabelIds: []string{"UNREAD"},
		}
		_, err = srv.Users.Messages.Modify(user, lastMessage.Id, modifiedMessage).Do()
		if err != nil {
			log.Printf("Unable to modify message: %v", err)
			return nil
		}

		message, err := srv.Users.Messages.Get(user, lastMessage.Id).Format("full").Do()
		if err != nil {
			log.Printf("Can't retrieve last message: %v", err)
			return nil

		}
		fmt.Println(message)

		return message
	}

	return nil
}

func Parse_mail(mail *gmail.Message) string {
	var res string
	parts := mail.Payload.Parts
	for _, part := range parts {
		if part.MimeType == "text/plain" {
			data, err := base64.URLEncoding.DecodeString(part.Body.Data)
			if err != nil {
				log.Printf("an error has occured while parsing mail")
				return ""
			}
			res = strings.TrimSpace(strings.ReplaceAll(string(data), "<br>", ""))
			break
		}
	}
	return res
	/*if mail. == "" {
		log.Printf("mail's body is empty")
		return nil
	}

	bodyDecrypted, err := base64.URLEncoding.DecodeString(mail.Payload.Body.Data)
	if err != nil {
		log.Printf("Can't decrypt mail's body: %v", err)
		return nil
	}

	body := string(bodyDecrypted)
	arr := strings.Split(body, "\r\n")
	println("mail parsed")
	return arr*/
}
