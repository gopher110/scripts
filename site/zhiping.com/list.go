package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func getJobList(context string) (retVal []string) {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(context))
	if err != nil {
		fmt.Println("+++++getJobList++++", err)
		return
	}

	doc.Find("a[ka^='job_list_']").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			retVal = append(retVal, fmt.Sprintf("https://www.zhipin.com%s", href))
		}
	})

	return
}
