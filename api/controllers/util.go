package controllers

import (
	"errors"

	"github.com/astaxie/beego/context"
	"github.com/mdiazp/sirel-server/api/models"
)

// GetAuthorFromInputData ...
func GetAuthorFromInputData(ctx *context.Context) (*models.User, error) {
	x := ctx.Input.Data()["Author"]
	if auth, ok := x.(*models.User); ok {
		return auth, nil
	}
	return nil, errors.New("Not user founded in ctx.Input.Data[\"Author\"]")
}

// PagAndOrdOptions ...
type PagAndOrdOptions struct {
	Limit          int
	Offset         int
	OrderBy        string
	orderDirection string
}
