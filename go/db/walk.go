package main

import (
	"io/ioutil"
	"fmt"
//	"time"
    "path/filepath"
    "os"
	"regexp"
)


func GetAllFile(pathname string) (error){
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			//fmt.Printf(pathname+fi.Name())
			GetAllFile(pathname + fi.Name() + "/")
		} else {
			fmt.Println(fi.Name())
		}
	}
	return err
}

func GetAllFile1(pathname string){
    filepath.Walk(pathname, func (path string, info os.FileInfo, err error) error {
		if info.IsDir() {

		}else{
			r,_:=regexp.Compile("json")
			is_json:=r.MatchString(path)
			//fmt.Println(info.Name())
			if is_json {
				fmt.Println(path)
			}
		}
        return nil
    })
}
    
func test(){
	p:="../data/"
	//GetAllFile(p)
	GetAllFile1(p)
}
