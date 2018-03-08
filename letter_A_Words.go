package main

import(

	_"github.com/andybalholm/cascadia"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	_ "strings"
)

func main() {



	doc, err := goquery.NewDocument("https://dictionary.cambridge.org/dictionary/english/")
	if err != nil {
		log.Fatal(err)
	}



	doc.Find(".entry_title span .base").Each(func(i int, s *goquery.Selection) {
		word := s.Find("b").Text()

		fmt.Printf("Word %d: %s \n", i, word)

	})
	

}
