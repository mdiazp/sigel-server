package app

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
	"gitlab.com/manuel.diaz/sirel/server/api/pkg/authproviders"
	"gitlab.com/manuel.diaz/sirel/server/api/pkg/authproviders/ldap"
	"gitlab.com/manuel.diaz/sirel/server/api/pkg/authproviders/xxx"
	"gitlab.com/manuel.diaz/sirel/server/api/pkg/cryptoutil"
)

var (
	model        orm.Ormer
	crypto       *cryptoutil.JWTCrypt
	ldapProvider authproviders.Provider
	xxxProvider  authproviders.Provider
)

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
}

func Model() orm.Ormer {
	return model
}

func Crypto() *cryptoutil.JWTCrypt {
	return crypto
}

const (
	AuthProviderLdap  = "ldap"
	AuthProviderSIREL = "xxx"
)

func AuthProvider(t string) authproviders.Provider {
	switch t {
	case AuthProviderLdap:
		return ldapProvider
	default:
		return xxxProvider
	}
}