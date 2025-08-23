package mailclient

import (
	"log"
	"net/smtp"
)

func smtpClient() {
	client, err := smtp.Dial("localhost:1025")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	err = client.Mail("me@test.com")
	if err != nil {
		log.Fatal(err)
	}
	err = client.Rcpt("you@test.com")
	if err != nil {
		log.Fatal(err)
	}
	w, err := client.Data()
	if err != nil {
		log.Fatal()
	}

	_, err = w.Write([]byte(`
		From: me@test.com
		To: you@test.com
		Subject: Hey bro

		Hey dude bro man guy!
		`))
	if err != nil {
		log.Fatal(err)
	}
	w.Close()
	client.Quit()
}
