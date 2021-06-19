package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {
	from := "from@gmail.com"

	to := "to@gmail.com"
	//host for gmail
	//host := "smtp.gmail.com"

	auth := smtp.PlainAuth("", from, "password", "smtp.gmail.com")

	message := `To: Oladeji Rafiat <oyebodeoladejiabiodun@gmail.com>
	From: Oyebode Amirdeen <oyebodeoladejiabiodun@gmail.com>
	Subject: Demo mail

	Message: This is a demo mail
	`
	//port for gmail(:587 for TLS and :465 for SSL)
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, []byte(message))
	if err != nil {
		log.Fatalln("Error sending mail:", err)
	}

	fmt.Println("Email sent")
}
