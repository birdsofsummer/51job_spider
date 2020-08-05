superagent=require("superagent")
qs=require('qs')
charset = require("superagent-charset");
charset(superagent)
utility =require('utility')
moment=require('moment')


sleep=(n=1000)=>new Promise((x,y)=>setTimeout(x,n))
today=()=>moment().format('YYYYMMDD_hhmmss')


HEADERS={
        "User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:80.0) Gecko/20100101 Firefox/80.0",
        "Accept": "application/json, text/javascript, */*; q=0.01",
        "Accept-Language": "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en;q=0.3,en-US;q=0.2",
        "X-Requested-With": "XMLHttpRequest",
        "Pragma": "no-cache",
        "Cache-Control": "no-cache",
        "referrer": "",
}

get=((u="/",q={})=>superagent
    .get(encodeURI(u))
    .query(q)
    .set({
        ...HEADERS,
        referrer: encodeURI(u)+"?"+qs.stringify(q)
    })
    .charset('gbk')
   //.buffer(false)
    .type("json")
    .then(x=>x.text)
    .then(JSON.parse)
    )

get_list=(keyword="java",page=1)=>{
   let u="https://search.51job.com/list/040000,000000,0000,00,9,99,"+keyword+",2,"+page+".html"
   console.log(u)
   let q={
              lang: 'c',
              postchannel: '0000',
              workyear: '99',
              cotype: '99',
              degreefrom: '99',
              jobterm: '99',
              companysize: '99',
              ord_field: '0',
              dibiaoid: '0',
              line: '',
              welfare: '',
    }
    return get(u,q)
}

get_lists=async(keyword="java",from=1,to=Infinity)=>{
    let file_name="./data/list_"+keyword+"_"+today()+'.json'
    let first=await get_list(keyword,1)
    result=[]
    let {
        curr_page,
        total_page,
        engine_search_result,
        jobid_count,
    }=first
    let from1=Math.max(from,1)
    let total=Math.min(to,total_page)

    for (let i=from1;i<=total;i++){
        console.log("start %s %d/%d",keyword,i,total)
        try {
            let r=await get_list(keyword,i)
            let d=r.engine_search_result
            result.push(...d)
            console.log("done %s %d/%d",keyword,i,total)
            //await sleep(1e3)
        }catch(e){
            console.log(e)
            break
        }
    }
    console.log("all done %s %d->%d ,%s",keyword,from1,total,file_name)
    utility.writeJSON(file_name,result)
    return result
}

start_list=async ()=>{
    let k=[
        "前端"
    ]
    for (let i of k) {
        console.log("fetch %s",k)
        get_lists(i)
    }
}

start_list()





