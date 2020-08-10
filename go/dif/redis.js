const { promisify } = require("util");
let bird=require("bluebird")
let redis=bird.promisifyAll(require("redis"));
//redis = require("redis");
let c= redis.createClient();

const k="black"
const add=(d=[])=>c.SADDAsync(k,d)
const del=(d=[])=>c.SREMAsync(k,d)
const list=()=>c.smembersAsync(k)

module.exports={
    add,
    del,
    list,
}

