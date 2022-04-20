package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

func main() {
	c := colly.NewCollector(colly.Debugger(&debug.LogDebugger{}))
	c.OnHTML(".title-content", func(el *colly.HTMLElement) {
		fmt.Printf("%+v %+v\n", el.ChildText(".title-content-title"), el.Attr("href"))
	})
	c.OnRequest(func(req *colly.Request) {
		fmt.Println("Visiting", req.URL)
	})
	err := c.Visit("http://www.baidu.com")
	if err != nil {
		return
	}
}
