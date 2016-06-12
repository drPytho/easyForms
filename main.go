package main

import (
	"flag"
	"log"
)

var (
	testMail bool
	email    string
	password string
	host     string
	port     string
)

func init() {
	flag.BoolVar(&testMail, "test", false, "This can be set to test the mail configuraton. This will then send a standard mail to test")
	flag.StringVar(&email, "email", "", "This is the email to witch the messages will be sent, and used to log into the email server")
	flag.StringVar(&password, "password", "", "This is the password to the email server for the user email")
	flag.StringVar(&host, "host", "", "This is the server host name e.g. smtp.mailserver.se")
	flag.StringVar(&port, "port", "25", "This is the port to which this service will connect to the mail server")
}

func main() {
	log.Println("Starting up form submition server")
	flag.Parse()

	if email == "" || host == "" {
		log.Println("You need to enter both email and host name to start this service")
		return
	}

	//Construc the mailgun
	fmg := SetupFMG(email, password, host, port)

	if testMail {
		// Test the mail interface
		s := CreateSimpleMail()
		s.SetHeader("From", email)
		s.SetHeader("Subject", "This is a test")
		s.SetBody("This is a test mail to make sure that the mail configuration is correct on your form submition service")

		if err := fmg.SendMail(s); err != nil {
			// Woppsie
			log.Println(err.Error())
			return
		}
		log.Println("Email was sent successfully")
	}

}
