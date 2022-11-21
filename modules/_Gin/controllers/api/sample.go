package api

import (
	"net/http"
	"project/modules/_Gin/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	// "log"
)

func InfoSample(c *gin.Context) {

	// ip := c.PostForm("xxx")	//get Front-End PostForm
	// DefaultQuery := c.DefaultQuery("first", "None") // ex. http://localhost/test?first=123 當 first 參數沒給時，帶入Default值None
	// Query := c.Query("first", "None") // ex. http://localhost/test?first=123 當 first 參數沒給時，等同沒有值

	cid := c.Param("id")         //get URL parameter /sample2/:id/*action
	caction := c.Param("action") //get URL parameter /sample2/:id/*action
	now := time.Now()
	returnData := sample{Status: 200, Data: cid + " | " + caction, Time: now.Format("2006-01-02 15:04:05 +0800")}

	c.JSON(http.StatusOK, returnData)
}

func AddSample(c *gin.Context) {
	num1, _ := strconv.Atoi(c.Param("num1"))
	num2, _ := strconv.Atoi(strings.TrimPrefix(c.Param("num2"), "/")) // <-- 注意:使用 * 方式取值，須將第一碼 "/" 刪除

	sum := models.Add(num1, num2)
	c.JSON(http.StatusOK, gin.H{"sum": sum, "input": (c.Param("num1") + "+" + strings.TrimPrefix(c.Param("num2"), "/"))})
}
