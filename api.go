package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetFormHandler(m MailGun) http.HandlerFunc {
	// Number of allowed messages per client IP address
	spam := CreateSpamFilter(50)
	return func(w http.ResponseWriter, r *http.Request) {
		//TODO add spam filter here
		if !spam.OK(r.RemoteAddr) {
			// This person has sent a message more than 50 times
			log.Println("I belive we have cought ourselve a spammer. IP: ", r.RemoteAddr)
			return
		}
		decoder := json.NewDecoder(r.Body)
		data := make(map[string]string)
		if err := decoder.Decode(&data); err != nil {
			// Error while parsing
			log.Println("Error while parsing requests body")
			// Maybe print som details
			log.Printf("Error from JSON encoder \n%s \nBody: %s", err.Error(), decoder.Buffered())
			return
		}
		//TODO: Escape user data
		s := CreateSimpleMail()
		for k := range data {
			if k == "email" || k == "Email" {
				s.SetHeader("From", data[k])
				s.SetHeader("Reply-To", data[k])
			}
		}
		s.SetHeader("MIME-Version", "1.0")
		s.SetHeader("Content-Type", "text/html")
		s.SetHeader("Subject", "Message from contact form")
		bodyMsg := "<html><body>"
		//Build message here
		bodyMsg += "<h3>User message</h3>"
		for k, v := range data {
			bodyMsg += "<b>" + k + "</b>: " + v + "<br><hr>\n"
		}
		bodyMsg += "</body></html>"
		s.SetBody(bodyMsg)
		m.SendMail(s)
	}
}
