package main

import (
    "fmt"
    "github.com/gocolly/colly/v2"
    "github.com/tebeka/selenium"
    "strconv"
    "strings"
    "time"
)

const (
    URL = "https://cafe.naver.com"

    seleniumPath     = "selenium-server-standalone-3.141.59.jar"
    chromeDriverPath = "chromedriver"
    port             = 4444
)

func main() {
    c := colly.NewCollector(
        colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36"),
        colly.AllowURLRevisit(),
    )
    DayAndUrl := make(map[string]string)

    c.OnHTML("a.article", func(e *colly.HTMLElement) {
        day := strings.TrimSpace(e.Text)
        url := e.Attr("href")
        if strings.Contains(day, "CF") {
            DayAndUrl[day] = url
        }
    })

    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL.String())
    })

    if err := c.Visit(URL + "/ArticleList.nhn?search.clubid=29406493&search.menuid=4&search.boardtype=L"); err != nil {
        fmt.Println(err)
    }
    for d := range DayAndUrl {
        if strings.Contains(d, getDate()) {
            getWod(URL + DayAndUrl[d])
            fmt.Printf("Day: %s, Url: %s\n", d, DayAndUrl[d])
        }
    }
}

func getDate() string {
    year, mon, day := time.Now().Date()
    var yearString = ""
    var monString = ""
    var dayString = ""
    yearString = strconv.Itoa(year)[2:]
    if int(mon) < 10 {
        monString = "0" + strconv.Itoa(int(mon))
    }
    if day < 10 {
        dayString = "0" + strconv.Itoa(day)
    }
    fullDate := yearString + monString + dayString
    return fullDate
}

func getWod(url string) {
    opts := []selenium.ServiceOption{
        selenium.ChromeDriver(chromeDriverPath),
    }
    service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
    if err != nil {
        panic(err)
    }
    defer service.Stop()
    caps := selenium.Capabilities{
        "browserName": "chrome",
    }
    wd, err := selenium.NewRemote(caps, "")
    if err != nil {
        panic(err)
    }
    wd.Get(url)
    time.Sleep(3 * time.Second)
    defer wd.Quit()

}
