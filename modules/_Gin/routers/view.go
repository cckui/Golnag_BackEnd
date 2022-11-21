package route

import (
	"net/http"
	web "project/modules/_Gin/controllers/view"

	"github.com/gin-gonic/gin"
)

func (r *Router) InitWeb() {
	r.Routers = r.Root.Group("/")

	r.Routers.GET("/logout", web.IndexLogout)
	r.Routers.GET("/", web.IndexSample)
	r.Routers.GET("/dashboard", web.RedirectIndexSample)
	r.Routers.GET("favicon.ico", func(c *gin.Context) {
		file, _ := r.FileList.ReadFile("public/favicon.ico")
		c.Data(
			http.StatusOK,
			"image/x-icon",
			file,
		)
	})
	r.Routers.Use(ginAuthBasic()) // 加入這行後以下都需驗證，以上不需驗證
	r.Routers.POST("/login", web.IndexLogin)
	r.Routers.GET("/login", web.IndexLogin)
	// r.Routers.POST("/test/:id", api.xxx)
}
