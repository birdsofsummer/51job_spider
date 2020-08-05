// http://i.51job.com/userset/my_apply.php?lang=c
// http://i.51job.com/userset/my_apply.php?lang=c&type=sh&tagType=all&page=2


say=(x="")=>(y="")=>console.log(x,y)

init_db=(db)=>{
  var objectStore = db.createObjectStore('t', {
      keyPath: 'id' ,
      autoIncrement: true,

  });
  objectStore.createIndex('name', 'name', { unique: true });
}


db=null


request = window.indexedDB.open("html", 1);
request.onsuccess=e=>{
    say('open db')(e)
    db=request.result
}
request.onerror = function(event) {
    say('eee')(event)
};
request.onupgradeneeded = e=>{
    db=e.target.result
    db.onclose=say('close')
    db.onabort=say('abort')
    db.onerror=say('error')
    db.onversionchange=say('v-change')
    init_db(db)
}
//request.onsuccess=e=>db=e.target.result;

save=(d={})=>db.transaction(['t'], 'readwrite').objectStore('t').add(d)
saves=(d=[])=>Promise.all(d.map(save))

readAll=()=>new Promise((a,b)=>{
  let r=[]
  var objectStore = db.transaction('t').objectStore('t');
   objectStore.openCursor().onsuccess = function (event) {
     var cursor = event.target.result;
     if (cursor) {
         r.push(cursor.value)
         cursor.continue();
     } else {
         console.log('done');
     }
    a(r)
  }
})


//-----------------------------------------------------------
get_co=()=>{
    e=[...document.querySelectorAll('.exmsg .e')]
    let co=e.map(a=>{
        let l1=a.querySelector('.l1')
        let l2=a.querySelector('.l2 .rq span')?.innerText
        let l3=a.querySelector('.l3 span.c_orange')?.innerText
        let l4=a.querySelector('.l4 .view span')?.innerText
        let r=a.querySelector('.l4 div.d.now')?.innerText
        let cvlogid=a.querySelector('[name="cvlogid"]')?.value
        let gs=l1.querySelector('a.gs')
        let zhn=l1.querySelector('a.zhn')
        return {
            name:   gs?.innerText,
            url:    gs?.href,
            title:  zhn?.innerText,
            url1:   zhn?.href,
            s:   l1.querySelector('span.xz')?.title,
            time:    l2,
            qty:    l3,
            status:    l4,
            status1:    r,
            cvlogid,
        }
    })
    .filter(x=>x?.name)
    return co
}


// ({"lastlogin":"1\u5c0f\u65f6\u524d","cycle":"1\u5929"})
hr=(cvlogid="13247146459")=>{
      u="http://i.51job.com/userset/ajax/apply_hr.php"
      t=Math.random()
      let q={
        _:(new Date).getTime(),
        [t]:"",
        cvlogid,
      }
      u1=u + "?" + new URLSearchParams(q)
     return fetch(u1)
        .then(x=>x.text())
        .then(x=>x.slice(1,-1))
        .then(JSON.parse)
}

start=async ()=>{
     let d=get_co()
     for (let i of d) {
       let z=await hr(i.cvlogid)
        Object.assign(i,z)
     }
     console.log('get done!',d)
     return d
}




/*
https://jobs.51job.com/shenzhen-lhq/123256674.html?s=01&t=0

<div class="op">
            <a track-type="jobsButtonClick" event-type="1" class="but_sq" id="app_ck" href="javascript:void(0);" onclick="delivery('hidJobID', '1', '//i.51job.com', 'c', 'search.51job.com', '01', '01', '//img06.51jobcdn.com');return false;">	<img src="//img02.51jobcdn.com/im/jobs/but_sq_arr.png" alt="" width="19" height="24">申请职位</a>
            <a track-type="jobsButtonClick" event-type="4" class="icon_b i_upline" id="123256674" target="_blank" href="//i.51job.com/userset/bounce_window_redirect.php?jobid=123256674&amp;redirect_type=2">竞争力分析</a>
            <a track-type="jobsButtonClick" event-type="3" class="icon_b i_collect" href="javascript:void(0);" onclick="saveCollection('123256674');return false;">收藏</a>
                        <div class="clear"></div>
        </div>




delivery('hidJobID', '1', '//i.51job.com', 'c', 'search.51job.com', '01', '01', '//img06.51jobcdn.com')


{
	"jobid": "(123256674:0)",
	"prd": "search.51job.com",
	"prp": "01",
	"cd": "jobs.51job.com",
	"cp": "01",
	"resumeid": "",
	"cvlan": "",
	"coverid": "",
	"qpostset": "",
	"elementname": "hidJobID",
	"deliverytype": "1",
	"deliverydomain": "//i.51job.com",
	"language": "c",
	"imgpath": "//img06.51jobcdn.com"
}
*/

apply=(jobid='(121733912:0)')=>{
    let u="https://i.51job.com/delivery/delivery.php"
    let t=Math.random()
    let q={
        'rand' : t,
        //'jsoncallback' : 'jQuery183029066898483028014_1588062627057',
        jobid,
        'prd' : 'search.51job.com',
        'prp' : '01',
        'cd' : 'jobs.51job.com',
        'cp' : '01',
        'resumeid' : '',
        'cvlan' : '',
        'coverid' : '',
        'qpostset' : '',
        'elementname' : 'hidJobID',
        'deliverytype' : '1',
        'deliverydomain' : '//i.51job.com',
        'language' : 'c',
        'imgpath' : '//img02.51jobcdn.com',
        _:(new Date).getTime(),
    }

    u1=u + "?" + new URLSearchParams(q)
    return fetch(u1)
        .then(x=>x.text())
        .then(x=>x.slice(1,-1))
        .then(JSON.parse)
}


save1=async ()=>{
    z=await readAll()
    n=z.map(y=>y.name)
    s=JSON.stringify(n)
    console.log(s)
    localStorage.black=s
}



start()
    .then(saves)
    .then(say("save done"))
    .then(readAll)
    .then(say('read\n'))



