package app

import (
	"html/template"

	"github.com/damifur/catalog/dao"
	"github.com/gin-gonic/gin"
)

var (
	engine *gin.Engine
)

func StartApp() {

	dao.CreateDatabase("prod")
	//defer services.DBClose()

	newEngine()
	engine.Run(":80")

}

func newEngine() {
	engine = gin.New()
	//engine.Use(gin.RecoveryWithWriter(logger.GetOut()))
	//store := sessions.NewCookieStore([]byte("a"))
	//engine.Use(sessions.Sessions("spotifysession", store))
	// This is in case we need to check that the user is logged (for security reasons) before handling a page
	//engine.Use(auth.SpotifyAuthFromConfig())
	loadHTMLGlob(engine, "static/templates/*.html", template.FuncMap{} /* If you want to use functions with the templates they should be passed here */)
	engine.Static("/static", "./static")
	//engine.NoRoute(controllers.RouteNotFound)

	mapURLSToControllers()

}
