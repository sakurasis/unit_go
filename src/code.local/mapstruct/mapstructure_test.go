package mapstruct

import (
	"fmt"
	"log"
	"testing"

	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name   string
	Age    int
	Gender string
}

func Test_mapstruct(t *testing.T) {
	// 示例数据
	data := map[string]interface{}{
		"name":   "John",
		"age":    30,
		"gender": "male",
	}

	// 创建一个空的Person实例
	var person Person

	// 使用mapstructure库将map数据映射到Person实例
	err := mapstructure.Decode(data, &person)
	if err != nil {
		log.Fatal(err)
	}

	// 打印解析后的Person实例
	fmt.Printf("Name: %s\nAge: %d\nGender: %s\n", person.Name, person.Age, person.Gender)
}
