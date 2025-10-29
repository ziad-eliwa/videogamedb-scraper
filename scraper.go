package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly"
)


func fileWriter(filename string) (*csv.Writer, *os.File){
	file, error := os.Create(filename)
	if error != nil {
		fmt.Println(error)
	}
	writer := csv.NewWriter(file)
	
	return writer, file
}

func render(url string) string {
	opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", false))
    allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
    defer cancel()
    ctx, cancel := chromedp.NewContext(allocCtx)
    defer cancel()
	
	var renderedHTML string
	err := chromedp.Run(ctx,
        chromedp.Navigate(url),
        chromedp.Sleep(5*time.Second),
        chromedp.OuterHTML("html", &renderedHTML),
    )
    if err != nil {
        log.Fatal("ChromeDP error:", err)
    }

	return renderedHTML
}


func main() {
	c := colly.NewCollector()
	var renderedHTML, url string


	url = "https://www.mobygames.com/genre/"
	renderedHTML = render(url)
	writer, file := fileWriter("csv/category.csv")
	defer file.Close()
	defer writer.Flush()
	
	c.OnHTML("section", func(e *colly.HTMLElement) {
		title := e.ChildText("h3")
		e.ForEach("ul li", func(i int, h *colly.HTMLElement) {
			data := h.Text
			var str string
			for i := 0; i < len(data); i++{
				if data[i] == '(' {
					str = data[0:i]
				}
			} 
			writer.Write([]string{title,strings.TrimSpace(str)})
		})
	})

	c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("html",renderedHTML)
	})
    c.Visit(url)
}
