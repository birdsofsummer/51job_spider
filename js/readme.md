
```bash
cp config.example.js config.js
```


+ edit config.js   

```javascript

//copy the cookie from browser
const c="_ujz=***; ps=***; guid=***; partner=***; slife=***; search=***; nsearch=***; 51job=***; guid=***'"  

const config={
  "headers": {
        "User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:79.0) Gecko/20100101 Firefox/79.0",
        "Accept": "*/*",
        "Accept-Language": "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2",
        "Pragma": "no-cache",
        "Cache-Control": "no-cache"
    },
  "keyword":"java开发",  //edit the keyword
  "cookie":c,      
}



```

+ run 

```bash
# node v14
mkdir data
node index.js
# need a real cookie in config.js
node me.js  

```
