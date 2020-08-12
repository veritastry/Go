package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 简单爬虫代码

func httpGet(url string) string {

	//发送一个url
	resp, err := http.Get(url)
	if err != nil {
		//返回失败
		fmt.Println("http get error")
		return ""
	}

	//resp 里面存放的就是从url返回的数据
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//返回失败
		fmt.Println("readall error")
		return ""
	}

	content := string(data)

	resp.Body.Close()
	return content
}

func main() {

	content := httpGet("https://www.qq.com/")

	fmt.Println("从百度获取的数据 ", content)

}
