package utility

import (
	"bytes"
	"fmt"
	"goroutine-optimize/config"
	"goroutine-optimize/logger"
	"goroutine-optimize/templates"
	"net/smtp"

	"gopkg.in/gomail.v2"
)

func SendEmail(email string, token string) {
	// sender data
	from := "syahruldualima@gmail.com"
	password := "fxbnhvhbcnixzcdn"

	// recive address
	toEmail := email
	to := []string{toEmail}

	// smtp email protocol
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	// create message email
	subject := "Email Verification\r\n"
	// body := "<html><h5>Click this <a href='http://" + config.SERVER_ADDRESS + ":" + config.SERVER_PORT + "/verify?token=" + token + "'/>link</a> for verifcation your email</html>"
	body := "<html><ul><li>tai</li></ul></html>"
	mime := "MIME-version 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n\r\n"

	message := []byte(subject + mime + body)

	// authentication data
	// func PlanAuth(identity, username, password, host string) auth
	auth := smtp.PlainAuth("", from, password, host)

	// send email
	// func sendEmail(addr string, a Auth, from string, to []string, msg []byte) error
	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		logger.Error("\nerror send email : " + err.Error() + "\n")
		return
	}

	logger.Info("\n success send email \n")
}

func SendEmailV2(email, token string) {
	// smtp email protocol
	host := "smtp.gmail.com"
	port := 587

	// sender data
	from := "syahruldualima@gmail.com"
	password := "fxbnhvhbcnixzcdn"

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "this is email verification!")
	// body := `Click this <a href="http://` + config.SERVER_ADDRESS + `:` + config.SERVER_PORT + `/verify?token=` + token + `"/>link</a> for verifcation your email`
	var body bytes.Buffer
	templates.MyTemplates.ExecuteTemplate(&body, "email.html", struct {
		Link string
	}{
		Link: fmt.Sprintf("http://%s:%s/verify?email=%s&token=%s", config.SERVER_ADDRESS, config.SERVER_PORT, email, token),
	})

	m.SetBody("text/html", body.String())

	// Send the email to Bob
	d := gomail.NewDialer(host, port, from, password)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
