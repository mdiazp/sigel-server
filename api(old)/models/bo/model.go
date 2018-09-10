package bo

import (
	"database/sql"

	"github.com/astaxie/beego/orm"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

func init() {
	// if beego.BConfig.RunMode == "dev" {
	//	orm.Debug = true
	//	orm.DebugLog = orm.NewLog(os.Stdout)
	// }
	orm.RegisterModel(new(models.User))
	orm.RegisterModel(new(models.Area))
}

// Model implement models.Model interface using beego.orm
type Model struct {
	Db  *sql.DB
	orm orm.Ormer
}

func NewModel(db *sql.DB) (models.Model, error) {
	o, e := orm.NewOrmWithDB("postgres", "sirel", db)
	if e != nil {
		return nil, e
	}
	return &Model{
		Db:  db,
		orm: o,
	}, nil
}
