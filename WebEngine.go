package main

import (
	"math/rand"
	"strconv"
	"strings"
	"net/http"
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"time"
)


type RowWordData struct {
	Word           string
	Definitions    map[string][]string
	UsageExamples  map[string][]string
	WordRang       int

}


const (
	WORD_COUNT_URL        = "http://www.wordcount.org/dbquery.php?"
	OXFORD_DICTIONARY_URL = "https://en.oxforddictionaries.com/definition/"

)

func GetRandomWord() (*RowWordData, error) {
	var (
		err error
		wordData = new(RowWordData)

	)
	rand.Seed(time.Now().UnixNano())
	wordData.Word, err = getWordData(rand.Intn(86799))
	if err != nil {
		return nil, err
	}

	wordData.Definitions, wordData.UsageExamples,err = getExtendedWordData(wordData.Word)
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


func getExtendedWordData(word string) (Definitions map[string][]string, UsageExamples map[string][]string, err error) {
	urlToParse := OXFORD_DICTIONARY_URL + word
	wordPage,err := goquery.NewDocument(urlToParse)
	if err!=nil {
		return nil,nil,err
	}

	definitions := make(map[string][]string)
	usageExamples := make( map[string][]string )

	wordPage.Find(".gramb ").Each(func(i int, s *goquery.Selection) {
		languagePart := s.Find("h3 .pos").Text()
		definition := s.Find(".semb .trg .ind").Text()
		usageExample :=	s.Find(".semb li .trg .examples .exg .ex em").Text()
		usageExamples[languagePart] = append(usageExamples[languagePart] , usageExample)
		definitions[languagePart] = append(definitions[languagePart] , definition)
	})
	return definitions, usageExamples,nil
}
