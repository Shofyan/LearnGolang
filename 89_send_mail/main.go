package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"net/smtp"
	"strings"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_EMAIL = "tokonabila45@gmail.com"
const CONFIG_PASSWORD = "JamaludinHasan"

func sendMail(to []string, cc []string, subject, message string) error {

	body := "from :" + CONFIG_EMAIL + "\n" +
		"TO : " + strings.Join(to, ",") + "\n" +
		"CC:" + strings.Join(cc, ",") + "\n" +
		"Subject:" + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", CONFIG_EMAIL, CONFIG_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, CONFIG_EMAIL, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}

func main() {
	to := []string{"shofyan@agricon.com"}
	cc := []string{"shofyan@terminix.co.id"}
	subject := "test email golang"
	message := "test email body hello"

	err := sendMail(to, cc, subject, message)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_EMAIL)
	mailer.SetHeader("To", "shofyan@agricon.com")
	mailer.SetAddressHeader("CC", "shofyan@terminix.co.id", "shofyan terminix")
	mailer.SetHeader("Subject", "tes email hoho")
	mailer.SetBody("text/html", "hello , <b> have a nive day </b>")
	mailer.Attach("./sample.png")

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_EMAIL,
		CONFIG_PASSWORD,
	)

	err = dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent! using gomail.v2")

}
