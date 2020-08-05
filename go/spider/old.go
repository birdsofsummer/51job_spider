package spider

import (
//	"io/ioutil"
	"io"
    "fmt"
    js "github.com/dop251/goja"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"strconv"
	"regexp"
	.. "./http1"
	.. "./model"
	"../util"
)





func Get_list(keyword string,page int) (io.Reader){
	u:="https://search.51job.com/list/040000,000000,0000,00,9,99," + keyword + ",2," + fmt.Sprintf("%v", page) + ".html"
	d:=map[string]interface{}{
			"lang" : "c",
			"postchannel" : "0000",
			"workyear" : "99",
			"cotype" : "99",
			"degreefrom" : "99",
			"jobterm" : "99",
			"companysize" : "99",
			"ord_field" : "0",
			"dibiaoid" : "0",
			"line" : "",
			"welfare" : "",
		}
	r,_:=Get(u,d)
	fmt.Println(r.Status)
//	defer r.Body.Close()
//	b, _ := ioutil.ReadAll(r.Body)
	b,err:=DecodeHTMLBody(r.Body,"gbk")
	if err != nil {
		fmt.Println("ZZZZZZ")
	}
	return b
}


func Get_detail(u string)(io.Reader){
	d:=map[string]interface{}{

	}
	r,_:=Get(u,d)
	fmt.Println(r.Status)
//	defer r.Body.Close()
//	b, _ := ioutil.ReadAll(r.Body)
	b,err:=DecodeHTMLBody(r.Body,"gbk")
	if err != nil {
		fmt.Println("ZZZZZZ")
	}
	return b
}


func ParseTotal(doc *goquery.Document)(int){
	//total, _ := doc.Find("#hidTotalPage").Attr("value")
	total:= doc.Find(".td").Text() //  "共 409 页"

	reg:= regexp.MustCompile(`[0-9]+`)
	n1:=reg.FindAllString(total, -1)
	n:="0"
	if len(n1) >0 {
		n=n1[0]
	}
	total1,err := strconv.Atoi(n)
	if err !=  nil{
		total1=0
	}
	return total1
}


func ParseList(b io.Reader)(int,[]Job){
	//s1:=strings.NewReader(s)
	doc, err := goquery.NewDocumentFromReader(b)
	if err != nil {
		fmt.Println(err)
	}
	total:=ParseTotal(doc)
	resultList:=doc.Find("#resultList .el")

	var r []Job

	resultList.Each(func (i int,x *goquery.Selection){
		if i==0 {
		   return	
		}
		id,_:=x.Find(".t1 input").Attr("value")
		url,_:=x.Find(".t1 a").Attr("href")
		title,_:=x.Find(".t1 a").Attr("title")
		company_url,_:=x.Find(".t2 a").Attr("href")
		company_title,_:=x.Find(".t2 a").Attr("title")
		address:=x.Find(".t3").Text()
		money:=x.Find(".t4").Text()
		date:=x.Find(".t5").Text()

		j:=Job{
			Id:id,
			Url:url,
			Title:title,
			Address:address,
			CompanyTitle:company_title,
			CompanyUrl:company_url,
			Money:money,    
			Date:date,    
		}
		//fmt.Println(i,j)
		r=append(r,j)
	})
	return total,r
}



func test(){
	r:=`
		<input class='a' id='a' value='123' /> 
		  <p  class='z' id='b'>dddd </p>
		  <p  class='z' >eee </p>
	`
	s:=strings.NewReader(r)
	doc, err := goquery.NewDocumentFromReader(s)
	if err != nil {
		fmt.Println(err)
	}
	o, _ := doc.Find("#a").Attr("value")
	fmt.Println(o)

	doc.Find(".z").Each(func(i int, node *goquery.Selection) {
		tt:=node.Text()
		fmt.Println(tt)
		ok := strings.Contains(tt, "dd")
		fmt.Println(i,ok)
	})

	sel :=doc.Find(".z")
	for i := range sel.Nodes {
		node := sel.Eq(i)
		tt:=node.Text()
		fmt.Println(tt)
		ok := strings.Contains(tt, "dd")
		fmt.Println(i,ok)

	}

	//[]byte("ddd")
}

func test1(){
    vm := js.New() 
    r, _ := vm.RunString(`
	  1+1
    `)
	v, _ := r.Export().(int64) 
    fmt.Println(r,v)
}

func GetLists(k string,path string) {
	s:=Get_list(k,1)
	total,r:=ParseList(s)
    fmt.Println(total,r)
	c := make(chan []Job,total)
	go func (){
		to:=total
		for i:=2;i<to;i++{
			fmt.Println("start",k,i,to)
			_,r1:=ParseList(Get_list(k,i))
			c <- r1
			fmt.Println("end",k,i,to)
		}
		close(c)
	}()
	var result Jobs
    for  i := range c {
        //fmt.Println(i)
		result=append(result,i...)
    }
	fmt.Println(k,len(r)," all done")
	file_name:=path+k+"_list.json"
	result.Save(file_name) //db...
}

func start(){
	kk:=[]string{
		"前端",
		"开发工程师",
		"java开发",
		"python",
		"nodejs",
	}
	path:="../data/"+util.Today()+"/"
	util.Md(path)
	for k,v:=range(kk){
		fmt.Println(k)
		GetLists(v,path)
	}
	fmt.Println(kk,"done")
}

func main(){
	start
}
