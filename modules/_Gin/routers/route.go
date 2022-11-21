package route

import (
	"embed"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

type Router struct {
	Root *gin.Engine
	// Helloworld *gin.RouterGroup
	// Stats      *gin.RouterGroup
	Routers *gin.RouterGroup

	Melody *melody.Melody

	FileList *embed.FS
}

var BaseRouter *Router

func InitRoute(router *gin.Engine, m *melody.Melody, file *embed.FS) {
	BaseRouter := &Router{Root: router, Melody: m, FileList: file}

	// >>> 指定Template
	// BaseRouter.Root.HTMLRender = templateRender()

	// BaseRouter.Root.SetHTMLTemplate(initTemplates()) // for release mode

	// >>> 指定靜態資源Mapping路徑
	//BaseRouter.Root.Static("/static", "./public")
	// BaseRouter.Root.Static("/assets", "./assets") // <-- Example
	// BaseRouter.Root.StaticFS("/more_static", http.Dir("my_file_system")) <-- Example
	// BaseRouter.Root.StaticFile("/favicon.ico", "./resources/favicon.ico") <-- Example

	// >>> 指定網站Logo
	//BaseRouter.Root.StaticFile("/favicon.ico", "./public/favicon.ico")

	// >>> API
	BaseRouter.InitInfoAPI()

	// >>> Website
	BaseRouter.InitWeb()

	// >>> WebSocket
	// BaseRouter.InitWS()

	return
}

// func templateRender() multitemplate.Renderer {
// 	r := multitemplate.NewRenderer()
// 	r.AddFromFiles("index", "views/index.html")
// 	// r.AddFromFiles("dashboard", "views/dashboard.html","views/_footer.tpl","views/_nav.tpl", "views/_sidebar.tpl" )
// 	// r.AddFromFiles("agv_list_tpl", "views/agv_list_tpl.tpl")
// 	return r
// }
