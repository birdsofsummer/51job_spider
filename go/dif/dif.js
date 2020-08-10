const R=require('ramda')
const u=require('utility')
const Mustache=require("mustache")
const c=require('./redis')
const fs=require('fs')

const ls=(p="../data")=>{
    let d=fs.readdirSync(p).reverse()
    let dir=d.map(x=>p+"/"+x).flatMap(x=>fs.readdirSync(x).map(y=>x+'/'+y))
    return dir
}

const find_dif=async (z=[])=>{
    let black=await c.list()
    let [a,b]=await Promise.all(z.map(u.readJSON))
    let dif=R.differenceWith((x,y)=>x.jobid==(y.jobid ? y.jobid :  y.id),a,b)
    let d2=dif
        .filter(x=>!/千/.test(x.providesalary_text))
        .filter(x=>!black.includes(x.company_name))
        .filter(x=>!/销售|实习|就业指导|招|初级|采购|课程顾问/.test(x.job_name))
        .filter(x=>/开发|工程|全栈|前端|python|node|javascript|程|系统|架构|软件|web|dev/i.test(x.job_name))
    let ddd=R.pipe(R.sortBy(x=>x.workarea),R.addIndex(R.map)((v,i)=>(v.index=i,v)))(d2)
    let d1={
        data:ddd,
        data1:JSON.stringify(ddd),
        now:new Date(),
    }
    let template=fs.readFileSync("./dif.mustache").toString()
    let t=Mustache.render(template, d1);
    tt="/tmp/1.html"
    fs.writeFileSync(tt,t)
    console.log(ddd.length,tt)
    return ddd
}

const find_today=()=>{
    let p='../data'
    let dir=ls(p)
    let z=dir.slice(0,2)
    return find_dif(z)
}

const find_m=()=>{
    let p='../data'
    let dir=ls(p)
    let z=[dir[0],dir.slice(-1)[0]]
    return find_dif(z)
}

const creat_template=(a=[])=>{
    k={data:R.keys(a[0])}
    tag=['<%', '%>']
tmplate1=`
<%#data%>
<dl class="<%.%>">
 <dt>
     <%.%>
 </dt>
 <dd>
   {{<%.%>}}
 </dd>
</dl>
<%/data%>
 `
tmplate=`
<%#data%>
    <div class="cell <%.%>">
      {{<%.%>}}
    </div>
<%/data%>
`

    t=Mustache.render(tmplate,k,{},tag)
    file="/tmp/1.mustache"
    fs.writeFileSync(file,t)
    console.log(file)
}



find_m()
