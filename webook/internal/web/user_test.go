package web

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	password := "hello#world123"
	// 加密
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}

	// 比较
	err = bcrypt.CompareHashAndPassword(encrypted, []byte(password))

	assert.NoError(t, err)
}

func TestNil(t *testing.T) {
	testTypeAssert(nil)
}

func testTypeAssert(c any) {
	// 类型断言，如果只接受一个返回值，那么类型断言不成立时候是会panic的
	// 两个返回值都接受时，断言不成立不会panic，可以对第二个返回值ok进行判断断言是否成立

	claims := c.(*UserClaims)
	println(claims)

	claims, ok := c.(*UserClaims)
	if !ok {
		fmt.Println("类型断言不成立")
		return
	}
	println(claims.Uid)
}
