package main

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"encoding/json"
	"os"
	"io/ioutil"
    "path/filepath"
	"regexp"
	"math"
	. "./conn"
	. "./models"
	rc "./redis"
)

var println = fmt.Println

func ReadJson(file string) ([]Job,error){
	var r []Job
    data, err := ioutil.ReadFile(file)
    if err != nil {
        fmt.Println("File reading error", err)
        return r,err
    }
	err = json.Unmarshal(data, &r)
	return r, err
}





func SaveJson(file string){
    var d1 []Job
	var h []History

	d,e1:=ReadJson(file)
	if e1!=nil{
		println("fail to read json file",e1.Error())
	}

    //去重
	m:= make(map[string]Job)
	for _,v := range(d) {
		m[v.Id]=v
	}
	for v:=range(m) {
		vv:=m[v]
		d1=append(d1,vv)
		hh:=History{vv.Id,vv.Date}
		h=append(h,hh)
	}

    // 一次1000条
	pages:=int(math.Ceil(float64(len(d1))/1000))
	for i:=0;i<pages;i++{
		from:=i*1000
		to:=(i+1)*1000

		println("save page", i,"/",pages,from,"-",to)

		d2:=d1[from:to]
		h2:=h[from:to]

		if i==pages-1 {
			d2=d1[from:]
			h2=h[from:]
		}

		affected, err := Engine.Insert(d2,h2)
		if err!=nil{
			println("fail to save job to db",err.Error())
			return
		}
		println(affected, err)
	}
}

func SaveJsons(pathname string) {
    filepath.Walk(pathname, func (path string, info os.FileInfo, err error) error {
		if info.IsDir() {

		}else{
			r,_:=regexp.Compile("json")
			is_json:=r.MatchString(path)
			//fmt.Println(info.Name())
			if is_json {
				fmt.Println("save ",path)
				SaveJson(path)
			}
		}
        return nil
    })
}

func sync(){
	d:=new(Job)
	h:=new(History)
	//Engine.DropTables(d,h)
    //engine.CreateTables(d,h)
	err := Engine.Sync2(d,h)
	if err!=nil{
		println("sync job error",err.Error())
	}
}



func List() []Job{
	var d []Job
	err := Engine.Find(&d)
	if err!=nil{
		println("list job error",err.Error())
	}
    k:="job"
	for _,v:=range(d){
		_, err = rc.RC.Do("SADD", k, v.Id)
		if err != nil {
			fmt.Println(err)
		}
	}
	return d 
}

//	users := make([]*User, 10)
//	for i:=0;i<10;i++{
//		users[i] = new(User)
//		users[i].Name = "name_" + Random()
//	}



func Add1(users []User){
	affected, err := Engine.Insert(users)
	if err!=nil{
		println("eeee",err.Error())
	}
	println(affected, err)
}



func Test() {
//	err,engine:=Conn()
//	if err!=nil{
//		return
//	}
    sync()

//	SaveJson("../data/2/开发工程师_list.json")
//	SaveJsons("../data/")

//	d1:=List()
//	for _,i :=range d1 {
//		println(i)
//	}

//	d2:=rc.SMEMBERS("job")
//	for _,i :=range d2 {
//		println(i)
//	}

}

func main(){
	Test()
}
