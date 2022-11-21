package route

import api "project/modules/_Gin/controllers/api"

func (r *Router) InitInfoAPI() {
	r.Routers = r.Root.Group("/api")
	r.Routers.GET("/add/:num1/*num2", api.AddSample)
	r.Routers.GET("/sample1/:id", api.InfoSample)
	r.Routers.Use(authBasic()) // 加入這行後以下都需驗證，以上不需驗證
	r.Routers.GET("/sample2/:id/*action", api.InfoSample)
	// r.Routers.POST("/test/:id/*action", api.xxx) // <-- 裡面的 :id/*action， : 跟 * 的差異在於，: 如果沒給參數的話會 404，* 可以不給沒關係，*方式只能放在最後．
}
