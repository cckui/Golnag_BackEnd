package _Gin

import (
	"embed"
	"html/template"
	"net/http"
	route "project/modules/_Gin/routers"
	_GinZap "project/modules/_Ginzap"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

type ModuleCfgStruct struct {
	Mode string `json:"Mode"`
	IP   string `json:"Format"`
}

var ModuleCfg ModuleCfgStruct

func GinInit(fileList embed.FS) {
	//===== 設定Gin運行模式 Debug or Release
	switch ModuleCfg.Mode {
	case "d":
		gin.SetMode(gin.DebugMode)
	// case "r":
	// 	gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.ReleaseMode)
		// fmt.Println("** -m Error Input -> release(r) or debug(d)")
		// os.Exit(0)
	}

	ginServer := gin.New()
	ginMelody := melody.New()

	// ====== 使用 sessions,需開啟 ======
	memstore := memstore.NewStore([]byte("secret"))         // <-- 使用 memstore 當儲存媒介，secret是用於加密的金鑰
	ginServer.Use(sessions.Sessions("mysession", memstore)) // <-- mysession指的是session的名子，memstore可更換其他儲存系統 ex.cookie、memstore、redis、memcached、mongodb
	// =================================

	ginServer.Use(_GinZap.Ginzap()) // Log包細項設定
	ginServer.Use(gin.Recovery())   // Server crash recovery

	ginServer.StaticFS("/static", http.FS(fileList))
	ginServer.SetHTMLTemplate(initTemplates(&fileList)) // for release mode

	//===== 初始化路由部分，ginMelody for websocket 使用
	route.InitRoute(ginServer, ginMelody, &fileList)

	//===== IP Port
	// go ginServer.RunTLS(":443", "./certs/ssl.pem", "./certs/ssl.key") // <-- 走 https 方式
	ginServer.Run(ModuleCfg.IP)
}

func initTemplates(file *embed.FS) *template.Template {

	t := template.New("")
	nav, _ := file.ReadFile("views/nav.tpl")
	index, _ := file.ReadFile("views/index.html")

	t.New("index.sample").Parse(string(index) + string(nav))
	// t.New("network.html").Parse(string(network) + string(nav) + string(sidebar) + string(footer))

	return t
}
