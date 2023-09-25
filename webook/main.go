package main

import (
	"github.com/aniviaH/basic-go/webook/internal/web"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func main() {
	/*
		server := gin.Default()
		u := &web.UserHandler{}

		// 注册
		//server.POST("/users/signup", func(context *gin.Context) {
		//
		//})
		// 这是 RESTful 风格
		//server.PUT("/user", func(context *gin.Context) {
		//
		//})
		server.POST("/users/signup", u.Gignup)

		// 登录
		//server.POST("/users/signin", func(context *gin.Context) {
		//
		//})
		server.POST("/users/signin", u.Signin)

		// 编辑用户
		//server.POST("/users/edit", func(context *gin.Context) {
		//
		//})
		// 这是 RESTful 风格
		//server.POST("/users/:id", func(context *gin.Context) {
		//
		//})
		server.POST("/users/edit", u.Edit)

		// 用户信息
		//server.GET("/users/profile", func(context *gin.Context) {
		//
		//})
		// REST 风格
		//server.GET("/users/:id", func(context *gin.Context) {
		//
		//})
		server.GET("/users/profile", u.Profile)

		server.Run(":8080")
	*/

	server := gin.Default()

	server.Use(func(context *gin.Context) {
		println("第一个 middleware")
	}, func(context *gin.Context) {
		println("第二个 middleware")
	})
	server.Use(func(context *gin.Context) {
		println("第三个 middleware")
	})

	// middleware方案：github.com/gin-gonic/contrib/gin-cors
	server.Use(cors.New(cors.Config{
		//出于安全考虑，这里不要用任意*号，公司里的域名个数一般都能容易列出来的。
		//另外前端xhr请求带了 withCredentials 属性时，也不能写*，否则会被浏览器认为跨域不通过而拦截
		//所以不要用*
		AllowOrigins:  []string{"http://localhost:3000"},
		AllowMethods:  []string{"POST", "GET", "OPTIONS"},        // 对应请求投头中的 Accecss-Control-Request-Method, 默认值是全部的simple methods
		AllowHeaders:  []string{"Content-Type", "authorization"}, // 对应请求投头中的 Accecss-Control-Request-Headers
		ExposeHeaders: []string{"Content-Length"},
		// 是否允许带 cookie 之类的东西
		AllowCredentials: true,
		// 如果 origin 判断逻辑复杂，可以用这个代替 AllowOrigins
		AllowOriginFunc: func(origin string) bool {
			//return origin == "https://github.com"
			if strings.HasPrefix(origin, "http://localhost") {
				// 开发环境
				return true
			}
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	//u := &web.UserHandler{}
	u := web.NewUserHanddler()
	u.RegisterRoutes(server)

	server.Run(":8080")
}
