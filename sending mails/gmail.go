package main

import (
	"fmt"
	"log"
	"net/smtp"
	"net/mail"
	//"strings"
	"os"
	"encoding/base64"
)

func main() {
	server := "smtp.gmail.com"

	pass := os.Getenv("emailPassword")

	auth := smtp.PlainAuth("", "oyebodeamirdeen@gmail.com", pass, server)

	from := mail.Address{"Needrima", "oyebodeamirdeen@gmail.com"}
	to := mail.Address{"", "oyebodeamirdeen@outlook.com"}
	to2 := mail.Address{"", "fatokunayodeji0@gmail.com"}
	title := "New blog post"

	body := fmt.Sprintf("New blog post <a href=%s>Visit</a>", "https://www.student-devs-blog.herokuapp.com")

	headers := map[string]string {
		"From": from.String(),
		"To": to.String(),
		"Subject": title,
		"Content-Type": "text/html; charset=utf-8",
		"Content-Transfer-Encoding": "base64",
	}

	var message string
	
	for i, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", i, v)
	}

	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	if err := smtp.SendMail(server+":587", auth, from.Address, []string{to.Address, to2.Address}, []byte(message)); err != nil {
		log.Fatalln("Error sending mail:", err)
	}

	fmt.Println("Sent")
}

// func main() {
// 	from := "oyebodeamirdeen@gmail.com"

// 	receiver := []string{"oyebodeamirdeen@gmail.com"}
	
// 	host := "smtp.gmail.com"

// 	password := os.Getenv("emailPassword")
// 	fmt.Println(password)

// 	auth := smtp.PlainAuth("", from, password, host)

// 	message := `To: Oladeji Rafiat <oladejirafiatade@gmail.com><br>
// 	From: Oyebode Amirdeen <oyebodeoladejiabiodun@gmail.com><br>
// 	Subject: Demo mail

// 	Message: This is a test mail <a href="http://www.google.com">Visit</a>
// 	`
// 	//port for gmail(:587 for TLS and :465 for SSL)
// 	err := smtp.SendMail(host+":587", auth, from, receiver, []byte(message))
// 	if err != nil {
// 		log.Fatalln("Error sending mail:", err)
// 	}

// 	fmt.Println("Email sent")
// }
