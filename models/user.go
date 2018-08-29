package models

import (
	"log"
	"github.com/xormplus/xorm"
	"errors"
	"github.com/lwnmengjing/goAdminBackand/utils"
)

type User struct {
	Id 		  	int64
	UserName  	string		`xorm:"varchar(100) notnull 'username'"`
	FirstName 	string		`xorm:"varchar(100) 'first_name'"`
	LastName 	string		`xorm:"varchar(100) 'last_name'"`
	CreatedAt 	int			`xorm:"created"`
	UpdatedAt 	int			`xorm:"updated"`
}

//func init() {
//	if xrom, found := utils.Engin.GetXormEngin("default"); !found {
//		log.Println("Database default is not found")
//	} else {
//		xrom.Ping()
//	}
//
//}

func (u *User) TableName() string {
	return utils.Config.TablePrefix + "user"
}

func (u *User) Engin() (e *xorm.Engine, err error) {
	var found bool
	if e, found = utils.Engin.GetXormEngin("default"); !found {
		log.Println("Database default is not found")
		err = errors.New("Database default is not found")
	} else {
		e.Ping()
	}
	return
}

func (u *User)Insert() (err error) {
	engin, err := u.Engin()
	if err != nil {
		return
	}
	id, err := engin.Insert(u)
	if err != nil {
		return err
	}
	u.Id = id
	return
}