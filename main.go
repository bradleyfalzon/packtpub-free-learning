package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, _ := goquery.NewDocument("https://www.packtpub.com/packt/offers/free-learning")
	fmt.Println(strings.Trim(doc.Find("h2").First().Text(), "\n\t"))
}
