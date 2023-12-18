package domain

import "time"

// User 领域对象，对应的是 DDD 中的聚合根，或者叫entity
// BO(Business Object)
type User struct {
	Id int64
	//Addr Address
	Email    string
	Password string
	//ConfirmPassword string  service不需要，因为handler已经会处理这一块，confirmPassword更多是为了防止用户输入密码误操作的，service层更关系数据方面，所以这里不需要定义这个概念
	// 昵称
	NickName string
	// 生日
	BirthDay time.Time
	// 个人简介
	PersonalDesc string
}

//func (u *User) Encrypt() string {
//	u.Password = encrypt(u.Password)
//	return u.Password
//}
//
//func (u *User) ComparePassword(input string) {
//
//}

//type Address struct {
//}
