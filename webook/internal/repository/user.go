package repository

import (
	"context"
	"fmt"
	"github.com/aniviaH/basic-go/webook/internal/domain"
	"github.com/aniviaH/basic-go/webook/internal/repository/dao"
)

var ErrUserDuplicateEmail = dao.ErrUserDuplicateEmail
var ErrUserDuplicateEmailV1 = fmt.Errorf("%w 邮箱冲突", dao.ErrUserDuplicateEmail)

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) UserRepository {
	return UserRepository{
		dao: dao,
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

func (ur *UserRepository) FindById(int64) {
	// 先从 cache 里面找
	// 再从 dao 里面找
	// 找到了回写 cache
}
