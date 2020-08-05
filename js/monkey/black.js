// ==UserScript==
// @name     51job
// @version  1
// @grant    none
// ==/UserScript==

console.log("hello")

re=/java后|java开|java软|java高|java初|java中|Oracle|ios|安卓|android|c#|c\+\+|.net|单片机|嵌入|asp|供应商|采购|实习|应届|校招|u3d|unity|燃料|bms|mes|毕业生|硬件/gi

let re1=/龙岗|光明/


black=JSON.parse?.(localStorage.black ?? "[]")
console.log(black)

remove=()=>{
    console.log("check")

    el=[...document.querySelectorAll('#resultList div.el')].slice(1)
    e1=el.filter(x=>{
       let t=x.querySelector('.t2 a').title
       return black.includes(t)
    })

    el.forEach(x=>{

       if (x.querySelectorAll('[alt="校招"]')?.length){
            x.remove()
       }
       if (x.querySelectorAll('[alt="实习"]')?.length){
            x.remove()
       }

        let title=x.querySelector(".t1 a").title
        let area=x.querySelector('.t3').innerText
        if (re.test(title) || re1.test(area)) {
            console.log("....",title,area)
            x.remove()
        }
    })

    e1.forEach(x=>{
        let t=x.querySelector('.t2 a').title
        console.log("remove",t)
        x.remove()
    })

    console.log('done')
}

document.addEventListener("click",remove,false)
window.addEventListener("load",remove,false)

console.log("band")


onmessage=console.log

say=(x)=>(...y)=>console.log(x,...y)
window.addEventListener("message", say('msg'), false);



function receiveMessage(event) {
    //event.origin
    event.source.postMessage("--->" , event.origin);
}

postMessage('ddd',location.href)

/*
    u1="http://i.51job.com/userset/my_apply.php?lang=c"
    u2="https://search.51job.com/list/040000,000000,0000,00,9,99,%25E5%25BC%2580%25E5%258F%2591%25E5%25B7%25A5%25E7%25A8%258B%25E5%25B8%2588,2,2.html?lang=c&postchannel=0000&workyear=99&cotype=99&degreefrom=99&jobterm=99&companysize=99&ord_field=0&dibiaoid=0&line=&welfare="
*/
