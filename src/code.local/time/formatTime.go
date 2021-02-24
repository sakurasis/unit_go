package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Unix()
	fmt.Printf("当前时间时间戳: %d\n", now)
	timestamp := 1609981366
	timeLayout := "2006-01-02 15:04:05"                           //转化所需模板
	datetime := time.Unix(int64(timestamp), 0).Format(timeLayout) //时间戳转化为日期
	fmt.Println("时间戳转换为日期格式为:", datetime)

	fmt.Println("=============计算花费工时=================")
	clockIn, err1 := time.Parse("2006-01-02 15:04:05", "2020-09-01 08:09:32")
	clockOut, err2 := time.Parse("2006-01-02 15:04:05", "2020-09-01 18:09:32")
	if err1 != nil && err2 != nil {
		fmt.Println("时间格式错误")
	}
	betweenTime := clockOut.Sub(clockIn)
	if betweenTime <= 0 {
		fmt.Println("时间差错误")
	}
	minutes := betweenTime.Minutes()
	fmt.Printf("You costs %.2f hours in company", (minutes-90)/60)
	fmt.Println("=========================================")
}
