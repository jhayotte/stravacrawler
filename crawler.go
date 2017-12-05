package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Crawl(url string, d chan float64) error {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return err
	}

	// grabs the distance
	s := doc.Find(".distance").First()
	if s.Text() == "" {
		return fmt.Errorf("Not Found for %s", url)
	}
	km := s.Find("strong").Text()
	if km == "" {
		return fmt.Errorf("Not Found")
	}

	km = strings.TrimSpace(km)
	//remove two last chars 'km' to get the value
	km = km[:len(km)-2]

	f, err := strconv.ParseFloat(km, 32)
	if err != nil {
		return err
	}
	// returns the value
	d <- f
	return nil
}
