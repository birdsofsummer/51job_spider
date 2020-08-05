package main

import (
	_ "github.com/go-sql-driver/mysql"
	// "github.com/lib/pq"
    //"xorm.io/xorm"
	//"math/rand"
	//"time"
	"fmt"
	. "./conn"
	. "./models"
)

var println = fmt.Println

func List() []User{
	var allusers []User
	err := Engine.Find(&allusers)
	if err!=nil{
		println("oo",err.Error())
	}
	return allusers
}

func Add1(){
	users := make([]*User, 10)
	for i:=0;i<10;i++{
		users[i] = new(User)
		users[i].Name = "name_" + Random()
	}
	affected, err := Engine.Insert(users)
	if err!=nil{
		println("eeee",err.Error())
	}
	println(affected, err)
}

func sync(){
	u:=new(User)
	Engine.DropTables(u)
    //engine.CreateTables(u)
	err := Engine.Sync2(u)
	if err!=nil{
		println("eeee",err.Error())
	}
}


func Test() {
//	err,engine:=Conn()
//	if err!=nil{
//		return
//	}
    sync()
	Add1()
	u1:=List()
	for _,i :=range u1 {
		println(i.Id,i.Name)
	}
}

func main(){
	Test()
}
