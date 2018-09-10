package cryptoutil

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	// AuthHd header key of JWT value
	AuthHd = "authHd"
	// MalformedHd is the error message sent when the
	// header of a request is empty
	MalformedHd = "Malformed authHd header"
	// NotJWTUser is the error message sent when the
	// JWTUser type assertion fails. This is a fatal
	// security breach since it can only occurr when
	// the private key is compromised.
	NotJWTUser = `False JWTUser type assertion. 
	Security breach. Private key compromised`
)

// JWTCrypt is the object for encrypting JWT
type JWTCrypt struct {
	pKey *rsa.PrivateKey
}

// JWTUser adds jwt.StandardClaims to an User
type JWTUser struct {
	Username string `json:"user"`
	jwt.StandardClaims
}

// NewJWTCrypt creates a new JWTCrypt
func NewJWTCrypt() (j *JWTCrypt) {
	j = new(JWTCrypt)
	var e error
	j.pKey, e = rsa.GenerateKey(rand.Reader, 512)
	if e != nil {
		panic(e.Error())
	}
	return
}

func (j *JWTCrypt) Encrypt(username string) (s string, e error) {
	uc := &JWTUser{Username: username}
	uc.ExpiresAt = time.Now().Add(time.Hour).Unix()
	t := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), uc)
	s, e = t.SignedString(j.pKey)
	return
}

func (j *JWTCrypt) Decrypt(token string) (username string, e error) {
	if token == "" {
		return "", errors.New(MalformedHd)
	}
	t, e := jwt.ParseWithClaims(token, &JWTUser{},
		func(x *jwt.Token) (a interface{}, d error) {
			a, d = &j.pKey.PublicKey, nil
			return
		})
	var clm *JWTUser
	if e == nil {
		var ok bool
		clm, ok = t.Claims.(*JWTUser)
		if !ok || clm.Username == "" {
			panic(NotJWTUser)
			// { the private key was used to sign something
			//   different from a *JWTUser, which is not
			//   done in this program, therefore it has
			//   been compromised }
		}
	}
	if e == nil {
		username = clm.Username
	}
	return
}
