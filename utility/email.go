package utility

import (
	"goroutine-optimize/logger"
	"net/smtp"
)

func SendEmail(email string) {
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

	// message
	subject := "Subject : Email Verification\n"
	body := "this is email verification"
	message := []byte(subject + body)

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
