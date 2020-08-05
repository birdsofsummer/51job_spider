const R=require('ramda')
const utility =require('utility')
const dif=R.differenceWith((a,b)=>a.id==b.id)


const test=()=>{
    d=fs.readdirSync('./data/').filter(x=>/list/.test(x)).slice(0,2)
   // d.sort()
    d1=dif(...d.map(x=>require('./data/'+x)))
    file_name="./data/dif.json"
    utility.writeJSON(file_name,d1)
}

test()
