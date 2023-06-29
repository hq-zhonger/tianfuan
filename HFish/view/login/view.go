package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tianfuan/HFish/error"
	"tianfuan/HFish/utils/conf"
)

func Html(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func Jump(c *gin.Context) {
	account := conf.Get("admin", "account")
	loginCookie, _ := c.Cookie("is_login")
	if account != loginCookie {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
		return
	} else {
		c.Next()
	}
}

func Login(c *gin.Context) {
	loginName := c.PostForm("loginName")
	loginPwd := c.PostForm("loginPwd")

	account := conf.Get("admin", "account")
	password := conf.Get("admin", "password")

	if loginName == account {
		if loginPwd == password {
			c.SetCookie("is_login", loginName, 60*60*24, "/", "*", false, true)
			c.JSON(http.StatusOK, error.ErrSuccessNull())
			return
		}
	}

	c.JSON(http.StatusOK, error.ErrLoginFail())
}

func Logout(c *gin.Context) {
	c.SetCookie("is_login", "", -1, "/", "*", false, true)
	c.Redirect(http.StatusFound, "/login")
}
