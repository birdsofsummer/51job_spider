// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    topLevel, err := UnmarshalTopLevel(bytes)
//    bytes, err = topLevel.Marshal()

package models

import (
	"encoding/json"
	"os"
	"fmt"
)

type Joblist []EngineSearchResultElement


func (d Joblist) Save(n string){
	//b,_:=json.Marshal(d)
	b, _ := json.MarshalIndent(d, "", "\t")
	file,_:=os.Create(n)
	defer file.Close()
	_, err := file.Write(b)
	if err!=nil{
		fmt.Printf("[saved] error %s",n)
	}
	fmt.Printf("[saved]: %s \n",n)
}



func Unmarshal(data []byte) (JobList, error) {
	var r JobList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *JobList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type JobList struct {
	TopAds                []interface{}               `json:"top_ads"`               
	AuctionAds            []interface{}               `json:"auction_ads"`           
	MarketAds             []interface{}               `json:"market_ads"`            
	EngineSearchResult    []EngineSearchResultElement `json:"engine_search_result"`  
	JobidCount            string                      `json:"jobid_count"`           
	BannerAds             string                      `json:"banner_ads"`            
	IsCollapseexpansion   string                      `json:"is_collapseexpansion"`  
	CoAds                 []interface{}               `json:"co_ads"`                
	KeywordRecommendation KeywordRecommendation       `json:"keyword_recommendation"`
	SearchCondition       SearchCondition             `json:"search_condition"`      
	SearchedCondition     string                      `json:"searched_condition"`    
	CurrPage              string                      `json:"curr_page"`             
	TotalPage             string                      `json:"total_page"`            
	KeywordAds            []interface{}               `json:"keyword_ads"`           
	JobSearchAssistance   []JobSearchAssistance       `json:"job_search_assistance"` 
	SEOTitle              string                      `json:"seo_title"`             
	SEODescription        string                      `json:"seo_description"`       
	SEOKeywords           string                      `json:"seo_keywords"`          
}

type EngineSearchResultElement struct {
	Type              Type            `json:"type"`              
	Jt                string          `json:"jt"`                
	Tags              []string        `json:"tags"`              
	AdTrack           string          `json:"ad_track"`          
	Jobid             string          `json:"jobid"`             
	Coid              string          `json:"coid"`              
	Effect            string          `json:"effect"`            
	IsSpecialJob      string          `json:"is_special_job"`    
	JobHref           string          `json:"job_href"`          
	JobName           string          `json:"job_name"`          
	JobTitle          string          `json:"job_title"`         
	CompanyHref       string          `json:"company_href"`      
	CompanyName       string          `json:"company_name"`      
	ProvidesalaryText string          `json:"providesalary_text"`
	Workarea          string          `json:"workarea"`          
	WorkareaText      string          `json:"workarea_text"`     
	Updatedate        string          `json:"updatedate"`        
	IsIntern          string          `json:"isIntern"`          
	Iscommunicate     string          `json:"iscommunicate"`     
	CompanytypeText   string          `json:"companytype_text"`   //CompanytypeText 
	Degreefrom        string          `json:"degreefrom"`        
	Workyear          string          `json:"workyear"`          
	Issuedate         string          `json:"issuedate"`         
	IsFromXyz         string          `json:"isFromXyz"`         
	Jobwelf           string          `json:"jobwelf"`           
	JobwelfList       []string        `json:"jobwelf_list"`      
	AttributeText     []string        `json:"attribute_text"`    
	CompanysizeText   string          `json:"companysize_text"`  //CompanysizeText
	CompanyindText    string          `json:"companyind_text"`   
	Adid              string          `json:"adid"`              
}

type JobSearchAssistance struct {
	URL       string `json:"url"`      
	Img       string `json:"img"`      
	Txt       string `json:"txt"`      
	Vtxt      string `json:"vtxt"`     
	Startdate string `json:"startdate"`
	Enddate   string `json:"enddate"`  
	Indexform string `json:"indexform"`
	Isdefault string `json:"isdefault"`
}

type KeywordRecommendation struct {
	Title    string  `json:"title"`    
	DataType string  `json:"data_type"`
	Keyword  string  `json:"keyword"`  
	Data     []Datum `json:"data"`     
}

type Datum struct {
	Href  string `json:"href"` 
	Text  string `json:"text"` 
	Click string `json:"click"`
}

type SearchCondition struct {
	Lang           string `json:"lang"`          
	Keywordtype    string `json:"keywordtype"`   
	OrdField       string `json:"ord_field"`     
	Jobarea        string `json:"jobarea"`       
	CurrPage       string `json:"curr_page"`     
	District       string `json:"district"`      
	Dibiaoid       string `json:"dibiaoid"`      
	Postchannel    string `json:"postchannel"`   
	Reservechannel string `json:"reservechannel"`
	Issuedate      string `json:"issuedate"`     
	Providesalary  string `json:"providesalary"` 
	Degreefrom     string `json:"degreefrom"`    
	Companysize    string `json:"companysize"`   
	Cotype         string `json:"cotype"`        
	Workyear       string `json:"workyear"`      
	Industrytype   string `json:"industrytype"`  
	Funtype        string `json:"funtype"`       
	Jobterm        string `json:"jobterm"`       
	Keyword        string `json:"keyword"`       
	Welfare        string `json:"welfare"`       
	Address        string `json:"address"`       
	Line           string `json:"line"`          
	Confirmdate    string `json:"confirmdate"`   
	Radius         string `json:"radius"`        
	Lonlat         string `json:"lonlat"`        
}

type CompanysizeText string
const (
	The10000人以上 CompanysizeText = "10000人以上"
	The10005000人 CompanysizeText = "1000-5000人"
	The150500人 CompanysizeText = "150-500人"
	The5001000人 CompanysizeText = "500-1000人"
	The50150人 CompanysizeText = "50-150人"
	少于50人 CompanysizeText = "少于50人"
)

type CompanytypeText string
const (
	上市公司 CompanytypeText = "上市公司"
	合资 CompanytypeText = "合资"
	国企 CompanytypeText = "国企"
	外资欧美 CompanytypeText = "外资（欧美）"
	外资非欧美 CompanytypeText = "外资（非欧美）"
	民营公司 CompanytypeText = "民营公司"
)

type Type string
const (
	EngineSearchResult Type = "engine_search_result"
)


/*
22: 	TopAds                []interface{}               `json:"top_ads"`               
23: 	AuctionAds            []interface{}               `json:"auction_ads"`           
24: 	MarketAds             []interface{}               `json:"market_ads"`            
29: 	CoAds                 []interface{}               `json:"co_ads"`                
35: 	KeywordAds            []interface{}               `json:"keyword_ads"`           
*/
