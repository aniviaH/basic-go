package middleware

import (
	"encoding/gob"
	"github.com/aniviaH/basic-go/webook/internal/web"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
	"time"
)

// LoginJWTMiddlewareBuilder JWT登录校验
type LoginJWTMiddlewareBuilder struct {
	paths []string
}

func NewLoginJWTMiddlewareBuilder() *LoginJWTMiddlewareBuilder {
	return &LoginJWTMiddlewareBuilder{}
}

func (l *LoginJWTMiddlewareBuilder) IgnorePaths(path string) *LoginJWTMiddlewareBuilder {
	l.paths = append(l.paths, path)
	return l
}

func (l *LoginJWTMiddlewareBuilder) Build() gin.HandlerFunc {
	// 用 Go 的方式编码解码
	gob.Register(time.Now())

	return func(ctx *gin.Context) {
		// 不需要登录校验的
		for _, path := range l.paths {
			if ctx.Request.URL.Path == path {
				return
			}
		}

		// 我现在用 JWT 来校验
		tokenHeader := ctx.GetHeader("Authorization")
		// 校验格式
		if tokenHeader == "" {
			// 没登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		//segs := strings.SplitN(tokenHeader, " ", 2)
		segs := strings.Split(tokenHeader, " ")
		if len(segs) != 2 {
			// 数据的格式不正确，有人瞎搞
			//ctx.AbortWithStatus(http.StatusBadRequest)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// 校验内容
		tokenStr := segs[1]
		//token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		//	return []byte("yICPpbp2QnPmCfHGEryXLXFtkCyEsela"), nil
		//})
		claims := &web.UserClaims{}
		// 从token中解析出claims。ParseWithClaims 里面一定要传入指针
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("yICPpbp2QnPmCfHGEryXLXFtkCyEsela"), nil
		})
		if err != nil {
			//ctx.AbortWithStatus(http.StatusInternalServerError)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		//if claims.ExpiresAt.Time.Before(time.Now()) {
		//	// 过期了
		//}
		// err 为 nil，token 不为 nil，Uid 不为 0
		if token == nil || !token.Valid || claims.Uid == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 安全措施：校验下用户的客户端
		if claims.UserAgent != ctx.Request.UserAgent() {
			// 严重的安全问题
			// 你是要监控的
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 每十秒钟刷新一次
		now := time.Now()
		if claims.ExpiresAt.Sub(now) < time.Second*50 {
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute))
			tokenStr, err = token.SignedString([]byte("yICPpbp2QnPmCfHGEryXLXFtkCyEsela"))
			if err != nil {
				// 记录日志
				log.Fatalln("jwt 续约失败", err)
			}
			ctx.Header("x-jwt-token", tokenStr)
		}

		// 通过登录校验...

		// 因为在这里已经对用户信息进行了解析，别的地方接口需要用到这些信息的，可以直接存一下到ctx，这样别的需要用的地方不用再重新解析一遍
		ctx.Set("claims", claims)
		ctx.Set("userId", claims.Uid)
	}
}
