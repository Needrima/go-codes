package main

import (
	//"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
)

func main() {
	mail := gomail.NewMessage()

	mail.SetHeader("From", "from@gmail.com")

	mail.SetHeader("To", "to@gmail.com")

	mail.SetHeader("Subject", "Test email")

	mail.SetBody("text/plain", "This is a test mail")

	dialer := gomail.NewPlainDialer("smtp.gmail.com", 587, "from@gmail.com", "password")

	//dialer.TLSConfig = &tls.Config{InsecureSkipVerify: false}

	if err := dialer.DialAndSend(mail); err != nil {
		fmt.Println("Error sending mail:", err)
		return
	}
	fmt.Println("Mail sent")
}
