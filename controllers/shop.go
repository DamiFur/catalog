package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCatalog(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
	return
}
