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
	for _, wb := range websites {
		go goWebsite(wb)
	}
	cost := time.Since(started)
	fmt.Printf("It took %v nanoSeconds times to visit these websites.\n", cost.Nanoseconds())
	time.Sleep(3 * time.Second)
}

// the common function
func goWebsite(wb string) {
	_, err := http.Get(wb)
	if err != nil {
		fmt.Printf("ERR: %s is down\n", wb)
		return
	}
	fmt.Printf("SUCCESS: %s is running\n", wb)
}
