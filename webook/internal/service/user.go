package service

import (
	"context"
	"github.com/aniviaH/basic-go/webook/internal/domain"
	"github.com/aniviaH/basic-go/webook/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// SignUp 服务层的注册含义
// service层的命名，一般保持和 Handler 那边的命名对应
func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
	// 第2个参数 domain.User 的原因：
	// 分层情况是：
	// hanlder
	// service
	// service 在 handler 的下层，保持链路的单一性，service 不应该去访问 handler，handler 可以访问 service
	// 所以针对用户概念的抽象，再通过在 domain 层做一层定义

	// 你要考虑加密放在哪里的问题了
	encrypted, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(encrypted)

	// 然后就是，存起来
	return svc.repo.Create(ctx, u)
}
