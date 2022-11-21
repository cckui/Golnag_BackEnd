package route

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func authLogin() gin.HandlerFunc {

	return func(c *gin.Context) {

		session := sessions.Default(c)
		var loginSession string
		s := session.Get("login")

		switch s.(type) {
		case nil:
			loginSession = ""
		default:
			loginSession = fmt.Sprintf("%v", s)
		}

		if loginSession == "true" {
			// fmt.Println("通過驗證")
			c.Next()
		} else {
			// fmt.Println("驗證失敗")
			c.Redirect(302, "/")
			c.Abort()
		}
	}
}

func authMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"Status": 400,
			})
			c.Abort()
			// return
		} else {
			fmt.Println(token)
			c.Next()
		}
	}
}

func authBasic() gin.HandlerFunc {

	return func(c *gin.Context) {
		username, password, status := c.Request.BasicAuth()
		errMsg := "The authentication is failed"
		if !status {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": errMsg})
			return
		}
		if !validateUser(username, password) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": errMsg})
			return
		}
		c.Next()
	}
}
func validateUser(username string, password string) bool {
	// 驗證處理邏輯
	fmt.Printf("Account/Password： ")
	fmt.Printf(username + " / ")
	fmt.Println(password)
	return true
}

func ginAuthBasic() (auth gin.HandlerFunc) {

	auth = gin.BasicAuth(gin.Accounts{
		"user1": "user1",
		"user2": "user2",
		"user3": "user3",
	})

	return
}
