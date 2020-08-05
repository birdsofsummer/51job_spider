const fs = require("fs");
const cheerio = require("cheerio");
const superagent=require("superagent")
const charset = require("superagent-charset");
const utility =require('utility')
const moment=require('moment')
 _=require('lodash')


charset(superagent)
//superagent.buffer[mime] = true

const today=()=>moment().format('YYYYMMDD_hhmmss')

const config=require("./config")
const cookie=config.cookie    ///

const agent = superagent.agent();

const get_jsonp=(u="",d={})=>( agent.get(u)
                    .query( {
                        [Math.random()]:"",
                        _:(new Date).getTime(),
                        ...d,
                     })
                   .set({
                        ...config["headers"],
                        Cookie:cookie,
                        Referer: u,
                    })
                   .charset('gbk')
                   //.type('json')
                   .then(x=>x.text)
                   .then(x=>x.slice(1,-1))
                   .then(JSON.parse)
)


const get_apply_info=(d={})=>{
    let u="http://i.51job.com/userset/ajax/apply.php"
    return get_jsonp(u,d)
}


const sleep=(n=1000)=>new Promise((x,y)=>setTimeout(x,n))
const map=(fn,x)=>{
    let r=[]
    for (let i=0;i<x.length;i++){
        let o=fn(x.eq(i))
        r.push(o)
    }
    return r
}


const parse_total=(html="")=>{
    const $=cheerio.load(html)
    const total=+$('.p_in .td').text().match(/\d+/)?.[0] ?? 0
    return total
}

const parse_line=a=>({
    time:a.find('.l2 .rq span').text(),
    title:a.find('.l1 a.zhn').text(),
    url:a.find('.l1 a.zhn').attr('href'),
    company:{
        name: a.find('.l1 a.gs').text(),
        url:a.find('.l1 a.gs').attr('href'),
    },
    aCvInfo:{
		"cvlogid": a.find('[name="cvlogid"]').attr('value'),
		"time": a.find("#sd"+a.find('[name="cvlogid"]').attr('value')).attr('value'),
		"pass": a.find("#pf"+a.find('[name="cvlogid"]').attr('value')).attr('value'),
    },
    //status:a.find('.l4 .view span').text().trim(),
    //status1:a.find('.l4 div.d.now').text().trim(),
    cvlogid:a.find('[name="cvlogid"]').attr('value'),
    money: a.find('.l1 span.xz').text(),
    qty:  +a.find('.l3 span.c_orange').text(),
})

const parse_my_apply=(html="")=>{
    const $=cheerio.load(html)
    const table=$('.exmsg .e')
    return map(parse_line,table).filter(x=>x && x.title)
}


// ({"lastlogin":"1\u5c0f\u65f6\u524d","cycle":"1\u5929"})
const hr=(cvlogid="13247146459")=>{
      let u="http://i.51job.com/userset/ajax/apply_hr.php"
      let t=Math.random()
      let q={
        _:(new Date).getTime(),
        [t]:"",
        cvlogid,
        //jsoncallback:jQuery183003546339639476981_1594883724999,
      }
      let  h={
            ...config["headers"],
            Cookie:cookie,
            Referer: u,
      }
      return  agent.get(u)
                   .query(q)
                   .set(h)
                   .charset('gbk')
                   //.type('json')
                   .then(x=>x.text)
                   .then(x=>x.slice(1,-1))
                   .then(JSON.parse)
                   .catch(console.log)

}

const my_apply=(page=1)=>{
    let u="http://i.51job.com/userset/my_apply.php"
    let d={
        'lang' : 'c',
        'type' : 'sh',
        'tagType' : 'all',
        'page' :page ,
    }
    let  h={
          ...config["headers"],
          Cookie:cookie,
          Referer: u,
    }
    return agent
        .get(u)
        .query(d)
        .set(h)
        .charset('gbk')
        .then(x=>x?.text ?? "")
        .catch(console.log)

}

const my_applys=async (t=1e3)=>{
    let first=await my_apply(1)
    let total= parse_total(first) ||  0
    let to=Math.min(t,total)
    let r=[]
    console.log('start %d',to)
    for (let i=1;i<=to;i++){
        console.log('start...%d/%d',i,to)
        let d=await my_apply(i)
        let d1=parse_my_apply(d)

        const aCvInfo = _.keyBy(_.map(d1,'aCvInfo'),'cvlogid')
        const apply_info = await get_apply_info({cvinfo:aCvInfo})
        const ainfo= apply_info.html

        for (let ii of d1) {

            let cvlogid=ii.cvlogid
            ii.aCvInfo.info=ainfo?.[cvlogid]

            console.log('get hr status',ii)
            ii.hr=await hr(ii.cvlogid)

            console.log('get hr status done',)
        }

        console.log("%d",i,d1)
        r.push(...d1)
        console.log('done...%d/%d',i,to)
        await sleep(1000)
    }
    let file_name='./data/myapply_'+today()+'.json'
    console.log('all done',file_name)
    utility.writeJSON(file_name,r)
    return r
}







my_applys()
