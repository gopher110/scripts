package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
)

type jobDetail struct {
	//城市
	City string
	//职位链接
	Url string
	//职位名称
	Title string
	//职位描述
	Description string
	//薪资
	Salary string
	//经验
	Experience string
	//学历
	Degree string
	//公司名称
	Company string
	//公司地址
	Address string
	//成立日期
	RegisterDate string
	//简介
	Profile string
	//标签
	Tags []string
	//创建时间
	Time string
}

//Requirements:
// Company:  Stage: Scale: Industry: RegisterDate: Remark:}
func parseJob(url string) (job jobDetail) {
	if body, err := fetch(url); err == nil {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
		if err != nil {
			fmt.Printf("goquery NewDocumentFromReader err %v", err)
		}
		//职位链接、名称、薪资
		s := doc.Find(".job-banner ")
		updateTime := s.Find(".time").Text()
		getDate := func(str string) (retVal string) {
			reg, _ := regexp.Compile("\\d{4}-\\d{2}-\\d{2}")
			if reg.MatchString(str) == true {
				retVal = reg.FindStringSubmatch(str)[0]
			}
			return
		}
		job.Time = getDate(updateTime)

		doc.Find(".job-tags span").Each(func(i int, s *goquery.Selection) {
			job.Tags = append(job.Tags, s.Text())
		})
		job.Url, job.Title, job.Salary = url, doc.Find("title").First().Text(), s.Find("span[class=salary]").First().Text()

		//所在城市、工作经验、最低学历
		reg := regexp.MustCompile(`<p>([^<]+)<em class="\w+"></em>([^<]+)<em class="\w+"></em>([^<]+)</p>`)
		if reg.MatchString(body) {
			subMatch := reg.FindStringSubmatch(body)
			job.City, job.Experience, job.Degree = subMatch[1], subMatch[2], subMatch[3]
		}

		//公司名称、公司地址、成立日期
		s = doc.Find(".business-info")
		job.Company, job.Address = s.Find("h4").First().Text(), doc.Find(".job-location .location-address").First().Text()
		job.RegisterDate = getDate(s.Find("td span:contains(成立时间)").First().Parent().Text())

		//职位描述&公司介绍
		job.Description = strings.TrimSpace(doc.Find("h3:contains(职位描述)+div").Text())
		job.Profile = strings.TrimSpace(doc.Find("h3:contains(公司介绍)").Parent().Find("p").First().Text())
	} else {
		fmt.Println("########parseJob err#########", url, err)
	}

	return
}
