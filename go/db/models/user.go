package models

type User struct {
	Id   int    `json:"id" xorm:"not null pk autoincr comment('Id') unique INT(10)"`
	Name string `json:"name" xorm:"not null comment('NickName') unique VARCHAR(25)"`
}
