package main

import "fmt"

type Spoker interface {
	Say()
}

// siri
type Siri struct{}

// xiaoai
type XiaoAi struct{}

func (s Siri) Say() {
	fmt.Println("你好，我是语音助手siri")
}

func (x XiaoAi) Say() {
	fmt.Println("我是小爱，你的小米语音助手")
}

func main() {
	// 声明Spoker类型的变量rx
	var rx Spoker

	var s Siri = Siri{}

	var x XiaoAi = XiaoAi{}

	rx = s

	rx.Say()

	rx = x

	rx.Say()

}
