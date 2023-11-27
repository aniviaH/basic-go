package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrUserDuplicateEmail = errors.New("邮箱冲突")
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (ud *UserDAO) Insert(ctx context.Context, u User) error {
	// 时间的存储放dao来
	// 存毫秒数、存纳秒数 -> 一般存毫秒数就行
	// SELECT * FROM users where email=123@qq.com FOR UPDATE
	now := time.Now().UnixMilli()
	u.Ctime = now
	u.Utime = now
	err := ud.db.WithContext(ctx).Create(&u).Error // 注意，这里需要取u的地址
	// 下面这段代码与底层强耦合，就是针对底层使用的MYSQL数据库。如果不是，则也永远不会进入
	//if mysqlErr, ok := err.(*mysql.MySQLError); ok {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		const uniqueConflictsErrNo uint16 = 1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			// 邮箱冲突
			return ErrUserDuplicateEmail
		}
	}
	return nil
}

// User 在DAO里面，直接对应数据库表结构(一一对应)
// 有些人叫做 entity，有些人叫做 model，有些人叫做 PO(Persistent Object)，都是关于与数据库关联映射的概念
type User struct {
	Id int64 `gorm:"primaryKey,autoIncrement"`
	// 全部用户唯一
	Email    string `gorm:"unique"`
	Password string

	// 往这里面加其它的字段

	// 创建时间，毫秒数
	Ctime int64
	// 更新时间，毫秒数
	Utime int64
}

// UserDetail 用户详情表 - 可以用来放其它次要信息
type UserDetail struct {
}

// Address 地址信息表
type Address struct {
	Id     int64
	UserId int64
}
