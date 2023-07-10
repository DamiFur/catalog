package app

import (
	"html/template"

	"github.com/damifur/catalog/controllers"
	"github.com/gin-gonic/gin"
	"github.com/leekchan/gtf"
)

func mapURLSToControllers() {

	withCtx := engine.Group("/", controllers.GetDBSession)

	withCtx.GET("", controllers.GetCatalog)
	withCtx.GET("ap", controllers.GetAdminPanel)
	withCtx.POST("c", controllers.CreateCategory)
	withCtx.POST("i", controllers.CreateItem)
	withCtx.POST("upload", controllers.UploadImage)
	//engine.GET("spotify", controllers.GetMainPage)
	//engine.GET("login", auth.Login)
	//engine.GET("logout", auth.Logout)
	//engine.GET("callback", auth.HandleOAuth2Callback)
	//engine.POST("start", controllers.StartTimer)
	//engine.PUT("addfive", controllers.Add5ToTimer)
	//engine.PUT("subfive", controllers.SubstractToTimer)
	//engine.PUT("cancel", controllers.CancelTimer)

}

// Modified code from gin.LoadHTMLGlob
func loadHTMLGlob(engine *gin.Engine, pattern string, funcs template.FuncMap) {
	gtf.Inject(funcs)
	templ, err := template.New("").Funcs(funcs).ParseGlob(pattern)
	if err != nil {
		templ = template.Must(template.New("").Funcs(funcs).ParseGlob("../" + pattern))
	}
	engine.SetHTMLTemplate(templ)
}
