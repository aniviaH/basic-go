package web

import (
	"errors"
	"fmt"
	"github.com/aniviaH/basic-go/webook/internal/domain"
	"github.com/aniviaH/basic-go/webook/internal/service"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserHandler 我准备在它上面定义跟用户有关的路由
type UserHandler struct {
	svc *service.UserService
	//github.com/dlclark/regexp2
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	// 预编译正则表达式来提高校验速度。
	//return &UserHandler{
	//	emailExp:    regexp.MustCompile(emailRegexPattern, regexp.None),
	//	passwordExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
	//}

	const (
		emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	)
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	return &UserHandler{
		svc:         svc,
		emailExp:    emailExp,
		passwordExp: passwordExp,
	}
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
	server.POST("/users/signup", u.Signup)

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

	// 分组路由
	ug := server.Group("/users")
	ug.GET("/profile1", u.Profile) // /users/profile1
	ug.POST("/signup2", u.Signup)  // /users/signup2
}

func (u *UserHandler) Signup(ctx *gin.Context) {
	//ctx.String(http.StatusOK, "hello, 你在注册")
	type SignUpReq struct {
		// tag 字段的附带信息
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}

	var req SignUpReq
	// Bind方法会根据 Content-Type 来解析你的数据到 req 里面
	// 解析错了，就会直接写回一个 400 的错误
	if err := ctx.Bind(&req); err != nil {
		return
	}
	// 拿到数据
	fmt.Printf("req: %v", req)

	//const (
	//	emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	//	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	//)
	//ok, err = regexp.Match(emailRegexPattern, []byte(req.Email))
	// 使用 https://github.com/dlclark/regexp2
	//emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	//ok, err := emailExp.MatchString(req.Email)
	// 使用预编译
	ok, err := u.emailExp.MatchString(req.Email)
	if err != nil {
		// 你的正则表达式不对，才会出现error
		//ctx.String(http.StatusInternalServerError, "系统错误")
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		//ctx.String(http.StatusBadRequest, "你的邮箱格式不对")
		ctx.String(http.StatusOK, "你的邮箱格式不对")
		return
	}
	if req.ConfirmPassword != req.Password {
		ctx.String(http.StatusOK, "两次输入的密码不一致")
		return
	}
	//ok, err = regexp.Match(passwordRegexPattern, []byte(req.Password))
	// 使用 https://github.com/dlclark/regexp2
	//passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	//ok, err = passwordExp.MatchString(req.Password)
	// 使用预编译
	ok, err = u.passwordExp.MatchString(req.Password)
	if err != nil {
		// 记录日志
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "密码必须大于8位，包含数字、字母、特殊字符")
		return
	}

	// 这边就是数据库操作
	//fmt.Println("req:", req)

	// 考虑一下，能不能直接在 UserHanlder 里面操作数据库？
	// 不能。因为 Handler 只是负责和 HTTP 有关的东西。我们需要一个代表数据库抽象的东西。
	//db := gorm.Open()

	// 通过service层信息，调用一下 svc 的方法
	// 这里的第2个参数，用户，使用 domain.User，代表的是业务意义的User
	err = u.svc.SignUp(ctx.Request.Context(), domain.User{
		// 做转化
		Email:    req.Email,
		Password: req.Password,
	})
	//if err == service.ErrUserDuplicateEmail {
	if errors.Is(err, service.ErrUserDuplicateEmail) {
		ctx.String(http.StatusOK, "邮箱冲突")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统异常")
		return
	}

	ctx.String(http.StatusOK, "注册成功")
}

func (u *UserHandler) Signin(ctx *gin.Context) {
}

func (u *UserHandler) Edit(ctx *gin.Context) {

}

func (u *UserHandler) Profile(ctx *gin.Context) {

}
