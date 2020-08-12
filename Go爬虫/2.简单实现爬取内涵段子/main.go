package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Spider struct {
	page int // 表示当前爬虫正在爬取哪一页
}

func (this *Spider) HttpGet(url string) (content string,statusCode int) {
	resp,err := http.Get(url)
	if err != nil{
		fmt.Println(err)
		statusCode = -100
		return
	}
	defer resp.Body.Close()

	data,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println(err)
		statusCode = resp.StatusCode
		return
	}
	content = string(data)
	statusCode = resp.StatusCode
	return
}

//存储一页的段子到文件中
func (this *Spider) Store_dz_to_file(titles []string,contents []string) {
	filename := "MyDuanzi.txt"
	f,err := os.OpenFile(filename,os.O_CREATE | os.O_APPEND | os.O_RDWR,0644)
	if err != nil {
		fmt.Println("open file error!!!!!")
		return
	}
	defer f.Close()

	for i := 0 ; i < len(titles) ; i++{
		f.WriteString("\n=====================\n")
		f.WriteString(titles[i])
		f.WriteString("\n=====================\n")
		f.WriteString(contents[i])
	}
}

// 爬取一页段子的方法
func (this *Spider) SpiderOneDz(url string)(dz_title string,dz_content string) {

	fmt.Println("正在爬取",url)

	content,rcode := this.HttpGet(url)
	if rcode < 0 {
		fmt.Println("http get error rcode = " , rcode)
		return "",""
	}
	//得到标题
	title_exp := regexp.MustCompile(`<h1 class = "title">(?s:(.*?))</h1>`)
	titles := title_exp.FindAllStringSubmatch(content,-1)
	for _,title := range titles{
		dz_title = title[1]
		break
	}
	//得到内容
	content_exp := regexp.MustCompile(`</a></p>(?s:(.*?))<div class = "ad610">`)
	contents := content_exp.FindAllStringSubmatch(content,-1)
	for _,content_dz := range contents{
		dz_content = content_dz[1]
		dz_content = strings.Replace(dz_content,"/r/n","/n",-1)
		strings.Replace(dz_content,"<p>","",-1)
		strings.Replace(dz_content,"</p>","",-1)
		break
	}
	return
}

// 爬取一页段子的方法
func (this *Spider) SpiderOnePage() {
	fmt.Println("正在爬取第",this.page,"页")
	url := ""
	if this.page == 1 {
		url = "http://www.neihan8.com/article/index.html"
	}else {
		url = "http://www.neihan8.com/article/index_" + strconv.Itoa(this.page) + ".html"
	}
	fmt.Println("url = " , url)

	content , rcode := this.HttpGet(url)
	if rcode < 0  {
		fmt.Println("http get error rcode = " , rcode)
		return
	}
	//当前页面的段子的标题和内容
	title_slice := make([]string ,0)
	content_slice := make([]string ,0)

	dz_url_exp := regexp.MustCompile(`<h3><a href="(?s:(.*?))"`)
	urls := dz_url_exp.FindAllStringSubmatch(content,-1)
	for _,dz_url := range urls{
		//fmt.Println("dz_url = " , dz_url[1])
		full_url := "http://www.neihan8.com" + dz_url[1]
		//fmt.Println(full_url)
		//爬取一个段子
		title,content := this.SpiderOneDz(full_url)
		//fmt.Println("title = " , title)
		//fmt.Println("content = " , content)
		title_slice = append(title_slice,title)
		content_slice = append(content_slice,content)
	}
	//把当前页面的爬取到的全部段子存入文件里
	this.Store_dz_to_file(title_slice,content_slice)
}

// 与用户交互的方法
func (this *Spider) Dowork() {
	fmt.Println("Spider begin to work")
	this.page = 1
	var cmd string
	for{
		fmt.Println("请输入任意键 爬取下一页 输入exit退出")
		fmt.Scanf("%s",&cmd)
		if cmd == "exit" {
			fmt.Println("exit")
			break
		}
		fmt.Println("正在爬取第",this.page,"页")
		this.SpiderOnePage()
		this.page++
	}

}

func main()  {
	sp:= new(Spider)
	sp.Dowork()
}
