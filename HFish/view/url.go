package view

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tianfuan/HFish/utils/conf"
	"tianfuan/HFish/view/api"
	"tianfuan/HFish/view/dashboard"
	"tianfuan/HFish/view/fish"
	"tianfuan/HFish/view/login"
	"tianfuan/HFish/view/mail"
	"tianfuan/HFish/view/setting"
)

// 解决跨域问题
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func LoadUrl(r *gin.Engine) {
	// 登录
	r.GET("/login", login.Html)
	r.POST("/login", login.Login)
	r.GET("/logout", login.Logout)

	// 仪表盘
	r.GET("/", login.Jump, dashboard.Html)
	r.GET("/dashboard", login.Jump, dashboard.Html)
	r.GET("/get/dashboard/data", login.Jump, dashboard.GetFishData)

	// 钓鱼列表
	r.GET("/fish", login.Jump, fish.Html)
	r.GET("/get/fish/list", login.Jump, fish.GetFishList)
	r.GET("/get/fish/info", login.Jump, fish.GetFishInfo)
	r.POST("/post/fish/del", login.Jump, fish.PostFishDel)

	// 邮件群发
	r.GET("/mail", login.Jump, mail.Html)
	r.POST("/post/mail/sendEmail", login.Jump, mail.SendEmailToUsers)

	// 设置
	r.GET("/setting", login.Jump, setting.Html)
	r.GET("/get/setting/info", login.Jump, setting.GetSettingInfo)
	r.POST("/post/setting/update", login.Jump, setting.UpdateEmailInfo)
	r.POST("/post/setting/checkSetting", login.Jump, setting.UpdateStatusSetting)

	// API 接口
	// WEB 上报钓鱼信息
	apiStatus := conf.Get("api", "status")

	// 判断 API 是否启用
	if apiStatus == "1" {
		r.Use(cors())

		apiUrl := conf.Get("api", "url")
		r.POST(apiUrl, api.ReportWeb)
	}
}
