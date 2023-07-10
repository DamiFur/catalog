package services

import (
	"github.com/damifur/catalog/dao"
	"gorm.io/gorm"
)

func GetDBSession() (*gorm.DB, error) {
	return dao.GetSession()
}
