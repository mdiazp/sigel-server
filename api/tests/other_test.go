package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	h "net/http"
	"testing"

	"gitlab.com/manuel.diaz/sirel/server/api/controllers"
)

const apihost string = "http://localhost:8080"

func Login() {
	var e error

	user := "ManUEL.diaz"
	pass := "123"

	tr := &h.Transport{
		Proxy: nil,
	}
	h.DefaultClient.Transport = tr

	c := controllers.GetCredentials()
	c.User = user
	c.Pass = pass

	var bs []byte
	bs, e = json.Marshal(c)
	var r *h.Response

	if e == nil {
		var rq *h.Request
		if e == nil {
			bf := bytes.NewReader(bs)
			rq, e = h.NewRequest(h.MethodPost, apihost+"/login", bf)
		}
		if e == nil {
			r, e = h.DefaultClient.Do(rq)
		}

		if e != nil {
			println("e: ", e)
			return
		}
	} else {
		println("e: ", e)
		return
	}

	if r.StatusCode != 200 {
		printBody(r)
		return
	}

	println(r.Status)

	body, e := ioutil.ReadAll(r.Body)
	if e != nil {
		println("e: ", e)
		return
	}

	authHd := string(body)
	println("OK")
	println("autHD: %s", authHd)
}

func TestLogin(t *testing.T) {
	Login()
}

func printBody(r *h.Response) {
	fmt.Printf("StatusCode: %d\n", r.StatusCode)
	fmt.Printf("    Status: %s\n", r.Status)

	body, e := ioutil.ReadAll(r.Body)
	if e == nil {
		r.Body.Close()
		fmt.Println(string(body))
	}
}
