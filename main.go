package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("https://www.packtpub.com/packt/offers/free-learning")
	if err != nil {
		log.Fatalln(err)
	}

	title := doc.Find("h2").First().Text()
	fmt.Println(strings.Trim(title, "\n\t"))
}
