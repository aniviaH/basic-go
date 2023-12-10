package web

import (
	"github.com/gin-gonic/gin"
)

//func RegisterRoutes() *gin.Engine {
//	server := gin.Default()
//
//	registerUserRoutes(server)
//
//	return server
//}

func registerUserRoutes(server *gin.Engine) {
	u := &UserHandler{}

	// 注册
	//server.POST("/users/signup", func(context *gin.Context) {
	//
	//})
	// 这是 RESTful 风格
	//server.PUT("/user", func(context *gin.Context) {
	//
	//})
	server.POST("/users/signup", u.Signup)

	// 登录
	//server.POST("/users/signin", func(context *gin.Context) {
	//
	//})
	server.POST("/users/signin", u.Login)

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

	//server.Run(":8080")
}
