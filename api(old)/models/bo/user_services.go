package bo

import (
	"github.com/astaxie/beego/orm"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

var uzero = models.User{}

func (this *Model) CreateUser(u models.User) (models.User, error) {
	u.Id = 0

	// Validate unique username constraint
	e := this.orm.QueryTable(&models.User{}).Filter("username", u.Username).Limit(1).One(&models.User{})
	if e == nil {
		return uzero, models.ErrUserWithSameUsernameAlreadyExist
	}
	if e != orm.ErrNoRows {
		//internal error
		return uzero, e
	}

	// then the unique username constraint it not violated
	// and user can be inserted
	_, e = this.orm.Insert(&u)
	if e != nil {
		return uzero, e
	}

	return u, nil
}

func (this *Model) GetUserById(id int) (models.User, error) {
	u := models.User{}
	e := this.orm.QueryTable(&models.User{}).Filter("id", id).Limit(1).One(&u)

	if e == orm.ErrNoRows {
		return uzero, models.ErrResultNotFound
	}
	if e != nil {
		return uzero, e
	}

	return u, nil
}

func (this *Model) GetUserByUsername(username string) (models.User, error) {
	u := models.User{}
	e := this.orm.QueryTable(&models.User{}).Filter("username", username).Limit(1).One(&u)

	if e == orm.ErrNoRows {
		return uzero, models.ErrResultNotFound
	}
	if e != nil {
		return uzero, e
	}

	return u, nil
}

func (this *Model) UpdateUser(nu models.User) (models.User, error) {
	// Check that user exist
	u, e := this.GetUserById(nu.Id)
	if e != nil {
		return uzero, e
	}

	//Check unique username constraint
	if nu.Username != u.Username {
		_, e = this.GetUserByUsername(u.Username)
		if e == nil {
			return uzero, models.ErrUserWithSameUsernameAlreadyExist
		}
		if e != models.ErrResultNotFound {
			return uzero, e
		}
	}

	_, e = this.orm.Update(&nu)

	if e != nil {
		return uzero, e
	}

	return nu, nil
}

func (this *Model) DeleteUser(id int) error {
	num, e := this.orm.Delete(&models.User{Id: id})
	if e == nil {
		if num == 0 {
			return models.ErrResultNotFound
		}
		return nil
	}
	return e
}

func (this *Model) GetUserQuerySeter() models.UserQuerySeter {
	return this.orm.QueryTable(&models.User{})
}
