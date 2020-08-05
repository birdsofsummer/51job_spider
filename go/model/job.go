package model

import  (
	"encoding/json"
	"os"
	"fmt"
)

type Job struct {
	Id           string `json:"id" xorm:"not null pk VARCHAR(100)"`
	Url          string `json:"url" xorm:"not null VARCHAR(100)"`
	Title        string `json:"title" xorm:"not null VARCHAR(100)"`
	Address      string `json:"address" xorm:" VARCHAR(100)"`
	CompanyTitle string `json:"company_title" xorm:"not null VARCHAR(100)"`
	CompanyUrl   string `json:"company_url" xorm:"not null VARCHAR(100)"`
	Money        string `json:"money" xorm:"VARCHAR(100)"`
	Date         string `json:"date" xorm:"not null VARCHAR(100)"`
}


type Jobs []Job

func (d Jobs) Save(n string){
	//b,_:=json.Marshal(d)
	b, _ := json.MarshalIndent(d, "", "\t")
	file,_:=os.Create(n)
	defer file.Close()
	_, err := file.Write(b)
	if err!=nil{
		fmt.Println("[saved] error")
	}
	fmt.Println("[saved]:",n)
}

