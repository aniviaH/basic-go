package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aniviaH/basic-go/webook/internal/domain"
	"github.com/redis/go-redis/v9"
	"time"
)

// var ErrUserNotFound =
var ErrKeyNotExist = redis.Nil

// 较差的方案
//type Cache interface {
//	GetUser(ctx context.Context, id int64) (domain.User, error)
//
//	// 读取文章
//	GetArticle(ctx context.Context, aid int64)
//
//	// 还有别的业务
//	// 。。。
//}
//type CacheV1 interface {
//	// 你的中间件团队去做的(统一的cache方案-cache api，或者有的会称为unify cache（统一缓存）)
//	Get(ctx context.Context, key string) (any, error)
//}

// 最理想的方案
//type UserCache struct {
//	cache CacheV1
//}
//
//func (u *UserCache) GetUser(ctx context.Context, id int64) (domain.User, error) {
//	return domain.User{}, nil
//}

// 正常方案

type UserCache struct {
	// 很差的写法
	//client        *redis.Client
	//clusterClient *redis.ClusterClient

	// 面向接口编程 -> 传单机的 Redis 可以，传 cluster 的 Redis 也可以，这里并不关心，只要你实现了我期望的那些接口
	client     redis.Cmdable
	expiration time.Duration
}

// 依赖注入：我要用client，我绝对不会自己去初始化它，我要求的是外面传进来
/**
原则：
	A 用到了 B，B 一定是接口
	A 用到了 B，B 一定是 A 的字段
	A 用到了 B，A 绝对不初始化 B，而是外面注入
*/

// NewUserCache 就这样写!!! 由外部传进来client(外部初始化好)，它爱怎么初始化就怎么初始化，想用什么参数就用什么参数，你这边根本一点不care。你这边就非常轻量，就是一层皮，不做其它任何事
func NewUserCache(client redis.Cmdable, expiration time.Duration) *UserCache {
	// 业务专门的cache，我这个UserCache就是专门处理User业务的
	return &UserCache{
		client:     client,
		expiration: time.Minute * 15,
		//expiration: expiration,
	}
}

func (uc *UserCache) GetUser(ctx context.Context, id int64) (domain.User, error) {
	return domain.User{}, nil
}

func (uc *UserCache) Key(id int64) string {
	// user:info:123
	// user_info_123
	// bumen_xiaozu_user_info_key
	return fmt.Sprintf("user:info:%d", id)
}

func (uc *UserCache) Set(ctx context.Context, user domain.User) error {
	// redis存储的数据需要进行 序列化

	// 不要忽略这里的问题，方便序列化出错时问题定位
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	key := uc.Key(user.Id)

	return uc.client.Set(ctx, key, data, uc.expiration).Err()
}

// 只要 error 为 nil，就认为 user 一定在，缓存里面一定有数据
// 如果没有数据，返回一个特定的 error
// 缓存里面有没有数据或其它问题，只有你能告诉我
func (uc *UserCache) Get(ctx context.Context, id int64) (domain.User, error) {
	key := uc.Key(id)
	// 数据不存在，err = redis.nil
	data, err := uc.client.Get(ctx, key).Bytes()
	if err != nil {
		return domain.User{}, err
	}

	var u domain.User
	err = json.Unmarshal(data, &u)
	if err != nil {
		return domain.User{}, err
	}
	return u, nil
}

// 不要这样写!!!
func NewUserCache2(cfg Config) *UserCache {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.Addr,
		// 可能还要初始化其它的字段，那么Config类型里就要对应定义对应的，最终可能就会形成Config与redis.Options一模一样，但这是没有必要的
		// 既然是接收外部的，那直接让外部初始化client直接传进来就行了。内部不要自己去初始化
	})
	return &UserCache{
		client: client,
	}
}

// 这样写更差!!!(万恶之源)
//func NewUserCache3() *UserCache {
//	client := os.LookupEnv("REDIS_ADDR")
//	return &UserCache{
//		client: client,
//	}
//}

type Config struct {
	Addr string
	// 增加其它字段定义，又要与redis.Options对应
	// 。。。
}

//func (u *UserCache) GetUser(ctx context.Context, id int64) (domain.User, error) {
//	if u.client == nil {
//		return u.clusterClient.xx
//	}
//	//
//}
