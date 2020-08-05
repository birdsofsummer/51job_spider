const fs = require("fs");
const cheerio = require("cheerio");
const superagent=require("superagent")
const charset = require("superagent-charset");
const utility =require('utility')
const moment=require('moment')

charset(superagent)
//superagent.buffer[mime] = true

const config=require("./config")

const today=()=>moment().format('YYYYMMDD_hhmmss')

const format_url=(page=1)=>{
   let u="https://search.51job.com/list/040000,000000,0000,00,9,99," + config.keyword + ",2,"+page+".html"
  //u="https://search.51job.com/list/040000,000000,0000,00,9,99,%25E5%2589%258D%25E7%25AB%25AF,2,4.html"
   let u1=new URL(u)
   let d={
        'lang' : 'c',
        'postchannel' : '0000',
        'workyear' : '99',
        'cotype' : '99',
        'degreefrom' : '99',
        'jobterm' : '99',
        'companysize' : '99',
        'ord_field' : '0',
        'dibiaoid' : '0',
        'line' : '',
        'welfare' : '',
    }
   let q = new URLSearchParams(d)
    q.forEach((v,k)=>u1.searchParams.append(k,v))
    return u1
}


const get=(u="")=>(superagent
      .get(u)
      .set({...config.headers,referrer:u.toString()})
      .charset('gbk')
      //.buffer(false)
      .then(x=>x.text))




const get_page=(n=1)=>get(format_url(n))
const get_detail=(u="")=>get(u).then(parse_detail)


const map=(fn,x)=>{
    let r=[]
    for (let i=0;i<x.length;i++){
        let o=fn(x.eq(i))
        r.push(o)
    }
    return r
}

const parse_table=(x)=>{
  let  id=x.find('.t1 input').attr('value')
  let  url=x.find('.t1 a').attr('href')
  let  title=x.find('.t1 a').attr('title')

  let  company_url=x.find('.t2 a').attr('href')
  let  company_title=x.find('.t2 a').attr('title')
  let  address=x.find('.t3').text()
  let  money=x.find('.t4').text()
  let  date=x.find('.t5').text()
  return {
        id,
        url,
        title,
        address,
        company_title,
        company_url,
        address,
        money,
        date,
  }
}


const parse_total=(html="")=>{
    const $=cheerio.load(html)
    const resultList=$('#resultList')
    //title=map(x=>x.text(),resultList.find('.title span'))

    const total_page = +$('#hidTotalPage').attr('value')
    const total = +$('.dw_tlc .rt').eq(0).text().trim().match(/\d+/)?.[0]
    return { total , total_page}
}

const parse_detail=(html="")=>{
    const $=cheerio.load(html)

    const title=$('.cn h1').attr('title')
    const id=$('#hidJobID').attr('value')
    const money=$('.cn strong').text()
    const msg=$('.msg.ltype').text().trim().split('|').map(x=>x.trim())
    const tag=map(x=>x.text(),$('.t1 span'))
    const content=$('.job_msg.inbox').text().trim()
    const zhineng=map(x=>x.text(),$('.el.tdn'))
    const company_name=$('.catn').text()
    const address=($('.bmsg .icon_b.i_map').attr('onclick')||"").match(/\, '(.+)\'/)?.[1] ?? ""

    return {
        title,
        id,
        money,
        msg,
        tag,
        content,
        zhineng,
        company_name,
        company_url:$('.catn').attr('href'),
        company_all:$('.icon_b.i_house').attr('href'),
        company_logo:$('.com_name.himg img').attr('src'),
        company_tag:map(x=>x.attr('title'),$('.com_tag p')),
        address,
    }
}

const list_parser=(html="")=>{
    const $=cheerio.load(html)
    const resultList=$('#resultList')
    //title=map(x=>x.text(),resultList.find('.title span'))

    if (resultList.length > 0) {
       let d=map(parse_table,resultList.find('.el').slice(1))
       return d
    }else{
       return []
    }
}


const sleep=(n=1000)=>new Promise((x,y)=>setTimeout(x,n))



const start_list=async (to=3,from=1)=>{
    let n=from
    const first=await  get_page(n)
    const { total , total_page} = parse_total(first)
    const r= list_parser(first)
    const to1= to ? Math.min(to,total_page) : total_page
    console.log('%d/%d',to1,total_page)
    for (n=from+1;n<=to1;n++){
        console.log('start...%d/%d',n,to1)
        let t=await get_page(n)
        let d=list_parser(t)
        console.log("page %d",n,d)
        r.push(...d)
        console.log('done...%d/%d',n,to1)
        await sleep(100)
    }
    let file_name="./data/list_"+today()+'.json'
    console.log('all done',file_name)
    utility.writeJSON(file_name,r)
    return r
}

const get_details=async (u=[])=>{
    let r=[]
    let to=100 //u.length
    for (let i=0;i<to;i++){
        console.log('start...%d/%d',i,to)
        let d=await get_detail(u[i].url)
        let d1={list:u[i],detail:d}
        console.log("%d/%d",i,to,d1)
        r.push(d1)
        console.log('done...%d/%d',i,to)
        await sleep(1000)
    }
    let file_name='./data/detail_'+today()+'.json'
    console.log('all done',file_name)
    utility.writeJSON(file_name,r)
    return r
}

const start_detail=async (name="./data/list_20200716_050120.json")=>{
    let d1=await utility.readJSON(name)
    get_details(d1)
}



start_list(1e3)
//start_detail("./data/list_20200717_072729.json")
