package sinonimo

import (
	"errors"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	errorWhileSearchingInWordReference = "Error while search in wordreference"
)

type SearchSynonyms interface {
	search(word string) ([]string, error)
}

type WordReference struct {
}

func (w WordReference) search(word string) ([]string, error) {
	words := make([]string, 0)

	res, err := http.Get("https://www.wordreference.com/sinonimos/" + word)
	if err != nil {
		log.Fatal(err)
		errors.New(errorWhileSearchingInWordReference)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		errors.New(errorWhileSearchingInWordReference)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div#article > div > ul > li").Each(func(i int, s *goquery.Selection) {
		regex := regexp.MustCompile(`,\s+`)
		wordsResult := regex.ReplaceAllString(s.Text(), ", ")
		words = append(words, strings.Split(wordsResult, ", ")...)
	})

	return words, nil
}
