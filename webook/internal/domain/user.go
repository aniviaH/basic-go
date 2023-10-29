package domain

type User struct {
	//Addr Address
	Email    string
	Password string
	//ConfirmPassword string  service不需要，因为handler已经会处理这一块，confirmPassword更多是为了防止用户输入密码误操作的，service层更关系数据方面，所以这里不需要定义这个概念
}

//type Address struct {
//}
