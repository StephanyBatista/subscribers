package email

import (
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

func Send(from, to, subject, body, subscriberId string) bool {

	if os.Getenv("sub_gotest") == "true" {
		return false
	}

	link := os.Getenv("sub_url_app") + "/subscribers/" + subscriberId + "/read"
	img := `<img src="` + link + `">`
	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", `<b>message `+img+`</b>`)

	n := gomail.NewDialer(
		os.Getenv("sub_email_smtp"),
		587,
		os.Getenv("sub_email_user"),
		os.Getenv("sub_email_password"))

	if err := n.DialAndSend(msg); err != nil {
		log.Println(err.Error())
		return false
	}

	return true
}
