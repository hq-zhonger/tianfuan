package setting

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"tianfuan/HFish/core/protocol/mysql"
	"tianfuan/HFish/core/protocol/redis"
	"tianfuan/HFish/core/protocol/ssh"
	"tianfuan/HFish/utils/conf"
	"tianfuan/HFish/view"
	"time"
)

func RunWeb(template string, static string, url string) http.Handler {
	r := gin.New()
	r.Use(gin.Recovery())

	// 引入html资源
	r.LoadHTMLGlob("HFish/web/github/html/*")

	// 引入静态资源
	r.Static("HFish/static", "HFish/web/github/static")

	r.GET(url, func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	return r
}

func RunAdmin() http.Handler {
	gin.DisableConsoleColor()
	f, _ := os.Create("./hfish.log")
	gin.DefaultWriter = io.MultiWriter(f)
	// 引入gin
	r := gin.Default()

	r.Use(gin.Recovery())
	// 引入html资源
	r.LoadHTMLGlob("HFish/admin/*")

	// 引入静态资源
	r.Static("HFish/static", "HFish/static")

	// 加载路由
	view.LoadUrl(r)

	return r
}

func Run() {
	// 启动 Mysql 钓鱼
	mysqlStatus := conf.Get("mysql", "status")

	// 判断 Mysql 钓鱼 是否开启
	if mysqlStatus == "1" {
		mysqlAddr := conf.Get("mysql", "addr")

		// 利用 Mysql 服务端 任意文件读取漏洞
		mysqlFiles := conf.Get("mysql", "files")

		go mysql.Start(mysqlAddr, mysqlFiles)
	}

	//=========================//

	// 启动 Redis 钓鱼
	redisStatus := conf.Get("redis", "status")

	// 判断 Redis 钓鱼 是否开启
	if redisStatus == "1" {
		redisAddr := conf.Get("redis", "addr")
		go redis.Start(redisAddr)
	}

	//=========================//

	// 启动 SSH 钓鱼
	sshStatus := conf.Get("ssh", "status")

	// 判断 SSG 钓鱼 是否开启
	if sshStatus == "1" {
		sshAddr := conf.Get("ssh", "addr")
		go ssh.Start(sshAddr)
	}

	//=========================//

	// 启动 Web 钓鱼
	webStatus := conf.Get("web", "status")

	// 判断 Web 钓鱼 是否开启
	if webStatus == "1" {
		webAddr := conf.Get("web", "addr")
		webTemplate := conf.Get("web", "template")
		webStatic := conf.Get("web", "static")
		webUrl := conf.Get("web", "url")

		serverWeb := &http.Server{
			Addr:         webAddr,
			Handler:      RunWeb(webTemplate, webStatic, webUrl),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}

		go serverWeb.ListenAndServe()
	}

	//=========================//

	// 启动 admin 管理后台
	adminbAddr := conf.Get("admin", "addr")

	serverAdmin := &http.Server{
		Addr:         adminbAddr,
		Handler:      RunAdmin(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	serverAdmin.ListenAndServe()
}
