package models

import (
	"regexp"

	"github.com/astaxie/beego/validation"
)

const validCharacters = "abcdefghijklmopqrstiuv"

func validateNotEmptyString(sname string, s string, v *validation.Validation) {
	pattern := `^[\p{Latin}]`

	ln := len(s)

	if ln == 0 {
		v.SetError(sname, "Can not be empty")
		return
	}
	re := regexp.MustCompile(pattern)

	/*beego.Debug("===============> s = ", s)*/

	if !re.MatchString(s) {
		/*beego.Debug("=================> error")*/
		v.SetError(sname, "Invalid first character")
	}
}
