package main

import (
	"fmt"
	"github.com/aniviaH/basic-go/webook/internal/repository"
	"github.com/aniviaH/basic-go/webook/internal/repository/dao"
	"github.com/aniviaH/basic-go/webook/internal/service"
	"github.com/aniviaH/basic-go/webook/internal/web"
	"github.com/aniviaH/basic-go/webook/internal/web/middleware"
	"github.com/aniviaH/basic-go/webook/pkg/ginx/middlewares/ratelimit"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

func RangeKeyPairs(keyPairs ...[]byte) int {
	for k, v := range keyPairs {
		fmt.Println(k, v)
	}
	//for i := 0; i < len(keyPairs); i++ {
	//	fmt.Println(keyPairs[i])
	//}
	return 999
}
func Test(keyPairs ...[]byte) int {
	return RangeKeyPairs(keyPairs...)
}

func main() {
	//mainWebook()

	mainBase()
}

func mainBase() {
	server := gin.Default()

	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello, world 2023-12-30 18:41")
	})

	server.Run(":8080")
}

func mainWebook() {
	Test([]byte("abcde"), []byte("hjk"))

	//defaultLogicComment()

	// 初始化db
	db := initDb()

	// 初始化server
	server := initWebServer()

	// 初始化user的handler
	uh := initUser(db)

	// 注册user的路由
	uh.RegisterRoutes(server)

	// 启动服务
	server.Run(":8080")
}

func initWebServer() *gin.Engine {
	server := gin.Default()

	server.Use(func(context *gin.Context) {
		println("第一个 middleware")
	}, func(context *gin.Context) {
		println("第二个 middleware")
	})

	server.Use(func(context *gin.Context) {
		println("第三个 middleware")
	})

	// 基于redis的限流插件(使用redis进行统计判断是否达到限流的条件)
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	server.Use(ratelimit.NewBuilder(redisClient, time.Second, 100).Build())

	// middleware方案：github.com/gin-gonic/contrib/gin-cors
	server.Use(cors.New(cors.Config{
		//出于安全考虑，这里不要用任意*号，公司里的域名个数一般都能容易列出来的。
		//另外前端xhr请求带了 withCredentials 属性时，也不能写*，否则会被浏览器认为跨域不通过而拦截
		//所以不要用*
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"POST", "GET", "OPTIONS"},        // 对应请求投头中的 Accecss-Control-Request-Method, 默认值是全部的simple methods
		AllowHeaders: []string{"Content-Type", "Authorization"}, // 对应请求投头中的 Accecss-Control-Request-Headers
		// 这里控制哪些自定义设置的响应头可以被客户端获取，未在这个列表里的前端不能拿到
		ExposeHeaders: []string{"Content-Length", "x-jwt-token"},
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

	// session配置步骤1：初始化gin sessions的配置
	//store := cookie.NewStore([]byte("secret"))

	// 基于内存的实现，第一个参数是 authentication key，最好是 32 或者 64 位
	// 第二个参数是 encryption key
	//store := memstore.NewStore([]byte("authentication-key-123456"), []byte("encryption-key-123456"))
	store := memstore.NewStore([]byte("cBjJFkt0Kgs6CKD4cr6QYYd8qIaQi8pds7Py3kYEkibIzjf1hRFe3EnLoCfhk2BI"), []byte("yICPpbp2QnPmCfHGEryXLXFtkCyEsela"))

	// 基于redis的实现
	// 第一个参数是最大空闲连接数量
	// 第二个参数是网络类型：tcp 或 udp，你不太可能用udp
	// 第三个参数是host:端口
	// 第四哥参数是密码
	// 最后两个就是两个key
	//store, err := redis.NewStore(16, "tcp", "localhost:6379", "", []byte("cBjJFkt0Kgs6CKD4cr6QYYd8qIaQi8pds7Py3kYEkibIzjf1hRFe3EnLoCfhk2BI"), []byte("yICPpbp2QnPmCfHGEryXLXFtkCyEsela"))
	//if err != nil {
	//	panic(err)
	//}

	// cookie的名字叫mysession
	server.Use(sessions.Sessions("mysession", store))

	// 面向接口编程，可以自己提供一个sqlx版本的sessions实现
	//mystore := &sqlx_store.Store{}
	//server.Use(sessions.Sessions("mysession", mystore))

	// session配置步骤3: 对路由访问进行登session的校验和拦截（封装出去作为一个中间件函数）
	// v1
	//middleware.IgnorePaths = []string{"users/login", "users/signup"}
	//middleware.IgnorePaths = []string{"sss"}
	//server.Use(middleware.CheckLogin())
	// 问题点是，如果我有两个server，这个服务不能忽略sss这条路径。这个时候v1就做不到，因为内部是包变量，其只能有一个
	// 虽然说，登录的场景不会有两个server这种场景。但开发很多其它middlerware时是可能遇到这种场景的。
	//server1 := gin.Default()
	//server1.Use(middleware.CheckLogin())
	// v2
	//server.Use(middleware.CheckLoginV2([]string{"sss"}, "abc", "def"))
	//server1 := gin.Default()
	//server1.Use(middleware.CheckLoginV2([]string{}, "abc", "def"))
	// v3: 推荐写法
	// 作为中间件的提供者，如果你的设计有问题需要修改，那么使用你的中间件用户都得进行更新，这会是影响很大。所以中间件设计之初应该考虑好兼容性和扩展性
	// 如下的写法，它可以做到比较好的兼容性和扩展性
	//server.Use(
	//	middleware.NewLoginMiddlewareBuilder().
	//		IgnorePaths("/users/login").
	//		IgnorePaths("/users/signup").
	//		Build())
	// 使用JWT的登录校验middleware拦截器
	server.Use(middleware.NewLoginJWTMiddlewareBuilder().
		IgnorePaths("/users/login").
		IgnorePaths("/users/signup").
		Build())
	// 校验登录态 - 可以封装一下
	//server.Use(func(ctx *gin.Context) {
	//	// 不需要登录校验的路由
	//	if ctx.Request.URL.Path == "/users/login" || ctx.Request.URL.Path == "/users/signup" {
	//		return
	//	}
	//
	//	sess := sessions.Default(ctx)
	//	id := sess.Get("userId")
	//	if id == nil {
	//		// 没有登录
	//		ctx.AbortWithStatus(http.StatusUnauthorized)
	//		return
	//	}
	//})

	return server
}

func initUser(db *gorm.DB) *web.UserHandler {
	userDAO := dao.NewUserDAO(db)
	userRepo := repository.NewUserRepository(userDAO)
	userService := service.NewUserService(userRepo)
	//u := &web.UserHandler{}
	uh := web.NewUserHandler(userService)

	return uh
}

func initDb() *gorm.DB {
	// 初始化db
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		// mysql启动异常，直接panic，将当前goroutine直接结束
		// 只在初始化过程中panic，相当于整个 goroutine 结束
		// 一旦数据库连接初始化过程出错，应用就不要启动了
		panic(err)
	}

	// 初始化table
	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}

	return db
}

func defaultLogicComment() {
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
}
