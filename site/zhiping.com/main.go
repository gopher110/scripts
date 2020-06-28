package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFile, err := os.Create(fmt.Sprintf("zhiping-jobs-%s.csv", time.Now().Format("12150405")))
	if err != nil {
		panic("创建文件失败")
	}
	defer csvFile.Close()

	//csvFile, err := os.OpenFile("zhiping-jobs-628073740.csv", os.O_RDWR|os.O_CREATE, 0666)
	//if err != nil {
	//	fmt.Printf("can not create file, err is %+v \n", err)
	//}
	//defer csvFile.Close()
	//csvFile.Seek(0, io.SeekEnd)
	//csvFile.WriteString("\xEF\xBB\xBF")

	writer := csv.NewWriter(csvFile)
	defer fmt.Println("****Crawl Test END***")
	for i := 1; ; i++ {
		if i == 1 {
			line := []string{"城市", "职位名称", "职位描述", "薪资", "经验", "学历", "标签", "更新时间", "职位链接", "公司名称", "公司地址", "成立日期", "简介"}
			writer.Write(line)
		}
		url := fmt.Sprintf("https://www.zhipin.com/c101020100-p100116/?page=%d&ka=page-%d", i, i)
		fmt.Printf("***TEST CRAWL / START PARSE URL:%s ****\n", url)
		if body, err := fetch(url); err == nil {
			list := getJobList(body)
			if len(list) == 0 {
				fmt.Println("!!!!!!!", body)
				return
			}
			for _, v := range list {
				if job := parseJob(v); job.Title != "" && job.Salary != "" {
					line := []string{job.City, job.Title, job.Description, job.Salary, job.Experience, job.Degree, strings.Join(job.Tags, ","), job.Time,
						job.Url, job.Company, job.Address, job.RegisterDate, job.Profile}
					fmt.Printf("=========%v \n", line)
					writer.Write(line)
				}
			}
			writer.Flush()
		} else {
			fmt.Printf("!!!!!!!%s\n page:%d url：%s\n", body, i, url)
			fmt.Println(err.Error())
			return
		}
	}
}
