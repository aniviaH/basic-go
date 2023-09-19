package web

import "github.com/gin-gonic/gin"

// UserHandler 我准备在它上面定义跟用户有关的路由
type UserHandler struct {
}

//func (u UserHandler) RegisterRoutes(server *gin.Engine) {
//
//}

type ArticalHandler struct {
}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
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

	//server.Run(":8080")
}

func (u *UserHandler) Gignup(ctx *gin.Context) {

}

func (u *UserHandler) Signin(ctx *gin.Context) {

}

func (u *UserHandler) Edit(ctx *gin.Context) {

}

func (u *UserHandler) Profile(ctx *gin.Context) {

}
