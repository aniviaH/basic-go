package main

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type number int

type Integer int

// 泛型约束
type Number interface {
	~int | uint | int64 | float64
	// ~int 代表泛型约束为 int及int的衍生类型
}

func Sum[T Number](data ...T) T {
	var res T
	for _, v := range data {
		res = res + v
	}

	//var s1 string = "hello"
	//var s2 string = "world"
	//var s3 = s1 + s2

	return res
}

func SumV1[T Number](vals ...T) T {
	var a T
	return a
}

type MyMarshal struct {
}

func (m *MyMarshal) MarshalJSON() ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func UseSum() {
	fmt.Println(Sum[int](1, 2, 3))
	fmt.Println(Sum[float64](1.1, 2.2, 3.3))
	//fmt.Println(Sum[float32](1.1, 2.2, 3.3))

	fmt.Println(Sum[Integer](1, 2, 3)) // Cannot use Integer as the type Number Type does not implement constraint 'Number' because type is not included in type set ('int', 'uint', 'int64', 'float64')

	//Sum[float3](1, 2) // Cannot use float32 as the type Number Type does not implement constraint 'Number' because type is not included in type set ('int', 'uint')

	// Cannot use string as the type Number Type does not implement constraint 'Number' because type is not included in type set ('int', 'uint')
	//Sum("hello", "world")

	var j MyMarshal
	// 结构体实现了约束的接口，也可以传入泛型参数
	ReleaseResouce3[*MyMarshal](&j)

	// 用的时候一定是具体的类型
}

func ReleaseResouce[R io.Closer](r R) {
	r.Close()
}

func ReleaseResouce2[R time.Time](r R) {

}

func ReleaseResouce3[R json.Marshaler](r R) {
	r.MarshalJSON()
}
