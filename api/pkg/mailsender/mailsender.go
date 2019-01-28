package mailsender

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"
)

// Mail ...
type Mail struct {
	SenderID string
	ToIds    []string
	Subject  string
	Body     string
}

// SMTPServer ...
type SMTPServer struct {
	Host string
	Port string
}

// ServerName ...
func (s *SMTPServer) ServerName() string {
	return s.Host + ":" + s.Port
}

// BuildMessage ...
func (mail *Mail) BuildMessage() string {
	message := ""
	message += fmt.Sprintf("From: %s\r\n", mail.SenderID)
	if len(mail.ToIds) > 0 {
		message += fmt.Sprintf("To: %s\r\n", strings.Join(mail.ToIds, ";"))
	}

	message += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	message += "\r\n" + mail.Body

	return message
}

// SendMail ...
func SendMail(smtpServer SMTPServer, mail Mail, password string) error {

	messageBody := mail.BuildMessage()

	//build an auth
	auth := smtp.PlainAuth("", mail.SenderID, password, smtpServer.Host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer.Host,
	}

	conn, e := tls.Dial("tcp", smtpServer.ServerName(), tlsconfig)
	if e != nil {
		return e
	}

	client, e := smtp.NewClient(conn, smtpServer.Host)
	if e != nil {
		return e
	}

	// step 1: Use Auth
	if e = client.Auth(auth); e != nil {
		return e
	}

	// step 2: add all from and to
	if e = client.Mail(mail.SenderID); e != nil {
		return e
	}
	for _, k := range mail.ToIds {
		if e = client.Rcpt(k); e != nil {
			return e
		}
	}

	// Data
	w, e := client.Data()
	if e != nil {
		return e
	}

	_, e = w.Write([]byte(messageBody))
	if e != nil {
		return e
	}

	e = w.Close()
	if e != nil {
		return e
	}

	client.Quit()

	return nil

}
