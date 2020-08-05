package models

type Job struct {
	Id           string `json:"id" xorm:"not null pk VARCHAR(100)"`
	Url          string `json:"url" xorm:"not null VARCHAR(100)"`
	Title        string `json:"title" xorm:"not null VARCHAR(100)"`
	Address      string `json:"address" xorm:"not null VARCHAR(100)"`
	CompanyTitle string `json:"company_title" xorm:"not null VARCHAR(100)"`
	CompanyUrl   string `json:"company_url" xorm:"not null VARCHAR(100)"`
	Money        string `json:"money" xorm:"not null VARCHAR(100)"`
	Date         string `json:"date" xorm:"not null VARCHAR(100)"`
}
