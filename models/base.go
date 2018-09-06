package models

import (
	"github.com/xormplus/xorm"
	"errors"
	"github.com/lwnmengjing/goAdminBackand/utils"
)

type ActiveRecod interface {
	GetDb()	(e *xorm.Engine, err error)
	TableName() string
}

type Base struct {

}

func (u *Base) TableName() string {
	return utils.Config.TablePrefix + "user"
}

func (u *Base) GetDb() (e *xorm.Engine, err error) {
	var found bool
	if e, found = utils.Engin.GetXormEngin("default"); !found {
		err = errors.New("Database default is not found!")
	}
	return
}
