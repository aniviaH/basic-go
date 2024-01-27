package repository

import (
	"context"
	"github.com/aniviaH/basic-go/webook/internal/domain"
	"github.com/aniviaH/basic-go/webook/internal/repository/cache"
	"github.com/aniviaH/basic-go/webook/internal/repository/dao"
)

var (
	ErrUserDuplicateEmail = dao.ErrUserDuplicateEmail
	ErrUserNotFound       = dao.ErrUserNotFound
)

type UserRepository struct {
	dao *dao.UserDAO
	//redis *redis.Client
	cache *cache.UserCache
}

func NewUserRepository(dao *dao.UserDAO, c *cache.UserCache) UserRepository {
	return UserRepository{
		dao:   dao,
		cache: c,
	}
}

func (ur *UserRepository) Create(ctx context.Context, u domain.User) error {
	// 数据怎么存
	// 方式1： 存数据库
	return ur.dao.Insert(ctx, dao.User{
		// 转化一下
		Email:    u.Email,
		Password: u.Password,
		Ctime:    0,
		Utime:    0,
	})

	// 如果有缓存操作，就在这里进行操作...
}

func (ur *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := ur.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
	}, nil
}

func (ur *UserRepository) FindById(ctx context.Context, id int64) (domain.User, error) {
	// 先从 cache 里面找
	u, err := ur.cache.Get(ctx, id)
	if err == nil {
		// 必然是有数据
		return u, nil
	}
	// 没这个数据
	if err == cache.ErrKeyNotExist {
		// 去数据库里面加载
	}
	// 这里怎么办？ err = io.EOF
	// e: entity
	ue, err := ur.dao.FindById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	u = domain.User{
		Id:       ue.Id,
		Email:    ue.Email,
		Password: ue.Password,
	}

	// 存到userCache，这里可以开一个go routine，也可以不开
	go func() {
		err = ur.cache.Set(ctx, u)
		if err != nil {
			// 我这里怎么办？
			// 返回错误？
			//return domain.User{}, err
			// 还是不管？

			// 打日志，做监控
		}
	}()

	return u, nil
	// 要不要去数据库加载
	// 看起来我不应该加载？
	// 看起来我好像也要加载？

	// 选加载 -- 做好兜底，万一 Redis 真的崩了。你要保护住你的数据库
	// 我数据库限流呀！

	// 选不加载 -- 用户体验差一点

	// 用缓存始终要面对考虑的两个问题：
	// 1. 一致性问题
	// 2. 缓存崩了

	// 几种场景
	// 1. 缓存里面有数据
	// 2. 缓存里面没有数据
	// 3. 缓存出错了，你也不知道有没有数据

	// 再从 dao 里面找
	// 找到了回写 cache
}
