package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/validation"
	"gitlab.com/manuel.diaz/sirel/server/api/authproviders"
	"gitlab.com/manuel.diaz/sirel/server/api/authproviders/ldap"
	"gitlab.com/manuel.diaz/sirel/server/api/authproviders/xxx"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
	"gitlab.com/manuel.diaz/sirel/server/api/pkg/cryptoutil"
)

func AccessRolControl(this *beego.Controller, Rol string) {
	beego.Debug("llego a AccessController")
	u, ok := this.Ctx.Input.Data()["User"].(models.User)
	if !ok {
		this.Ctx.Output.SetStatus(http.StatusInternalServerError)
		beego.Error("User data not found in context")
		this.ServeJSON()
		this.StopRun()
	}
	if !u.HaveRol(Rol) {
		this.Ctx.Output.SetStatus(http.StatusUnauthorized)
		this.ServeJSON()
		this.StopRun()
	}
}

const (
	AuthProviderLdap  = "ldap"
	AuthProviderSIREL = "xxx"
)

type AuthProviderType string

func GetAuthProvider() authproviders.Provider {
	conf := beego.AppConfig
	t := conf.String("AUTH_PROVIDER")
	switch t {
	case AuthProviderLdap:
		return ldap.GetProvider(
			conf.String("AdAddress"),
			conf.String("AdSuff"),
			conf.String("AdBDN"),
			conf.String("AdUser"),
			conf.String("AdPassword"),
		)
	default:
		return xxx.GetProvider()
	}
}

/*
func GetLdap() *ldaputil.Ldap {
	conf := beego.AppConfig
	ldap := ldaputil.NewLdapWithAcc(
		conf.String("AdAddress"),
		conf.String("AdSuff"),
		conf.String("AdBDN"),
		conf.String("AdUser"),
		conf.String("AdPassword"),
	)

	return ldap
}
*/

var crypto = cryptoutil.NewJWTCrypt()

func GetCrypto() *cryptoutil.JWTCrypt {
	return crypto
}

func wre(this *beego.Controller, statusCode int, ms ...interface{}) {
	this.Ctx.Output.SetStatus(statusCode)
	if len(ms) > 0 {
		this.Data["json"] = ms[0]
	} else {
		this.Data["json"] = http.StatusText(statusCode)
	}
	this.ServeJSON()
	this.StopRun()
}

func GetAuthorFromInputData(ctx *context.Context) (models.User, error) {
	x := ctx.Input.Data()["Author"]
	if auth, ok := x.(models.User); ok {
		return auth, nil
	}
	return models.User{}, errors.New("Not user founded in ctx.Input.Data[\"Author\"]")
}

func GetAuthor(this *beego.Controller) models.User {
	// Author of request must be loggued
	u, e := GetAuthorFromInputData(this.Ctx)
	if e != nil {
		// Then the authenticator filter fail
		beego.Error(e.Error())
		wre(this, 500)
	}
	return u
}

func ReadInputBody(this *beego.Controller, obj interface{}) {
	e := json.Unmarshal(this.Ctx.Input.RequestBody, &obj)
	if e != nil {
		wre(this, 400)
	}
}

func Validate(this *beego.Controller, obj interface{}) {
	valid := validation.Validation{}
	ok, err := valid.Valid(obj)
	if err != nil {
		beego.Error(err.Error())
		wre(this, 500)
	}
	if !ok {
		beego.Debug(valid.Errors)
		wre(this, 400)
	}
}

func ReadPagAndOrdOptions(this *beego.Controller) PagAndOrdOptions {
	var (
		opt PagAndOrdOptions
		e   error
	)

	opt.Limit, e = this.GetInt("limit")
	if opt.Limit <= 0 || opt.Limit > 100 {
		opt.Limit = 20
	}
	opt.Offset, e = this.GetInt("offset")
	opt.OrderBy = this.GetString("orderby")
	opt.Desc, e = this.GetBool("desc")

	if e != nil {
		wre(this, 400)
	}

	return opt
}

func fmtorder(opt *PagAndOrdOptions) string {
	exp := opt.OrderBy
	if opt.Desc {
		exp = "-" + exp
	}
	return exp
}

type PagAndOrdOptions struct {
	Limit   int
	Offset  int
	OrderBy string
	Desc    bool
}
