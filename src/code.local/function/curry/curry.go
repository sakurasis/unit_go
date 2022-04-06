package curry

import "fmt"

// 定义可柯里化函数形式
type Function func(...interface{}) interface{}

// 通用柯里化函数
func (f Function) Curry(i interface{}) func(...interface{}) interface{} {
	return func(values ...interface{}) interface{} {
		values = append([]interface{}{i}, values...)
		fmt.Println(values)
		return f(values...)
	}
}
