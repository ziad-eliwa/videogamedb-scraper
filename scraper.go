package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"

	//"log"
	"os"
	//"regexp"
	//"slices"
	//"strconv"
	"strings"
	"time"

	//"github.com/PuerkitoBio/goquery"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly"
)

func fileWriter(filename string) (*csv.Writer, *os.File) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	writer := csv.NewWriter(file)
	return writer, file
}

func render(url string) (string, error) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", false))
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	var renderedHTML string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(1250*time.Millisecond),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.Evaluate(`document.readyState === "complete"`, nil),
		chromedp.OuterHTML("html", &renderedHTML),
	)

	return renderedHTML, err
}

func main() {
	var url, renderedHTML string
	var err error
	// url = "https://www.mobygames.com/genre/"
	// renderedHTML, err = render(url)
	// if err != nil {}

	// genrec := colly.NewCollector()
	// categorywriter, category_file := fileWriter("csv/category.csv")
	// defer category_file.Close()
	// defer categorywriter.Flush()

	// genrec.OnHTML("section", func(e *colly.HTMLElement) {
	// 	title := e.ChildText("h3")
	// 	e.ForEach("ul li", func(i int, h *colly.HTMLElement) {
	// 		data := h.Text
	// 		var str string
	// 		for i := 0; i < len(data); i++ {
	// 			if data[i] == '(' {
	// 				str = data[0:i]
	// 			}
	// 		}
	// 		categorywriter.Write([]string{title, strings.TrimSpace(str)})
	// 	})
	// })

	// genrec.OnRequest(func(r *colly.Request) {
	// 	r.Ctx.Put("html", renderedHTML)
	// })
	// genrec.Visit(url)

	// //techspecs scraping
	// techspecsc := colly.NewCollector()
	// url = "https://www.mobygames.com/attributes/tech-specs/"
	// renderedHTML, err = render(url)
	// if err != nil {
	// 	fmt.Println("ChromeDP error", err)
	// }
	// var tech_specs_links []string
	// var tech_specs_titles []string
	// techspecsc.OnHTML("#main > ul:nth-child(4)", func(e *colly.HTMLElement) {
	// 	e.ForEach("li", func(i int, h *colly.HTMLElement) {
	// 		title := strings.TrimSpace(h.Text)
	// 		link := h.ChildAttr("a", "href")
	// 		fmt.Println(title, link)
	// 		tech_specs_links = append(tech_specs_links, link)
	// 		tech_specs_titles = append(tech_specs_titles, title)
	// 	})
	// })

	// techspecsc.OnError(func(r *colly.Response, err error) {
	// 	fmt.Println("Error", err)
	// })

	// techspecsc.OnResponse(func(r *colly.Response) {
	// 	fmt.Println("Visited link: ", url, "\n", "Got response, status:", r.StatusCode)
	// })

	// techspecsc.OnRequest(func(r *colly.Request) {
	// 	r.Ctx.Put("html", renderedHTML)
	// })

	// techspecsc.Visit(url)

	techspecs_writer, tech_file := fileWriter("csv/tech-specs.csv")
	defer tech_file.Close()
	defer techspecs_writer.Flush()
	// for i := 0; i < len(tech_specs_links); i++ {
	// 	renderedHTML, err := render("https://www.mobygames.com"+tech_specs_links[i])
	// 	if err != nil {
	// 		fmt.Println("ChromeDP error", err)
	// 	}

	// 	var attributes []string
	// 	techspecsc.OnHTML("#main > ul", func(e *colly.HTMLElement) {
	// 		e.ForEach("li",func(i int, h *colly.HTMLElement) {
	// 			if !slices.Contains(attributes,h.Text) {
	// 				attributes = append(attributes, strings.TrimSpace(h.Text))
	// 			}
	// 		})
	// 	})

	// 	techspecsc.OnError(func(r *colly.Response, err error) {
	// 		fmt.Println("Error", err)
	// 	})

	// 	techspecsc.OnResponse(func(r *colly.Response) {
	// 		fmt.Println("Visited link: ", tech_specs_links[i], "\n", "Got response, status:", r.StatusCode)
	// 	})

	// 	techspecsc.OnRequest(func(r *colly.Request) {
	// 		r.Ctx.Put("html", renderedHTML)
	// 	})

	// 	techspecsc.Visit("https://www.mobygames.com"+tech_specs_links[i])
	// 	for n := 0 ; n < len(attributes); n++ {
	// 		techspecs_writer.Write([]string{tech_specs_titles[i],attributes[n]})
	// 		techspecs_writer.Flush()
	// 	}
	// }

	// // maturity rating scraping
	// maturityc := colly.NewCollector()
	// url = "https://www.mobygames.com/attributes/ratings/"
	// renderedHTML, err = render(url)
	// if err != nil {
	// 	fmt.Println("ChromeDP error", err)
	// }
	// var maturity_rating_links []string
	// var maturity_titles []string
	// maturityc.OnHTML("#main > ul:nth-child(4)", func(e *colly.HTMLElement) {
	// 	e.ForEach("li", func(i int, h *colly.HTMLElement) {
	// 		title := strings.TrimSpace(h.Text)
	// 		link := h.ChildAttr("a", "href")
	// 		fmt.Println(title, link)
	// 		maturity_rating_links = append(maturity_rating_links, link)
	// 		maturity_titles = append(maturity_titles, title)
	// 	})
	// })

	// maturityc.OnError(func(r *colly.Response, err error) {
	// 	fmt.Println("Error", err)
	// })

	// maturityc.OnResponse(func(r *colly.Response) {
	// 	fmt.Println("Visited link: ", url, "\n", "Got response, status:", r.StatusCode)
	// })

	// maturityc.OnRequest(func(r *colly.Request) {
	// 	r.Ctx.Put("html", renderedHTML)
	// })

	// maturityc.Visit(url)

	maturitywriter, maturityfile := fileWriter("csv/maturity-rating.csv")
	defer maturityfile.Close()
	defer maturitywriter.Flush()

	// for i := 0; i < len(maturity_rating_links); i++ {
	// 	renderedHTML, err := render("https://www.mobygames.com"+maturity_rating_links[i])
	// 	if err != nil {
	// 		fmt.Println("ChromeDP error", err)
	// 	}

	// 	var attributes []string
	// 	maturityc.OnHTML("#main > ul", func(e *colly.HTMLElement) {
	// 		e.ForEach("li",func(i int, h *colly.HTMLElement) {
	// 			if !slices.Contains(attributes,h.Text) {
	// 				attributes = append(attributes, strings.TrimSpace(h.Text))
	// 			}
	// 		})
	// 	})

	// 	maturityc.OnError(func(r *colly.Response, err error) {
	// 		fmt.Println("Error", err)
	// 	})

	// 	maturityc.OnResponse(func(r *colly.Response) {
	// 		fmt.Println("Visited link: ", maturity_rating_links[i], "\n", "Got response, status:", r.StatusCode)
	// 	})

	// 	maturityc.OnRequest(func(r *colly.Request) {
	// 		r.Ctx.Put("html", renderedHTML)
	// 	})

	// 	maturityc.Visit("https://www.mobygames.com"+maturity_rating_links[i])
	// 	for n := 0 ; n < len(attributes); n++ {
	// 		maturitywriter.Write([]string{maturity_titles[i],attributes[n]})
	// 		maturitywriter.Flush()
	// 	}
	// }

	/*
		Platform Scraping
	*/

	// platformc := colly.NewCollector()

	// url = "https://www.mobygames.com/platform/"
	// renderedHTML, err = render(url)
	// var platform_links []string
	// var years_active []string
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
	// }

	// platformc.OnHTML("#results", func(e *colly.HTMLElement) {
	// 	e.ForEach("tr", func(i int, h *colly.HTMLElement) {
	// 		year_active := strings.TrimSpace(h.ChildText("td:nth-child(5)"))
	// 		year2, err := strconv.Atoi(year_active[len(year_active)-4:])
	// 		if (year2 >= 2020 && err == nil) || (year_active == "Undated–Undated") {
	// 			link := strings.TrimSpace(h.ChildAttr("td:nth-child(1) a", "href"))
	// 			platform_links = append(platform_links, "https://www.mobygames.com"+link)
	// 			years_active = append(years_active, year_active)
	// 		}
	// 	})
	// })

	// platformc.OnError(func(r *colly.Response, err error) {
	// 	fmt.Println("Error", err)
	// })

	// platformc.OnResponse(func(r *colly.Response) {
	// 	fmt.Println("Visited link: ", url, "\n", "Got response, status:", r.StatusCode)
	// })

	// platformc.OnRequest(func(r *colly.Request) {
	// 	r.Ctx.Put("html", renderedHTML)
	// })

	// platformc.Visit(url)

	// platformwriter, platformfile := fileWriter("csv/platform.csv")
	// defer platformfile.Close()
	// defer platformwriter.Flush()

	// for i := 0; i < len(platform_links); i++ {
	// 	renderedHTML, err = render(platform_links[i])

	// 	if err != nil {
	// 		fmt.Println("Error: ", err)
	// 		continue
	// 	}

	// 	var title string
	// 	var overview string
	// 	platformc.OnHTML("#main > h1", func(e *colly.HTMLElement) {
	// 		title = strings.TrimSpace(e.Text)
	// 		title = strings.TrimSpace(strings.ReplaceAll(title,"Games List"," "))
	// 	})

	// 	platformc.OnHTML("#main > p:nth-child(3)", func(e *colly.HTMLElement) {
	// 		overview = strings.TrimSpace(e.Text)
	// 		overview = strings.ReplaceAll(overview, "\n", " ")
	//  		overview = strings.ReplaceAll(overview, "\r", " ")
	// 	})

	// 	platformc.OnError(func(r *colly.Response, err error) {
	// 		fmt.Println("Error", err)
	// 	})

	// 	platformc.OnResponse(func(r *colly.Response) {
	// 		fmt.Println("Visited link: ", url, "\n", "Got response, status:", r.StatusCode)
	// 	})

	// 	platformc.OnRequest(func(r *colly.Request) {
	// 		r.Ctx.Put("html", renderedHTML)
	// 	})
	// 	platformc.Visit(platform_links[i])

	// 	platformwriter.Write([]string{title,overview,years_active[i]})
	// 	platformwriter.Flush()
	// }

	/*
		Games Scraping
	*/
	var games_links []string
	for i := 1; i <= 14; i++ {
		opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", false))
		allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), opts...)

		url := fmt.Sprintf("https://www.mobygames.com/game/from:2020/include_dlc:false/include_nsfw:false/release_status:all/until:2025/sort:moby_score/page:%v/", i)
		ctx, ctxCancel := chromedp.NewContext(allocCtx)

		var rendered string

		err := chromedp.Run(ctx,
			chromedp.Navigate(url),
			chromedp.WaitVisible("div.browser-grid", chromedp.ByQuery),
			chromedp.Sleep(1*time.Second),
			chromedp.OuterHTML("html", &rendered, chromedp.ByQuery),
		)
		if err != nil {
			log.Fatal(err)
		}

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(rendered))
		if err != nil {
			log.Fatal(err)
		}

		doc.Find("div.browser-grid.mb").Each(func(i int, sel *goquery.Selection) {
			sel.Find("table.table.table-hover.mb > tbody > tr > td:nth-child(1) > a:nth-child(2)").Each(func(j int, td *goquery.Selection) {
				link, c := td.Attr("href")
				if c {
					games_links = append(games_links, link)
				}
			})
		})
		ctxCancel()
		allocCancel()
	}

	// Gamewriter, Gamefile := fileWriter("csv/Games.csv")
	// defer Gamefile.Close()
	// defer Gamewriter.Flush()

	// GameDirectorwriter, GameDirectorfile := fileWriter("csv/gamedirector.csv")
	// defer GameDirectorfile.Close()
	// defer GameDirectorwriter.Flush()

	// Peoplewriter, Peoplefile := fileWriter("csv/People.csv")
	// defer Peoplefile.Close()
	// defer Peoplewriter.Flush()

	// MobyGameReviewwriter, MobyGameReviewfile := fileWriter("csv/MobyGameReview.csv")
	// defer MobyGameReviewfile.Close()
	// defer MobyGameReviewwriter.Flush()

	// GameReleasewriter, GameReleasefile := fileWriter("csv/GameRelease.csv")
	// defer GameReleasefile.Close()
	// defer GameReleasewriter.Flush()

	// GameCategorywriter, GameCategoryfile := fileWriter("csv/GameCategory.csv")
	// defer GameCategoryfile.Close()
	// defer GameCategorywriter.Flush()

	TechSpecsReleasewriter, TechSpecsReleasefile := fileWriter("csv/TechSpecsRelease.csv")
	defer TechSpecsReleasefile.Close()
	defer TechSpecsReleasewriter.Flush()

	MaturityRatingReleasewriter, MaturityRatingReleasefile := fileWriter("csv/MaturityRatingRelease.csv")
	defer MaturityRatingReleasefile.Close()
	defer MaturityRatingReleasewriter.Flush()

	// Publisherwriter, Publisherfile := fileWriter("csv/Publisher.csv")
	// defer Publisherfile.Close()
	// defer Publisherwriter.Flush()

	// Developerwriter, Developerfile := fileWriter("csv/Developer.csv")
	// defer Developerfile.Close()
	// defer Developerwriter.Flush()

	game := colly.NewCollector()
	// credits := colly.NewCollector()
	// releases := colly.NewCollector()
	specs := colly.NewCollector()
	// people := colly.NewCollector()

	var Name string
	//var URL, description string
	game.OnHTML("#main > div.flex.flex-space-between > div.mb > h1", func(e *colly.HTMLElement) {
		Name = strings.TrimSpace(e.Text)
		fmt.Println("Name: ", Name)
	})

	// game.OnHTML("#cover > img", func(e *colly.HTMLElement) {
	// 	URL = strings.TrimSpace(e.Attr("src"))
	// 	fmt.Println("Image URL: ", URL)
	// })

	// game.OnHTML("#description-text", func(h *colly.HTMLElement) {
	// 	h.ForEach("p", func(i int, h *colly.HTMLElement) {
	// 		pgph := h.Text
	// 		pgph = strings.TrimSpace(h.Text)
	// 		pgph = strings.ReplaceAll(pgph, "\n", "\t")
	// 		description += pgph
	// 	})
	// })

	// var company_urls []string
	// var publisher, developer string
	// game.OnHTML("#publisherLinks", func(e *colly.HTMLElement) {
	// 	e.ForEach("li", func(j int, h *colly.HTMLElement) {
	// 		link := strings.TrimSpace(h.Attr("href"))
	// 		publisher = strings.TrimSpace(h.Text)
	// 		fmt.Println(Name, "Publisher:", publisher)
	// 		if !slices.Contains(company_urls, link) {
	// 			company_urls = append(company_urls, link)
	// 		}
	// 		Publisherwriter.Write([]string{publisher, Name})
	// 		Publisherwriter.Flush()
	// 	})
	// })

	// game.OnHTML("#developerLinks", func(e *colly.HTMLElement) {
	// 	e.ForEach("li a", func(j int, h *colly.HTMLElement) {
	// 		link := strings.TrimSpace(h.Attr("href"))
	// 		developer = strings.TrimSpace(h.Text)
	// 		fmt.Println(Name, "Developer:", developer)
	// 		if !slices.Contains(company_urls, link) {
	// 			company_urls = append(company_urls, link)
	// 		}
	// 		Developerwriter.Write([]string{developer, Name})
	// 		Developerwriter.Flush()
	// 	})
	// })

	// game.OnHTML("#infoBlock > div.info-genres > dl", func(e *colly.HTMLElement) {
	// 	var currentKey string
	// 	e.DOM.Children().Each(func(_ int, s *goquery.Selection) {
	// 		if goquery.NodeName(s) == "dt" {
	// 			currentKey = s.Text()
	// 		} else if goquery.NodeName(s) == "dd" {
	// 			s.Find("a").Each(func(_ int, a *goquery.Selection) {
	// 				fmt.Println(currentKey, a.Text())
	// 				//GameCategorywriter.Write([]string{Name, currentKey, a.Text()})
	// 				//GameCategorywriter.Flush()
	// 			})
	// 		}
	// 	})
	// })

	// game.OnError(func(r *colly.Response, err error) {
	// 	fmt.Println("Error", err)
	// })

	// game.OnResponse(func(r *colly.Response) {
	// 	fmt.Println("Visited link: ", url, "\n", "Got response, status:", r.StatusCode)
	// })

	// var rhtml string
	// credits.OnHTML("div.overflow-x-scroll > table > tbody", func(e *colly.HTMLElement) {
	// 	e.ForEach("tr", func(i int, h *colly.HTMLElement) {
	// 		role := strings.TrimSpace(h.ChildText("td:nth-child(1)"))
	// 		if strings.Contains(role, "Direct") {
	// 			h.ForEach("td:nth-child(2) a", func(i int, k *colly.HTMLElement) {
	// 				if slices.Contains(MainDirectors, role) {
	// 					fmt.Println(Name, k.Text, role)
	// 					link := k.Attr("href")
	// 					rhtml, err = render(link)
	// 					ctx := colly.NewContext()
	// 					ctx.Put("name", Name)
	// 					ctx.Put("role", role)
	// 					ctx.Put("html", rhtml)
	// 					people.Request("GET", link, nil, ctx, nil)
	// 				}
	// 			})
	// 		}
	// 	})
	// })

	// people.OnHTML("#main > div.grid-split-2-1 > div:nth-child(1)", func(h *colly.HTMLElement) {
	// 	directorname := h.ChildText("h1")
	// 	directoraka := h.ChildText("p")
	// 	re := regexp.MustCompile(`\d+`)
	// 	directorid := re.FindString(directoraka)

	// 	name := h.Request.Ctx.Get("name")
	// 	role := h.Request.Ctx.Get("role")

	// 	fmt.Println(directorid, directorname)
	// 	Peoplewriter.Write([]string{directorid, directorname})
	// 	Peoplewriter.Flush()
	// 	GameDirectorwriter.Write([]string{name, directorid, role})
	// 	GameDirectorwriter.Flush()
	// })

	// people.OnError(func(r *colly.Response, err error) {
	// 	fmt.Println("Error", err)
	// })

	// people.OnResponse(func(r *colly.Response) {
	// 	fmt.Println("Visited link:", r.Request.URL, "\nGot response, status:", r.StatusCode)
	// })

	// credits.OnError(func(r *colly.Response, err error) {
	// 	fmt.Println("Error", err)
	// })

	// credits.OnResponse(func(r *colly.Response) {
	// 	fmt.Println("Visited link:", url+"credits/", "\nGot response, status:", r.StatusCode)
	// })

	// releases.OnHTML("#main", func(e *colly.HTMLElement) {
	// 	var PlatformName string
	// 	var release_date string
	// 	var comment string
	// 	var retailprice float64
	// 	e.DOM.Children().Each(func(_ int, s *goquery.Selection) {
	// 		if goquery.NodeName(s) == "h4" {
	// 			PlatformName = s.Text()
	// 		} else if goquery.NodeName(s) == "table" {
	// 			s.Find("tbody > tr.bg-dark > td").Each(func(i int, s *goquery.Selection) {
	// 				release_date = s.Text()
	// 				release_date = strings.ReplaceAll(release_date, "Release", " ")
	// 				release_date = strings.TrimSpace(release_date)
	// 			})
	// 			s.Find("tbody > tr").Each(func(_ int, a *goquery.Selection) {
	// 				first := strings.TrimSpace(a.Find("td").Eq(0).Text())
	// 				second := strings.TrimSpace(a.Find("td").Eq(1).Text())
	// 				switch first {
	// 				case "Comments:":
	// 					comment = second
	// 				case "MSRP:":
	// 					retailprice, err = strconv.ParseFloat(second[1:], 64)
	// 					if err != nil {
	// 					}
	// 				}
	// 			})
	// 			fmt.Println(Name, PlatformName, release_date, comment, retailprice)
	// 			GameReleasewriter.Write([]string{Name, PlatformName, release_date, comment, strconv.FormatFloat(retailprice, 'f', -1, 64)})
	// 			GameReleasewriter.Flush()
	// 		}
	// 	})
	// })

	// releases.OnError(func(r *colly.Response, err error) {
	// 	fmt.Println("Error", err)
	// })

	// releases.OnResponse(func(r *colly.Response) {
	// 	fmt.Println("Visited link: ", url+"releases/", "\n", "Got response, status:", r.StatusCode)
	// })

	specs.OnHTML("#main > div.grid-split > div:nth-child(1) > table > tbody", func(e *colly.HTMLElement) {
		var platform string
		e.ForEach("tr", func(i int, h *colly.HTMLElement) {
			if h.ChildText("td:nth-child(2)") == "" {
				platform = strings.TrimSpace(h.ChildText("td:nth-child(1) > h4"))
				platform = strings.TrimSpace(strings.ReplaceAll(platform, "+", " "))
			} else {
				stype := strings.TrimSpace(h.ChildText("td:nth-child(1)"))
				h.ForEach("td:nth-child(2) > ul > li", func(i int, k *colly.HTMLElement) {
					sattr := strings.TrimSpace(k.ChildText("a"))
					fmt.Println(Name, platform, stype, sattr)
					techspecs_writer.Write([]string{stype,sattr})
					techspecs_writer.Flush()
					TechSpecsReleasewriter.Write([]string{Name, platform, stype, sattr})
					TechSpecsReleasewriter.Flush()
				})
			}
		})
	})

	specs.OnHTML("#main > div.grid-split > div:nth-child(2) > table > tbody", func(e *colly.HTMLElement) {
		var platform string
		e.ForEach("tr", func(i int, h *colly.HTMLElement) {
			if h.ChildText("td:nth-child(2)") == "" {
				platform = strings.TrimSpace(h.ChildText("th > h4"))
				platform = strings.TrimSpace(strings.ReplaceAll(platform, "+", " "))
			} else {
				stype := strings.TrimSpace(h.ChildText("td:nth-child(1)"))
				h.ForEach("td:nth-child(2) > ul > li", func(i int, k *colly.HTMLElement) {
					sattr := strings.TrimSpace(k.ChildText("a"))
					fmt.Println(Name, platform, stype, sattr)
					maturitywriter.Write([]string{stype,sattr})
					maturitywriter.Flush()
					MaturityRatingReleasewriter.Write([]string{Name, platform, stype, sattr})
					MaturityRatingReleasewriter.Flush()
				})
			}
		})
	})

	specs.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error", err)
	})

	specs.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited link: ", url+"specs/", "\n", "Got response, status:", r.StatusCode)
	})

	for i := 0; i < len(games_links); i++ {
		url = games_links[i]
		renderedHTML, err = render(url)
		if err != nil {
		}

		ctx := colly.NewContext()
		ctx.Put("url", url)
		ctx.Put("Name", Name)
		ctx.Put("html", renderedHTML)
		game.Request("GET", url, nil, ctx, nil) 

		// Gamewriter.Write([]string{Name, URL, description})
		// Gamewriter.Flush()
		// description = ""

		// renderedHTML, err = render(url + "credits/")
		// if err != nil {
		// }

		// ctxCredits := colly.NewContext()
		// ctxCredits.Put("url", url+"credits/")
		// ctxCredits.Put("Name", Name)
		// ctxCredits.Put("html", renderedHTML)
		// credits.Request("GET", url+"credits/", nil, ctxCredits, nil) 

		// renderedHTML, err = render(url + "releases/")
		// if err != nil {
		// }

		// ctxReleases := colly.NewContext()
		// ctxReleases.Put("url", url+"releases/")
		// ctxReleases.Put("Name", Name)
		// ctxReleases.Put("html", renderedHTML)
		// releases.Request("GET", url+"releases/", nil, ctxReleases, nil)

		renderedHTML, err = render(url + "specs/")
		if err != nil {
		}

		ctxSpecs := colly.NewContext()
		ctxSpecs.Put("url", url+"specs/")
		ctxSpecs.Put("Name", Name)
		ctxSpecs.Put("html", renderedHTML)
		specs.Request("GET", url+"specs/", nil, ctxSpecs, nil) 

		// opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", false))
		// allocCtx, Alloccancel := chromedp.NewExecAllocator(context.Background(), opts...)
		// ctxChrome, Ctxcancel := chromedp.NewContext(allocCtx)

		// var rendered string

		// err = chromedp.Run(ctxChrome,
		// 	chromedp.Navigate(url+"reviews/"),
		// 	chromedp.Sleep(3*time.Second),
		// 	chromedp.WaitVisible("div.overflow-x-scroll.mb", chromedp.ByQuery),
		// 	chromedp.OuterHTML("html", &rendered, chromedp.ByQuery),
		// )
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// doc, err := goquery.NewDocumentFromReader(strings.NewReader(rendered))
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// doc.Find("#main > div.overflow-x-scroll.mb > table > tbody").Each(func(i int, sel *goquery.Selection) {
		// 	sel.Find("tr").Each(func(j int, h *goquery.Selection) {
		// 		cell := h.Find("td:nth-child(1)")
		// 		cell.Find("small").Remove()
		// 		platform := strings.TrimSpace(cell.Text())
		// 		if platform != "Overall" {
		// 			crating := strings.TrimSpace(h.Find("td:nth-child(2)").Text())
		// 			crating = crating[0:2]
		// 			if crating != "n/" {
		// 				intcrating, err := strconv.ParseFloat(crating, 64)
		// 				if err != nil {
		// 				}
		// 				intcrating = intcrating * 10.0 / 100.0
		// 				crating = strconv.FormatFloat(intcrating, 'f', -1, 64)
		// 			} else {
		// 				crating = ""
		// 			}

		// 			ccount := h.Find("td:nth-child(2) > a").Text()
		// 			ccount = strings.TrimSpace(strings.ReplaceAll(ccount, "ratings", " "))
		// 			ccount = strings.TrimSpace(strings.ReplaceAll(ccount, "rating", " "))

		// 			prating := h.Find("td:nth-child(3) > span").AttrOr("data-tooltip", "")
		// 			if prating != "" {
		// 				prating = strings.TrimSpace(strings.ReplaceAll(prating, "stars", " "))
		// 				stars, err := strconv.ParseFloat(prating, 64)
		// 				if err == nil {
		// 					stars *= 2
		// 					prating = strconv.FormatFloat(stars, 'f', -1, 64)
		// 				}
		// 			}
		// 			pcount := h.Find("td:nth-child(3) > a").Text()
		// 			pcount = strings.TrimSpace(strings.ReplaceAll(pcount, "ratings", " "))
		// 			pcount = strings.TrimSpace(strings.ReplaceAll(pcount, "rating", " "))

		// 			mobyrating := h.Find("td:nth-child(4) > div").Text()
		// 			fmt.Println(platform, crating, ccount, prating, pcount, mobyrating)
		// 			MobyGameReviewwriter.Write([]string{Name, platform, crating, prating, pcount, ccount, mobyrating})
		// 			MobyGameReviewwriter.Flush()
		// 		}
		// 	})
		// })
		// Alloccancel()
		// Ctxcancel()
	}

	// /*
	// 	Companies Scraping
	// */
	// companyc := colly.NewCollector()
	// companywriter, companyfile := fileWriter("csv/companies.csv")
	// defer companyfile.Close()
	// defer companywriter.Flush()
	// var title string
	// companyc.OnHTML("#main > div.grid-split-2-1 > div:nth-child(1) > h1", func(e *colly.HTMLElement) {
	// 	title = strings.TrimSpace(e.Text)
	// })
	// var overview string
	// companyc.OnHTML("#description-text", func(e *colly.HTMLElement) {
	// 	overview = strings.TrimSpace(e.Text)
	// 	overview = strings.ReplaceAll(overview, "\n", " ")
	// 	overview = strings.ReplaceAll(overview, "\r", " ")
	// })
	// var link string
	// companyc.OnHTML("#companyLogo > figure > a > img", func(e *colly.HTMLElement) {
	// 	link = strings.TrimSpace(e.Attr("src"))
	// })
	// var yearsactive1, yearsactive2 string
	// companyc.OnHTML("#companyGames > h2", func(e *colly.HTMLElement) {
	// 	text := strings.TrimSpace(e.Text)
	// 	yearsactive1 = text[len(text)-4:]
	// 	yearsactive2 = text[len(text)-12 : len(text)-8]
	// 	fmt.Println("Years: ", yearsactive2, " ", yearsactive1)
	// })

	// companyc.OnError(func(r *colly.Response, err error) {
	// 	fmt.Println("Error", err)
	// })

	// companyc.OnResponse(func(r *colly.Response) {
	// 	fmt.Println("Visited link: ", "\n", "Got response, status:", r.StatusCode)
	// })

	// companyc.OnRequest(func(r *colly.Request) {
	// 	r.Ctx.Put("html", renderedHTML)
	// })

	// for i := 0; i < len(company_urls); i++ {
	// 	var company_details []string
	// 	renderedHTML, err = render(company_urls[i])
	// 	if err != nil {
	// 		fmt.Println("ChromeDP error", err)
	// 		continue
	// 	}

	// 	ctxCompany := colly.NewContext()
	// 	ctxCompany.Put("url", company_urls[i])
	// 	ctxCompany.Put("html", renderedHTML)
	// 	companyc.Request("GET", company_urls[i], nil, ctxCompany, nil)

	// 	found := false

	// 	var country string
	// 	words := strings.Fields(overview)

	// 	for j := 0; j < len(words) && !found; j++ {
	// 		word := strings.Trim(words[j], ".,!?;:()")
	// 		if slices.Contains(nationalities, word) {
	// 			country = nationalityMap[word]
	// 			found = true
	// 		} else if slices.Contains(countryNames, word) {
	// 			country = word
	// 			found = true
	// 		} else if slices.Contains(americanCities, word) {
	// 			country = "United States"
	// 			found = true
	// 		} else if slices.Contains(capitals, word) {
	// 			country = capitalCountryMap[word]
	// 			found = true
	// 		}
	// 	}

	// 	if !found {
	// 		fmt.Println("Country Not Found")
	// 	} else {
	// 		fmt.Println("Country Found : ", country)
	// 	}
	// 	company_details = append(company_details, title)
	// 	company_details = append(company_details, overview)
	// 	company_details = append(company_details, link)
	// 	company_details = append(company_details, yearsactive2)
	// 	company_details = append(company_details, yearsactive1)
	// 	company_details = append(company_details, country)
	// 	companywriter.Write(company_details)
	// 	companywriter.Flush()
	// }
}
