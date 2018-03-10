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
	Definitions   []string
	UsageExamples []string
	WordRang      int
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

	wordData.Definitions, wordData.UsageExamples, err = getExtendedWordData(wordData.Word)
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


func getExtendedWordData(word string) (Definitions []string, UsageExamples []string, err error) {
	urlToParse := OXFORD_DICTIONARY_URL + word
	wordPage,err := goquery.NewDocument(urlToParse)

	if err!=nil {

		return nil,nil,err

	}

	usageExamples := make([]string, 0)
	wordPage.Find(".examples .exg .ex").Each(func(i int, s *goquery.Selection) {
		usageExample := s.Find("em").Text()
		if(i<5) {
			usageExamples = append(usageExamples , usageExample)
		}
	})

	definitions := make([]string, 0)
	wordPage.Find(".semb li .trg ol li" ).Each(func(i int, s *goquery.Selection) {
		if i<5 {
			definitionExample := s.Find("span").Text()
			definitions = append(definitions, definitionExample)
		}
	})
	return definitions, usageExamples,nil
}


func getRandom() rand.Rand {
	var random rand.Rand

	for i:=0; i < 100; i++ {
		random.Intn(86800)
	}

	return random
}