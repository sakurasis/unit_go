package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	started := time.Now()

	websites := []string{
		"https://www.google.com",
		"https://www.baidu.com",
		"https://www.sina.com",
		"https://www.qq.com",
	}
	ch := make(chan string, 10)
	for _, wb := range websites {
		go goWebsite(wb, ch)
	}
	for i := 0; i < len(websites); i++ {
		fmt.Print(<-ch)
	}

	cost := time.Since(started)
	fmt.Printf("It took %v seconds times to visit these websites.\n", cost.Seconds())
	time.Sleep(3 * time.Second)
}

// the common function
func goWebsite(wb string, ch chan string) {
	_, err := http.Get(wb)
	if err != nil {
		ch <- fmt.Sprintf("ERR: %s is down\n", wb)
		return
	}
	ch <- fmt.Sprintf("SUCCESS: %s is running\n", wb)
}
