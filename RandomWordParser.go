package main

import(

	_"github.com/andybalholm/cascadia"
	"fmt"
	"net/http"
	"strconv"
	"io/ioutil"
	"log"
	"strings"
	"time"
	"math/rand"
	"github.com/PuerkitoBio/goquery"
	_"go/doc"
)

type RowWordData struct{
	Word string
	Definition []string
	UsageExample []string
	WordRating int
}

const (
	WORD_COUNT        = "http://www.wordcount.org/dbquery.php?toFind="
	OXFORD_DICTIONARY = "https://en.oxforddictionaries.com/definition/"
)

func main() {

	// making random int from 0 : 88800
	var example RowWordData
	rand.Seed(time.Now().UnixNano())
	//index:= rand.Intn(88800)
	// initializing the "WordRating" field in example of RowWordData
	example.WordRating = 2
	// trying to get value of "Word" to type it into "Word" field in example of RowWordData
	word := getWord(example)
	example.Word = word
	// Getting 5 usage examples for
	usageExamples := getUsageExamples(example)
	example.UsageExample = usageExamples
	// gettting 5 or less definitions
	getWordDefinitions(example)

}
// got example like an argument and return "word" type - string
func getWord(example RowWordData)(string) {

	// got an example like an argument and parsing word from site with "WordRating" rating
	client := http.Client{}
	url := WORD_COUNT + strconv.Itoa(example.WordRating-1)
	url = url + "&method=SEARCH_BY_INDEX"

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	wholeLine := string(body)
	//got a response from site
	word := wholeLine[strings.Index(wholeLine, "word0")+6:strings.Index(wholeLine, "&freq0")]
	//got a word from full response
	fmt.Println(word)
	//return word
	return word
}
func getUsageExamples(example RowWordData)([]string){
	//need to change the word to example.Word
	urlToParse := OXFORD_DICTIONARY + "chaos"

	doc2, err := goquery.NewDocument(urlToParse)
	if err != nil {
		log.Fatal(err)
	}

	usageExamples := make([]string, 0)
	doc2.Find(".examples .exg .ex").Each(func(i int, s *goquery.Selection) {

		usageExample := s.Find("em").Text()
		if(i<5) {
			usageExamples = append(usageExamples , usageExample)
		}

	})

	i:= 0
	example.UsageExample = make([]string, 5)
	for i < len(usageExamples) {
		value := usageExamples[i]
		example.UsageExample[i] = value
		i++
	}
	return example.UsageExample
}
func getWordDefinitions(example RowWordData)([]string){

	urlToParse := "https://en.oxforddictionaries.com/definition/"
	urlToParse = urlToParse + "chaos"

	doc2, err := goquery.NewDocument(urlToParse)
	if err != nil {
		log.Fatal(err)
	}

	definitionExamples := make([]string, 0)
	doc2.Find(".subSenses li .ind ").Each(func(i int, s *goquery.Selection) {

		definitionExample := s.Find("span").Text()
		if i<5 {
			definitionExamples = append(definitionExamples , definitionExample)
			fmt.Println(definitionExamples[i])
		}

	})

	i:= 0
	example.Definition = make([]string, 5)
	for i<len(definitionExamples) {
		value := definitionExamples[i]
		example.UsageExample[i] = value
		i++
	}

	return example.Definition
}

