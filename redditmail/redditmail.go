package main

import (
	"bytes"
	"github.com/dexterous/redditnews"
	"log"
	"net/smtp"
)

func main() {
	to := "satish@joshsoftware.com"
	subject := "Go articles on Reddit"
	message := Email()

	body := "To: " + to + "\r\nSubject: " +
		subject + "\r\n\r\n" + message

	auth := smtp.PlainAuth("", "satish.talim", "xlzcalahblgmklac", "smtp.gmail.com")
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"satish.talim@gmail.com",
		[]string{to},
		[]byte(body))
	if err != nil {
		log.Fatal("SendMail: ", err)
		return
	}
}

// Email prepares the body of an email
func Email() string {
	var buffer bytes.Buffer

	items, err := redditnews.Get("golang")
	if err != nil {
		log.Fatal(err)
	}

	// Need to build strings from items
	for _, item := range items {
		buffer.WriteString(item.String())
	}

	return buffer.String()
}
