package app

import (
	"github.com/astaxie/beego"
	"github.com/mdiazp/sirel-server/api/models"
	"github.com/mdiazp/sirel-server/api/pkg/authproviders"
	"github.com/mdiazp/sirel-server/api/pkg/authproviders/ldap"
	"github.com/mdiazp/sirel-server/api/pkg/authproviders/xxx"
	"github.com/mdiazp/sirel-server/api/pkg/cryptoutil"
	"github.com/mdiazp/sirel-server/api/pkg/mailsender"
)

var (
	model        models.Model
	crypto       *cryptoutil.JWTCrypt
	ldapProvider authproviders.Provider
	xxxProvider  authproviders.Provider
	mailSender   *MailSender
)

// InitApp ...
func InitApp() {
	model = models.NewModel()
	crypto = cryptoutil.NewJWTCrypt()
	ldapProvider = ldap.GetProvider(
		beego.AppConfig.String("AdAddress"),
		beego.AppConfig.String("AdSuff"),
		beego.AppConfig.String("AdBDN"),
		beego.AppConfig.String("AdUser"),
		beego.AppConfig.String("AdPassword"),
	)
	xxxProvider = xxx.GetProvider()

	mailSender = &MailSender{
		identity: "",
		user:     beego.AppConfig.String("MailSenderUser"),
		password: beego.AppConfig.String("MailSenderPassword"),
		host:     beego.AppConfig.String("MailSenderHost"),
		port:     beego.AppConfig.String("MailSenderPort"),
	}
}

// GetMailSender ...
func GetMailSender() *MailSender {
	return mailSender
}

// MailSender ...
type MailSender struct {
	identity string
	user     string
	password string
	host     string
	port     string
}

// SendMail ...
func (ms *MailSender) SendMail(to, mailBody string) error {
	return mailsender.SendMail(
		mailsender.SMTPServer{
			Host: ms.host,
			Port: ms.port,
		},
		mailsender.Mail{
			SenderID: ms.user,
			ToIds:    []string{to},
			Subject:  "SIGEL",
			Body:     mailBody,
		},
		ms.password,
	)
}

// Model ...
func Model() models.Model {
	return model
}

// Crypto ...
func Crypto() *cryptoutil.JWTCrypt {
	return crypto
}

const (
	// AuthProviderLdap ...
	AuthProviderLdap = "ldap"

	// AuthProviderSIREL ...
	AuthProviderSIREL = "xxx"
)

// AuthProvider ...
func AuthProvider(t string) authproviders.Provider {
	switch t {
	case AuthProviderLdap:
		return ldapProvider
	default:
		return xxxProvider
	}
}
