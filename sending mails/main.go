package main

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"os"
)

func main() {
	mail := gomail.NewMessage()

	mail.SetHeader("From", mail.FormatAddress("oyebodeamirdeen@gmail.com", "Needrima"))

	mail.SetHeaders(map[string][]string{
		//"From": {mail.FormatAddress("oyebodeamirdeen@gmail.com", "Needrima")},
		"To" : {"oyebodeamirdeen@outlook.com", "oyebodeamirdeen@gmail.com", "legendarydancelord@gmail.com", "oladejirafiatade@gmail.com"},
		"Subject": {"Test Mail"},
	})

	password := os.Getenv("emailPassword")
	fmt.Println("Password:", password)

	mail.SetBody("text/html", `New blog alert at student dev blog <a style="color:red;" href="http://www.student-devs-blog.herokuapp.com">Visit</a>`)

	dialer := gomail.NewPlainDialer("smtp.gmail.com", 587, "oyebodeamirdeen@gmail.com", password)

	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dialer.DialAndSend(mail); err != nil {
		fmt.Println("Error sending mail:", err)
		return
	}
	fmt.Println("Mail sent")
}
