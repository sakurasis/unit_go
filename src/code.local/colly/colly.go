package main

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
)

func main() {
	var captchaId, img string
	var ip = "10.123.25.13"
	c := colly.NewCollector(
		colly.Debugger(&debug.LogDebugger{}),
		colly.AllowedDomains(ip),
	)
	handleCommonBiz(c)

	c.OnResponse(func(r *colly.Response) {

		assembleResponseHeaders(r)
		fmt.Println("response status:", r.StatusCode)

		body := r.Body
		captcha := &Captcha{}
		convertBodyToJson(body, captcha)
		// write to temp file.
		writeImgToFile(captcha, img)
		// get context from ocr
		// TODO
		captchaId = captcha.Data.CaptchaId
		fmt.Println("captchaId:", captchaId)

	})
	// get captcha
	err := c.Visit("https://" + ip + "/usmapi/v1/misc/captcha?r=0.8388254867807843&type=login_image")
	if err != nil {
		log.Fatal(err)
		return
	}
	//
	c2 := c.Clone()
	handleCommonBiz(c2)

}

func handleCommonBiz(c *colly.Collector) {
	cookieJar, _ := cookiejar.New(nil)
	c.SetCookieJar(cookieJar)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c.WithTransport(tr)
	c.OnRequest(func(r *colly.Request) {
		assembleRequestHeader(r)
	})
}

func assembleRequestHeader(r *colly.Request) {
	r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 Edg/108.0.1462.54")
	r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	r.Headers.Set("Accept-Encoding", "gzip, deflate, br")
	r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	r.Headers.Set("Cache-Control", "no-cache")
	r.Headers.Set("Connection", "keep-alive")
	r.Headers.Set("sec-ch-ua", "Not?A_Brand\";v=\"8\", \"Chromium\";v=\"108\", \"Microsoft Edge\";v=\"108")
	r.Headers.Set("sec-ch-ua-mobile", "?0")
	r.Headers.Set("sec-ch-ua-platform", "Windows")
	r.Headers.Set("Sec-Fetch-Dest", "document")
	r.Headers.Set("Sec-Fetch-Mode", "navigate")
	r.Headers.Set("Sec-Fetch-Site", "same-origin")
	r.Headers.Set("Sec-Fetch-User", "?1")
}

func assembleResponseHeaders(r *colly.Response) {
	r.Headers.Set("Server", "nginx")
	r.Headers.Set("Content-Type", "text/json")
	r.Headers.Set("Transfer-Encoding", "chunked")
	r.Headers.Set("Connection", "keep-alive")
	r.Headers.Set("ETag", "W/\"631adc36-26f\"")
	r.Headers.Set("X-Frame-Options", "sameorigin")
}

func writeImgToFile(captcha *Captcha, img string) {
	img = captcha.Data.Image
	data, err := base64.StdEncoding.DecodeString(img)

	if err != nil {
		log.Fatal(err)
	}
	f, _ := os.OpenFile("test.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	_, err = f.Write(data)
	if err != nil {
		return
	}
}

func convertBodyToJson(body []byte, captcha *Captcha) {
	err := json.Unmarshal(body, captcha)
	if err != nil {
		log.Fatal("analyse the json has error:", err)
		return
	}
}

type Data struct {
	CaptchaId string `json:"captcha_id"`
	Image     string `json:"image"`
}

type Captcha struct {
	Code string `json:"code"`
	Data Data
	Msg  string `json:"msg"`
}
