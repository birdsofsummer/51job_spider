package main
import (
	"fmt"
	"io/ioutil"
	"strconv"
	"github.com/axgle/mahonia"
	. "../http1"
	"../db/models"
	"../util"
)

func Get_list(keyword string,page int) ([]byte){
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
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("ZZZZZZ")
	}
	return b
}


func ParseList(b []byte)(models.JobList){
	d:=mahonia.NewDecoder("gbk")
	_, e, _ := d.Translate(b, true)
	r,_ := models.Unmarshal(e)
	return r

}


func GetLists(keyword string,from int,to int,path string) (models.Joblist){
	var result models.Joblist

	first:=ParseList(Get_list(keyword,1))
	t:=first.TotalPage
	fmt.Printf("total page %s \n",t)

	to1:=to
	total1,err := strconv.Atoi(t)
	if err != nil {
		to1=0
	}else{
		if (to == 0)  || (total1 < to){
			to1=total1
		}
	}

	c := make(chan models.Joblist,to1)
	go func(){
		for i:=from; i<= to1;i++{
			fmt.Printf("get page start %d/%d \n",i,to1)
			d:=ParseList(Get_list(keyword,1))
			c <- d.EngineSearchResult
			fmt.Printf("get page done %d/%d \n",i,to1)
		}
		close(c)
	}()

	for i:=range c{
		result=append(result,i...)
	}

	file_name:=path+keyword+"_list.json"
	fmt.Printf("%d %d all done \n",to1,len(result))
	fmt.Printf("save to %s \n",file_name)
	result.Save(file_name) //db...
	return result
}



func start(){
	kk:=[]string{
		"前端",
		//"开发工程师",
		//"java开发",
		//"python",
		//"nodejs",
	}
	path:="../data/"+util.Today()+"/"
	util.Md(path)
	for k,v:=range(kk){
		fmt.Println(k)
		GetLists(v,1,0,path)
		//fmt.Println(r)
	}
	fmt.Println(kk,"done")
}


func main() {
	start()
}

