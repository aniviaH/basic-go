package web

import (
	"errors"
	"fmt"
	"github.com/aniviaH/basic-go/webook/internal/domain"
	"github.com/aniviaH/basic-go/webook/internal/service"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
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
	//server.POST("/users/login", u.Login)
	server.POST("/users/login", u.LoginJWT)

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
	//server.GET("/users/profile", u.Profile)
	server.GET("/users/profile", u.ProfileJWT)

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

func (u *UserHandler) Login(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}
	// 拿到数据
	fmt.Printf("req: %v", req)
	user, err := u.svc.Login(ctx.Request.Context(), req.Email, req.Password)
	if errors.Is(err, service.ErrInvalidUserOrPassword) {
		ctx.String(http.StatusOK, "用户名或密码不对")
		return
	}

	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}

	//http.Cookie{
	//	Name:       "",
	//	Value:      "",
	//	Path:       "",
	//	Domain:     "",
	//	Expires:    time.Time{},
	//	RawExpires: "",
	//	MaxAge:     0,
	//	Secure:     false,
	//	HttpOnly:   false,
	//	SameSite:   0,
	//	Raw:        "",
	//	Unparsed:   nil,
	//}

	// 在这里登录成功了
	// session配置步骤2: 进行设置当前的session（通过http响应头部Set-Cookie将cookie告知客户端进行存储）
	sess := sessions.Default(ctx)
	// 设置放在 session 里面的值
	sess.Set("userId", user.Id)
	sess.Options(sessions.Options{
		//Secure:   true,
		//HttpOnly: true,
		//Path: "/products",
		// 60s过期(单位是s)
		//MaxAge: 60,
		// 一分钟过期
		MaxAge: 60 * 2,
	})
	// 需要调一下Save方法
	sessionSaveErr := sess.Save()
	if sessionSaveErr != nil {
		ctx.String(http.StatusInternalServerError, "登录失败")
		return
	}

	ctx.String(http.StatusOK, "登录成功")
	return
}

func (u *UserHandler) LoginJWT(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}
	// 拿到数据
	fmt.Printf("req: %v", req)
	user, err := u.svc.Login(ctx.Request.Context(), req.Email, req.Password)
	if errors.Is(err, service.ErrInvalidUserOrPassword) {
		ctx.String(http.StatusOK, "用户名或密码不对")
		return
	}

	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}

	// 步骤2
	// 在这里用 JWT 设置登录态
	// 生成一个 JWT token
	// 如果我要在 JWT token 里面带我个人数据，该怎么带？
	// 比如，我要带 userId

	claims := UserClaims{
		Uid: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
		},
	}
	//token := jwt.New(jwt.SigningMethodHS512)
	// token中保存用户数据claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenStr, err := token.SignedString([]byte("yICPpbp2QnPmCfHGEryXLXFtkCyEsela"))
	if err != nil {
		ctx.String(http.StatusInternalServerError, "系统错误 ")
		return
	}
	fmt.Println("tokenStr:", tokenStr)
	fmt.Println("userId:", user.Id)

	// 将token放到响应头，发给客户端
	ctx.Header("x-jwt-token", tokenStr)

	ctx.String(http.StatusOK, "登录成功")
	return
}

func (u *UserHandler) Logout(ctx *gin.Context) {
	sess := sessions.Default(ctx)
	sess.Options(sessions.Options{
		// 让session的cookie失效
		MaxAge: -1,
	})
	sessionSaveErr := sess.Save()

	if sessionSaveErr != nil {
		ctx.String(http.StatusInternalServerError, "退出登录失败")
		return
	}

	ctx.String(http.StatusOK, "退出登录成功")
	return
}

func (u *UserHandler) Edit(ctx *gin.Context) {

}

func (u *UserHandler) Profile(ctx *gin.Context) {
	ctx.String(http.StatusOK, "这是你的profile")
	return
}

func (u *UserHandler) ProfileJWT(ctx *gin.Context) {
	claims, _ := ctx.Get("claims")
	userId, _ := ctx.Get("userId")
	// 你可以断定，必然有 claims
	//if !exists1 {
	//	// 异常情况，可以在这里做监控
	//	fmt.Println("用户信息不存在")
	//	ctx.String(http.StatusInternalServerError, "用户信息不存在")
	//	return
	//}
	//if !exists2 {
	//	// 异常情况，可以在这里做监控
	//	fmt.Println("userId不存在")
	//	ctx.String(http.StatusInternalServerError, "userId不存在")
	//	return
	//}
	//fmt.Println(claims, userId)

	// ctx.Get 获取的返回值数据类型是 any，所以需要做下断言
	claimsAsset, ok := claims.(*UserClaims)
	if !ok {
		// 异常情况，可以在这里做监控
		fmt.Println("用户信息不存在")
		ctx.String(http.StatusInternalServerError, "用户信息不存在")
		return
	}
	fmt.Println(claimsAsset.Uid, userId)

	// 后面继续补充 Profile 的其它代码
	ctx.String(http.StatusOK, "这是你的profile")
	return
}

type UserClaims struct {
	// 组合，实现 jwt.Claims 接口
	jwt.RegisteredClaims
	// 声明你自己的要放进去 token 里面的数据
	Uid int64
	// 自己随便加
	// 一些敏感的信息不要放到里面
}
