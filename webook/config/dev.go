//go:build !k8s

// 没有 k8s 这个编译标签

// xxx go:build dev
// xxx go:build test
// xxx go:build e2e

package config

var Config = config{
	DB: DBConfig{
		// 本地的连接
		DSN: "root:root@tcp(localhost:13316)/webook",
	},
	Redis: RedisConfig{
		Addr: "localhost:6379",
	},
}
