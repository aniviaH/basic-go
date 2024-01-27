package service

import (
	"context"
	"errors"
	"github.com/aniviaH/basic-go/webook/internal/domain"
	"github.com/aniviaH/basic-go/webook/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserDuplicateEmail    = repository.ErrUserDuplicateEmail
	ErrInvalidUserOrPassword = errors.New("账号/邮箱或密码不对")
)

type UserService struct {
	repo repository.UserRepository
	//redis *redis.Client
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

	//err := svc.repo.Create(ctx, u)
	//if err != nil {
	//	return err
	//}
	// 存redis，
	//val, err := json.Marshal(u)
	//if err != nil {
	//	return err
	//}
	// 注意这边，要求 u 的id 不为 0
	//err = svc.redis.Set(ctx, fmt.Sprintf("user:info:%d", u.Id), val, time.Minute*30).Err()
	//return err
}

func (svc *UserService) Login(ctx context.Context, email string, password string) (domain.User, error) {
	// 先找用户
	u, err := svc.repo.FindByEmail(ctx, email)
	//if err == repository.ErrUserNotFound {
	//	return ErrInvalidUserOrPassword
	//}
	if errors.Is(err, repository.ErrUserNotFound) {
		// 账号或密码不存在
		return domain.User{}, ErrInvalidUserOrPassword
	}
	if err != nil {
		// 其它错误
		return domain.User{}, err
	}

	// 比较密码
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		// 可以在这里打个日志，比如DEBUG日志
		return domain.User{}, ErrInvalidUserOrPassword
	}
	//if u.Password != password {
	//	return domain.User{}, ErrInvalidUserOrPassword
	//}
	return u, err
}

func (svc *UserService) Profile(ctx context.Context, id int64) (domain.User, error) {
	// 第一个念头
	//val, err := svc.redis.Get(ctx, fmt.Sprintf("user:info:%d", id)).Result()
	//if err != nil {
	//	return domain.User{}, err
	//}
	//var u domain.User
	//err = json.Unmarshal([]byte(val), &u)
	//if err == nil {
	//	return u, err
	//}

	u, err := svc.repo.FindById(ctx, id)
	return u, err

	// 接下来，就是从数据库里面查找
	return domain.User{}, nil
}
