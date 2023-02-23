package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

const Section int = 50
const File string = "/scripts/182.csv"
const TagId int = 186
const AccessToken = "66_9mgc156GwQkcsHOtIip-vJehsxnswmXOZQuR7t-tyNsB4zgmXG9a6Uh4mgyHQJuwnFyMRS-P3bvKbPKaeOoPCBUNE2vb6ER4yEeOJVBmjA3mQORq4vMqr4b0ZVkIFMhABAAXE"
const TagUrl = "https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging?access_token="

func main() {

	data := processCsv(File, Section)
	channel := make(chan struct{}, 5)
	var step int32 = 0

	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Printf("当前处理了: %d / 总计处理: %d \n", step, len(data))
			if step+1 == int32(len(data)) {
				break
			}
		}
	}()
	for one := range data {
		channel <- struct{}{}
		go func(one int) {
			// 处理微信的标签问题
			atomic.AddInt32(&step, 1)
			batchTag(data[one], TagId)
			<-channel
		}(one)
	}

}

// 用户批量打标签
func batchTag(people []string, tagId int) {
	url := TagUrl + AccessToken
	requestBody, err := json.Marshal(map[string]interface{}{
		"tagid":       tagId,
		"openid_list": people})
	if err != nil {
		log.Fatalln(err)
	}
	response, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))

}

func processCsv(file string, section int) [][]string {
	currentPath, _ := os.Getwd()
	filePath := currentPath + file
	fmt.Println(filePath)
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	// 临时容器
	var container []string
	// 返回容器
	var result [][]string
	records := csv.NewReader(f)
	for {
		record, err := records.Read()
		if err == io.EOF {
			if len(container) > 0 {
				result = append(result, append([]string(nil), container...))
			}
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		container = append(container, record[0])
		if len(container) == section {
			result = append(result, append([]string(nil), container...))
			container = nil
		}
	}
	return result
}
