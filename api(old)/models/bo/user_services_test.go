package bo

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

func opendb(t *testing.T) (models.Model, *sql.DB, error) {
	db, e := sql.Open("postgres", "user=sirel dbname=sirel password=123 sslmode=disable")
	if e != nil {
		t.Fatalf("Error opening database: %s", e.Error())
	}

	model, e := NewModel(db)
	if e != nil {
		t.Fatalf("Error creating model: %s", e)
	}

	return model, db, e
}

func getUserForTest(t *testing.T, model models.Model) models.User {
	u, e := model.GetUserByUsername("manuel.diaz")
	if e != nil {
		t.Fatalf("Error recovering user for test: %s", e.Error())
	}
	return u
}

func TestCreateUser(t *testing.T) {
	var (
		u models.User
		e error
	)

	model, db, e := opendb(t)

	//Test correct create user
	u, e = model.CreateUser(models.User{
		Id:       1,
		Username: "manuel.diaz",
	})
	if e != nil {
		t.Fatalf("Some error ocurred while create user for the first time: %s", e.Error())
	}

	uu, e := model.GetUserById(u.Id)
	if e != nil || u != uu {
		t.Fatalf("Error testing correct create user when call to GetUserById: %s", e.Error())
	}

	//Test Recover of uername unique constraint
	u, e = model.CreateUser(models.User{
		Username: "manuel.diaz",
	})
	if e == nil {
		t.Fatal("Error recovering username unique constraint: e is nil")
	}
	if e != models.ErrUserWithSameUsernameAlreadyExist {
		t.Fatalf("Error recovering username unique constraint: %s", e)
	}

	//Test Internal Error
	db.Close()
	u, e = model.CreateUser(models.User{
		Username: "manuel.diaz",
	})
	if !models.ErrNotImplementet(e) {
		t.Fatalf("Error testing Internal Error: %s", e.Error())
	}
}

func TestGetUserById(t *testing.T) {
	var (
		u models.User
		e error
	)

	model, db, e := opendb(t)
	uft := getUserForTest(t, model)

	// Test ErrUserNotFound
	u, e = model.GetUserById(10000)
	if e != models.ErrResultNotFound {
		if e == nil {
			t.Fatalf("Error testing ErrNoRow Id: e is nil")
		}
		t.Fatalf("Error testing ErrNoRow: %s", e.Error())
	}

	// Test Correct Id
	u, e = model.GetUserById(uft.Id)
	if e != nil {
		t.Fatalf("Error testing Correct Id: %s", e.Error())
	}
	if u != uft {
		t.Fatalf("Error testing Correct Id: diferent user")
	}

	//Test Internal Error
	db.Close()
	u, e = model.GetUserById(10000)
	if !models.ErrNotImplementet(e) {
		t.Fatalf("Error testing Internal Error: e=%s", e.Error())
	}
}

func TestGetUserByUsername(t *testing.T) {
	var (
		u models.User
		e error
	)

	model, db, e := opendb(t)
	uft := getUserForTest(t, model)

	// Test ErrNoRow
	u, e = model.GetUserByUsername("kokokoko")
	if e != models.ErrResultNotFound {
		if e == nil {
			t.Fatalf("Error testing ErrNoRow: error is nil")
		}
		t.Fatalf("Error testing ErrNoRow: %s", e.Error())
	}

	// Test Correct Username
	u, e = model.GetUserByUsername("manuel.diaz")
	if e != nil {
		t.Fatalf("Error testing Correct Username: %s", e.Error())
	}
	if u != uft {
		t.Fatalf("Error testing Correct Id: diferent user")
	}

	//Test Internal Error
	db.Close()
	u, e = model.GetUserByUsername("kokoko")
	if !models.ErrNotImplementet(e) {
		t.Fatalf("Error testing Internal Error: %s", e.Error())
	}
}

func TestUpdateUser(t *testing.T) {
	var (
		u models.User
		e error
	)

	model, db, e := opendb(t)
	uft := getUserForTest(t, model)

	// Test Correct update
	u, e = model.UpdateUser(models.User{
		Id:       uft.Id,
		Username: "manuel.diaz",
		Name:     "Manuel Alejandro Diaz Perez",
		Email:    "manuel.diaz@upr.edu.cu",
		Rol:      "Admin",
	})
	if e != nil {
		t.Fatalf("Error testing correct update: %s", e.Error())
	}
	uu, e := model.GetUserById(uft.Id)
	if e != nil {
		t.Fatalf("Error testing correct update when cal to GetUserById: %s", e.Error())
	}
	if u != uu {
		t.Fatal("Error testing correct update: User not updated")
	}

	// Test ErrUserWithSameUsernameAlreadyExist
	_, e = model.CreateUser(models.User{Username: "claupd"})
	if e != nil {
		t.Fatalf("Error creating new user in ErrUserWithSameUsernameAlreadyExist Test: %s", e.Error())
	}

	u.Username = "claupd"
	_, e = model.UpdateUser(u)
	if e != models.ErrUserWithSameUsernameAlreadyExist {
		kk := "e is nil"
		if e != nil {
			kk = e.Error()
		}
		t.Fatalf("Error testing ErrUserWithSameUsernameAlreadyExist: %s", kk)
	}

	// Test ErrNotImplementet
	db.Close()
	_, e = model.UpdateUser(u)
	if !models.ErrNotImplementet(e) {
		t.Fatalf("Error testing ErrNotImplementet: %s", e)
	}
}
func TestDeleteUser(t *testing.T) {
	var (
		u models.User
		e error
	)
	model, db, e := opendb(t)
	uft := getUserForTest(t, model)

	// Test correct delete user
	// // Deleting manuel.diaz
	u = uft
	e = model.DeleteUser(u.Id)
	if e != nil {
		t.Fatalf("Error testing correct delete user: %s", e.Error())
	}
	// // Deleting claupd
	u, e = model.GetUserByUsername("claupd")
	if e != nil {
		t.Fatalf("Error testing correct delete user when call to GetUserByUsername: %s", e.Error())
	}
	e = model.DeleteUser(u.Id)
	if e != nil {
		t.Fatalf("Error testing correct delete user: %s", e.Error())
	}

	// Test ErrUserNotFound
	e = model.DeleteUser(u.Id)
	if e != models.ErrResultNotFound {
		t.Fatalf("Error testing ErrUserNotFound: %s", e)
	}

	// Test ErrNotImplementet
	db.Close()
	e = model.DeleteUser(u.Id)
	if !models.ErrNotImplementet(e) {
		t.Fatalf("Error testing ErrNotImplement: %s", e.Error())
	}
}
