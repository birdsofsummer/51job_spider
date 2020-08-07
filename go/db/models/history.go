package models

type History struct {
	Id   string `json:"id" xorm:"not null pk VARCHAR(100)"  binding:"required"`
	Date string `json:"date" xorm:"not null pk VARCHAR(100)"`
}
