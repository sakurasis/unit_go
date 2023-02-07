package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"math/rand"
	"time"
)

var username string = ""
var password string = ""
var url = "https://10.123.25.13/login"
var selector = "body"
var name = `document.querySelector("#rc-tabs-1-panel-pwd > div > form")`

func main() {

	content, err := GetHtmlContent(url, selector, name)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(content)
}

func GetHtmlContent(url, selector string, sel interface{}) (string, error) {
	agent := randomUserAgent()

	options := []chromedp.ExecAllocatorOption{
		chromedp.ExecPath("C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe"),
		chromedp.Flag("headless", true),
		chromedp.Flag("blink-settings", "imagesEnabled=true"),
		chromedp.UserAgent(agent),
		chromedp.IgnoreCertErrors,
		chromedp.Flag("disable-web-security", true),
	}

	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	// 创建上下文传递配置参数
	ctx, _ := chromedp.NewExecAllocator(context.Background(), options...)
	// 创建上下文对象
	chromectx, cancel := chromedp.NewContext(ctx, chromedp.WithLogf(log.Printf))

	err := chromedp.Run(chromectx, make([]chromedp.Action, 0, 1)...)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	//创建一个上下文，超时时间为40s
	timeoutCtx, cancel := context.WithTimeout(chromectx, 40*time.Second)
	defer cancel()
	var htmlContent string
	err = chromedp.Run(timeoutCtx)
	if err != nil {
		log.Printf("Run err : %v\n", err)
		return "", err
	}
	//log.Println(htmlContent)

	return htmlContent, nil
}

func click(sel interface{}, htmlContent string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitReady(selector, chromedp.ByQuery),
		chromedp.SetValue(`#username`, username, chromedp.ByID),
		chromedp.SetValue(`#password`, password, chromedp.ByID),
		chromedp.OuterHTML(sel, &htmlContent, chromedp.ByJSPath),
		chromedp.Click(`#submit`, chromedp.NodeVisible),
		chromedp.Sleep(150 * time.Second),
	}
}

func randomUserAgent() string {
	var chromeVersions = []string{
		"65.0.3325.146",
		"64.0.3282.0",
		"41.0.2228.0",
		"40.0.2214.93",
		"37.0.2062.124",
	}

	var osStrings = []string{
		"Macintosh; Intel Mac OS X 10_10",
		"Windows NT 10.0",
		"Windows NT 5.1",
		"Windows NT 6.1; WOW64",
		"Windows NT 6.1; Win64; x64",
		"X11; Linux x86_64",
	}
	version := chromeVersions[rand.Intn(len(chromeVersions))]
	os := osStrings[rand.Intn(len(osStrings))]
	return fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", os, version)
}
