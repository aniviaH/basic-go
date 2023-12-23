package middleware

import (
	"encoding/gob"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LoginMiddlewareBuilder struct {
	paths []string
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}

func (l *LoginMiddlewareBuilder) IgnorePaths(path string) *LoginMiddlewareBuilder {
	l.paths = append(l.paths, path)
	return l
}

func (l *LoginMiddlewareBuilder) Build() gin.HandlerFunc {
	// 用 Go 的方式编码解码
	gob.Register(time.Now())

	return func(ctx *gin.Context) {
		// 不需要登录校验的路由
		//if ctx.Request.URL.Path == "/users/login" || ctx.Request.URL.Path == "/users/signup" {
		//	return
		//}
		for _, path := range l.paths {
			if ctx.Request.URL.Path == path {
				return
			}
		}

		sess := sessions.Default(ctx)
		id := sess.Get("userId")
		fmt.Println("sess id", id)

		if id == nil {
			// 没有登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		//now := time.Now().UnixMilli()
		//// 可以在这里做session的刷新，那我怎么知道，一分钟已经过去了？
		//// 先拿到上次的更新时间
		//updateTime := sess.Get("update_time")
		//if updateTime == nil {
		//	// 说明还没有刷新过，刚登录，还没刷新过
		//	sess.Set("update_time", now)
		//	sess.Save()
		//	return
		//}
		//// updateTime 是有的
		//updateTimeVal, ok := updateTime.(int64) // 这里断言一下自己的类型
		//if !ok {
		//	// 非系统正常行为，正常不会出现
		//	ctx.AbortWithStatus(http.StatusInternalServerError)
		//	return
		//}
		//if now-updateTimeVal > 60*1000 {
		//	// 固定间隔时间刷新，比如每分钟内第一次访问刷新。
		//	sess.Set("update_time", now)
		//	sess.Save()
		//	return
		//}

		// 登录校验通过之后，可以在这里做session的刷新，那我怎么知道，一分钟已经过去了？
		// 先拿到上次的更新时间
		updateTime := sess.Get("update_time")
		sess.Set("userId", id)
		sess.Options(sessions.Options{
			MaxAge: 60 * 2,
		})
		now := time.Now()
		if updateTime == nil {
			// 说明还没有刷新过，刚登录，还没刷新过
			sess.Set("update_time", now)
			err := sess.Save()
			if err != nil {
				//panic(err)
				ctx.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			return
		}
		// updateTime 是有的
		updateTimeVal, ok := updateTime.(time.Time) // 这里断言一下自己的类型
		if !ok {
			// 非系统正常行为，正常不会出现
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if now.Sub(updateTimeVal) > time.Second*10 {
			// 固定间隔时间刷新，比如每分钟内第一次访问刷新。这里time.Second*10=10s刷新一次
			sess.Set("update_time", now)
			sess.Save()
			return
		}
	}
}

var IgnorePaths []string

func CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 不需要登录校验的路由
		//if ctx.Request.URL.Path == "/users/login" || ctx.Request.URL.Path == "/users/signup" {
		//	return
		//}
		for _, path := range IgnorePaths {
			if ctx.Request.URL.Path == path {
				return
			}
		}

		sess := sessions.Default(ctx)
		id := sess.Get("userId")
		fmt.Println("sess id", id)
		if id == nil {
			// 没有登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

// CheckLoginV2 这种写法可以处理自定义不需要登录的路径，但调用时候就是需要传入参数
func CheckLoginV2(paths []string, abc string, def string) gin.HandlerFunc {
	if len(paths) == 0 {
		paths = []string{}
	}
	return func(ctx *gin.Context) {
		// 不需要登录校验的路由
		//if ctx.Request.URL.Path == "/users/login" || ctx.Request.URL.Path == "/users/signup" {
		//	return
		//}
		for _, path := range paths {
			if ctx.Request.URL.Path == path {
				return
			}
		}

		sess := sessions.Default(ctx)
		id := sess.Get("userId")
		fmt.Println("sess id", id)
		if id == nil {
			// 没有登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
