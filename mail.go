package main

import (
	"net/smtp"
)

type Mail interface {
	SetHeader(key, val string)
	GetHeader(key string) (val string, ok bool)
	SetBody(msg string)
	getMail() string
}

type MailGun interface {
	SendMail(m Mail) error
}

type SimpleMail struct {
	Headers map[string]string
	body    string
}

func CreateSimpleMail() *SimpleMail {
	return &SimpleMail{
		make(map[string]string),
		"",
	}
}

func (sm *SimpleMail) SetHeader(key, val string) {
	sm.Headers[key] = val
}

func (sm *SimpleMail) GetHeader(key string) (val string, ok bool) {
	val, ok = sm.Headers[key]
	return
}

func (sm *SimpleMail) SetBody(msg string) {
	sm.body = msg
}

func (sm *SimpleMail) getMail() string {
	var m string = ""
	for key, val := range sm.Headers {
		m += key + ": " + val + "\r\n"
	}
	m += "\r\n"
	m += sm.body + "\r\n"
	return m
}

type FixedMailGun struct {
	Host      string
	Port      string
	UserEmail string
	auth      smtp.Auth
}

func SetupFMG(userEmail, password, host, port string) *FixedMailGun {
	return &FixedMailGun{
		host,
		port,
		userEmail,
		smtp.PlainAuth("", userEmail, password, host),
	}
}

func (mg *FixedMailGun) SendMail(msg Mail) error {
	msg.SetHeader("To", mg.UserEmail)
	var from string
	var ok bool
	if from, ok = msg.GetHeader("From"); !ok {
		from = mg.UserEmail
	}
	return smtp.SendMail(mg.Host+":"+mg.Port, mg.auth, from, []string{mg.UserEmail}, []byte(msg.getMail()))
}
