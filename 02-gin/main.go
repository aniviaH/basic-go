package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin: https://github.com/gin-gonic/gin

func main() {
	server := gin.Default()
	// 路由 当一个 HTTP 请求，用 GET 方法访问的时候，如果访问路径是 /hello
	server.GET("/hello", func(c *gin.Context) {
		//c.Writer.
		//c.Request.
		// 就执行这段代码
		c.String(http.StatusOK, "hello, go")
	})
	// 注册路由 对应 HTTP 的方法，其都有对应的方法
	server.POST("/post", func(context *gin.Context) {
		context.String(http.StatusOK, "hello, post 方法")
	})

	//上面的都是静态路由，路径需要完全匹配

	// 参数路由
	server.GET("/user/:name", func(context *gin.Context) {
		//context.String(http.StatusOK, "hello, 这是参数路由")

		name := context.Param("name")
		context.String(http.StatusOK, "hello, 这是参数路由，这是你传过来的名字：%s", name)
	})

	// 通配符路由
	server.GET("/views/*.html", func(context *gin.Context) {
		//context.String(http.StatusOK, "hello, 这是通配符路由")

		page := context.Param(".html")
		context.String(http.StatusOK, "hello, 这是通配符路由，这是你传过来的页面名字：%s", page)
	})

	// 查询参数 order?id=123
	server.GET("/order", func(ctx *gin.Context) {
		oid := ctx.Query("id")
		// It is shortcut for `c.Request.URL.Query().Get(key)`
		oid2 := ctx.Request.URL.Query().Get("id")
		ctx.String(http.StatusOK, "hello，这是查询参数，订单id=%s， 订单id2=%s", oid, oid2)
	})

	// 通配符路由不能注册这种 /users/*, /users/*/a。也就是说，* 不能单独出现。
	// 编译运行会报错
	//server.GET("/users/*/", func(ctx *gin.Context) {
	//})

	// ok
	//server.GET("/items/", func(ctx *gin.Context) {
	//	ctx.String(http.StatusOK, "hello, 这是items")
	//})

	server.GET("/items/*zhangsan", func(ctx *gin.Context) {
		matchText := ctx.Param("zhangsan")
		ctx.String(http.StatusOK, "hello, 这是items，匹配的到字符为%s", matchText)
	})

	//go func() {
	//	server1 := gin.Default()
	//	server1.GET("/hello2", func(c *gin.Context) {
	//		c.String(http.StatusOK, "hello2, go")
	//	})
	//	server1.Run(":8082")
	//}()

	server.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
