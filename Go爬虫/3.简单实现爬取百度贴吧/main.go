/**
* @file tieba_spider.go
* @brief  百度贴吧小爬虫 基于go语言(并发的爬取)
* @author

Aceld(LiuDanbing)

email: danbing.at@gmail.com
Blog: http://www.gitbook.com/@aceld

* @version 1.0
* @date 2017-11-28
*/
package main

import "fmt"
import "net/http"
import "io/ioutil"
import "strconv"
import "os"

//第一页: https://tieba.baidu.com/f?kw=lol&ie=utf-8&pn=0
//第2页: https://tieba.baidu.com/f?kw=lol&ie=utf-8&pn=50
//第3页: https://tieba.baidu.com/f?kw=lol&ie=utf-8&pn=100
//第4页: https://tieba.baidu.com/f?kw=lol&ie=utf-8&pn=150

/* -------------------------------------------*/
/**
* @brief  将一个url资源down下来
*
* @param string
*
* @returns 资源的数据
 */
/* -------------------------------------------*/
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

//爬取一页
func spider_one_page(page int, ch chan int) {

	var pn int

	//page =1, pn = 0
	//page =2, pn = 50
	//page =3, pn = 100
	//page =4, pn = 150
	pn = (page - 1) * 50

	url := "https://tieba.baidu.com/f?kw=lol&ie=utf-8&pn=" + strconv.Itoa(pn)
	fmt.Println(url)

	content := httpGet(url)

	//此时的content就是当前页面的数据 将这个数据保存在本地
	//1.html, 2.html, 3.html
	filename := strconv.Itoa(page) + ".html"
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}

	//将网页数据存储到文件中
	f.WriteString(content)
	f.Close()

	ch <- page
}

//爬虫函数 爬取一个区间页码
func spider_tieba(begin_page int, end_page int) {

	fmt.Println("准备爬取 ", begin_page, " -- ", end_page, "页")

	ch := make(chan int)

	for page := begin_page; page <= end_page; page++ {
		fmt.Println("正在爬取", page, "页")

		//开启一个goroutine(go程)
		go spider_one_page(page, ch)
	}

	for i := begin_page; i <= end_page+i; i++ {
		page := <-ch
		fmt.Println(page, "已经爬取完毕")
	}
}

func main() {
	var begin_page string
	var end_page string

	fmt.Println("请输入要爬取的起始页码")
	fmt.Scanf("%s\n", &begin_page)
	fmt.Println("请输入要爬取的终止页码")
	fmt.Scanf("%s\n", &end_page)

	//将页码由字符串类型转换成int
	b, _ := strconv.Atoi(begin_page)
	e, _ := strconv.Atoi(end_page)

	//开始爬取主业务
	spider_tieba(b, e)

}
