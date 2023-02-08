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
var name = `document.querySelector("#rc-tabs-1-panel-pwd > div > form")`

func main() {

	content, err := GetHtmlContent()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(content)
}

func GetHtmlContent() (bool, error) {
	agent := randomUserAgent()

	options := []chromedp.ExecAllocatorOption{
		chromedp.ExecPath("C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe"),
		chromedp.Flag("headless", false),
		chromedp.Flag("blink-settings", "imagesEnabled=true"),
		chromedp.UserAgent(agent),
		chromedp.IgnoreCertErrors,
		chromedp.NoDefaultBrowserCheck,
		chromedp.NoFirstRun,
		chromedp.WindowSize(1080, 720),
		chromedp.Flag("disable-web-security", true),
	}

	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	// 创建上下文传递配置参数
	ctx, _ := chromedp.NewExecAllocator(context.Background(), options...)
	// 创建上下文对象实例
	chromectx, cancel := chromedp.NewContext(ctx, chromedp.WithLogf(log.Printf))

	//创建一个上下文，超时时间为40s
	timeoutCtx, cancel := context.WithTimeout(chromectx, 40*time.Second)
	defer cancel()
	// 验证码
	var captcha string

	err := chromedp.Run(timeoutCtx, click(captcha))
	if err != nil {
		log.Printf("Run err : %v\n", err)
		return false, err
	}
	//log.Println(htmlContent)

	return true, nil
}

func getContextByCaptcha() chromedp.Tasks {
	return chromedp.Tasks{}
}

// #rc-tabs-1-panel-pwd > div > form > div.b_captcha > div.b_c > img
func click(captcha string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitReady(`document.querySelector("#root > div > header > div.l_right > div.r_login_box")`, chromedp.ByJSPath),
		chromedp.SetValue(`#username`, username, chromedp.ByID),
		chromedp.SetValue(`#password`, password, chromedp.ByID),
		chromedp.SetValue("#captcha", captcha, chromedp.ByID),
		chromedp.Click(`#rc-tabs-1-panel-pwd > div > form > button`, chromedp.NodeVisible),
		chromedp.Sleep(10 * time.Second),
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
