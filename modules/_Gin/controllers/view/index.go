package web

import (
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var secrets = gin.H{
	"user1": gin.H{"email": "111@chen.com", "phone": "111111"},
	"user2": gin.H{"email": "222@chen.com", "phone": "222222"},
	"user3": gin.H{"email": "333@chen.com", "phone": "333333"},
}

func RedirectIndexSample(c *gin.Context) {

	c.Redirect(302, "/")
}

func IndexSample(c *gin.Context) {

	frontEndData := TodoPageData{
		PageTitle: []string{"Title1", "Title2", "Title3"},
		Todos: []Todo{
			{Title: "todo1", Done: true},
			{Title: "todo2", Done: false},
			{Title: "todo3", Done: true},
		},
	}

	session := sessions.Default(c)
	loginState := session.Get("loginState")
	loginTime := session.Get("loginTime")

	c.HTML(http.StatusOK, "index.sample", gin.H{
		"samplePageTitle": frontEndData.PageTitle,
		"sampleTodos":     frontEndData.Todos,
		"loginState":      loginState,
		"loginTime":       loginTime,
	})
}

func IndexLogin(c *gin.Context) {

	now := time.Now()

	session := sessions.Default(c)
	session.Set("loginState", true)
	session.Set("loginTime", now.Format("2006-01-02 15:04:05 +0800"))
	session.Save()

	user := c.MustGet(gin.AuthUserKey).(string) // <-- Gin提供的account撈取方式

	if secret, ok := secrets[user]; ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
	}

}

func IndexLogout(c *gin.Context) {

	session := sessions.Default(c)
	// session.Set("loginState", false)
	session.Delete("loginState")
	session.Save()
	session.Clear() // <-- 刪除整個session，選用非必須
	c.Redirect(302, "/")
}
