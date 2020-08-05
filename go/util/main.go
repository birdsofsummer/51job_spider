package util

import (
	"os"
	"time"
	"fmt"
)


func Md(path string){
	_, err := os.Stat(path)
	if err == nil {
		fmt.Println("exist")
	}else{
		err=os.MkdirAll(path,os.ModePerm)
		if err!=nil{
		   fmt.Println(err)
		   return
	    }
		fmt.Println("created")
	}
}

func Today()(string){
	now := time.Now()
	n:=now.Format("2006-01-02")
	return n
}


