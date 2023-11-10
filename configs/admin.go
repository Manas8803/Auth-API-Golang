package configs

import (
	"log"
	"net/smtp"
)

func NotifyAdmin(err error) error {
	auth := smtp.PlainAuth("", EMAIL(), PASSWORD(), "smtp.gmail.com")

	email := ADMIN()
	to := []string{email}

	message := []byte(
		"To:" + email + "\r\n" +
			"Subject: Error\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"utf-8\"\r\n\r\n" +
			"<html>" +
			"<head>" +
			"<title>Error in deployment</title>" +
			"</head>" +
			"<body style=\"font-family: Arial, sans-serif;\">" +
			"<div style=\"padding: 20px;\">" +
			"<h1 style=\"color: #333;\">An error occured just now!!!</h1>" +
			"<p style=\"font-size: 16px;\">ERROR : <strong>" + err.Error() + "</strong></p>" +
			"</div>" +
			"</body>" +
			"</html>")

	sendEmailErr := smtp.SendMail("smtp.gmail.com:587", auth, EMAIL(), to, message)

	log.Fatal(sendEmailErr)
	return sendEmailErr
}
