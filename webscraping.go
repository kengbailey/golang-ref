/*

    Ken Bailey
    3/10/18

    Webscraping using github.com/gocolly/colly
    --> 9Clacks

*/

package main

import (
    "fmt"
    "github.com/gocolly/colly"
    "strings"
)

func main() {
    // Instantiate default collector
    c := colly.NewCollector(
        // visit only domains - needed if crawling
        colly.AllowedDomains("www4.9clacks2.net"),
    )
    // On every link tag post call callback
    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        if strings.Contains(e.Text, "[") && len(e.Text) > 20 {
            // print print post title
            fmt.Printf("Post found: %s\n", e.Text)
        }
    })
    // visit url
    c.Visit("http://www4.9clacks2.net")
}
