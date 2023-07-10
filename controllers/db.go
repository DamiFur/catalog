package controllers

import (
	"net/http"

	"github.com/damifur/catalog/config"
	"github.com/damifur/catalog/services"
	"github.com/damifur/catalog/utils/errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetDBSession(c *gin.Context) {

	session, err := services.GetDBSession()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.InternalServer("Error while getting sql session", err))
		return
	}

	c.Set(config.DB, session)
	c.Next()
}

func GetDB(c *gin.Context) (*gorm.DB, *errors.Error) {
	db, ok := c.Get(config.DB)
	if !ok {
		return nil, errors.InternalServer("Couldn't retrieve database from user session")
	}
	return db.(*gorm.DB), nil
}
