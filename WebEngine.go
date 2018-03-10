package main

import (
	"math/rand"
	"strconv"
	"strings"
	"net/http"
	"bytes"
	"github.com/PuerkitoBio/goquery"
)


type RowWordData struct {
	Word          string
	Definitions   [][]string
	UsageExamples [][]string
	WordRang      int
	partOfLanguage []string
}


const (
	WORD_COUNT_URL        = "http://www.wordcount.org/dbquery.php?"
	OXFORD_DICTIONARY_URL = "https://en.oxforddictionaries.com/definition/"

)

func GetRandomWord() (*RowWordData, error) {
	var (
		err error
		wordData = new(RowWordData)
		random = getRandom()
	)

	wordData.Word, err = getWordData(random.Intn(86800))
	if err != nil {
		return nil, err
	}

	wordData.Definitions, wordData.UsageExamples, wordData.partOfLanguage ,err = getExtendedWordData(wordData.Word)
	if err != nil {
		return nil, err
	}

	return wordData, nil
}


func getWordData(index int) (word string, err error)  {
	url := WORD_COUNT_URL + "toFind=" + strconv.Itoa(index) + "&method=SEARCH_BY_INDEX"

	resp, err := new(http.Client).Get(url)
	if err != nil {
		return "",  err
	}
	defer resp.Body.Close()

	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(resp.Body)
	if err != nil {
		return "", err
	}

	parsedResponse := buffer.String()
	word = parsedResponse[strings.Index(parsedResponse, "word0") + 6 : strings.Index(parsedResponse, "&freq0")]


	return word,nil
}


func getExtendedWordData(word string) (Definitions [][]string, UsageExamples [][]string,partOftheLanguage []string, err error) {
	urlToParse := OXFORD_DICTIONARY_URL + word
	wordPage,err := goquery.NewDocument(urlToParse)

	if err!=nil {

		return nil,nil,nil,err

	}

	partsOfTheLanguage := make([]string, 0)

	wordPage.Find("section .gramb h3 .pos ").Each(func(i int, s *goquery.Selection) {

			value := s.Find("span").Text()
			partsOfTheLanguage = append(partsOfTheLanguage,value)


	})

	usageExamples := make([][]string,len(partsOfTheLanguage) , 0)
	definitions := make([][]string,len(partsOfTheLanguage) , 0)

	wordPage.Find("section .gramb h3 .pos ").Each(func(i int, s *goquery.Selection) {

		wordPage.Find(".semb li .trg .ind" ).Each(func(j int, s *goquery.Selection) {

			definitionExample := s.Find("span").Text()
			definitions[i] = append(definitions[i],definitionExample)



		})

		wordPage.Find(".examples .exg .ex" ).Each(func(k int, s *goquery.Selection) {
			if k < 5 {
				usageExample := s.Find("em").Text()
				usageExamples[i] = append(usageExamples[i], usageExample)
			}
		})



	})



	return definitions, usageExamples,partsOfTheLanguage,err
}


func getRandom() rand.Rand {
	var random rand.Rand

	for i:=0; i < 100; i++ {
		random.Intn(86800)
	}

	return random
}