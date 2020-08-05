# 51job spider

## function

+ job list
+ job detail
+ my apply history (cookie)


## 2020-7-30 改版

+ 前后端分离了 (vue+jquery。。。)
+ 一堆cookie...
+ 有反爬


###  部分接口

```javascript
/*
console.log(App)

fetchData: function () {
    var t = this;
    this.loading = !0,
    window.PageNp && window.PageNp.start(),
    Qt.requestHandler && Qt.requestHandler.abort(),
    Qt.requestHandler = $.getJSON(this.$route.fullPath, (function (e) {
        qt.log('请求结束', e),
        t.selectedItemList = [
        ],
        t.totalCnt = parseInt(e.jobid_count),
        t.totalPage = parseInt(e.total_page),
        t.currentPage = parseInt(e.curr_page),
        t.PARAMTEXT = e.searched_condition,
        t.top_ads = e.top_ads || [
        ],
        t.auction_ads = e.auction_ads || [
        ],
        t.market_ads = e.market_ads || [
        ],
        t.engine_search_result = e.engine_search_result || [
        ],
        Qt.coAds.info = e.co_ads || [
        ],
        Qt.keyword_ads.data = e.keyword_ads || [
        ],
        Qt.keyword_recommendation.data = e.keyword_recommendation || [
        ];
        var i = {
        };
        Object.keys(Qt.queryParam.data).map((function (t) {
            i[t] = e.search_condition[t]
        }));
        var n = {
        };
        Object.keys(Qt.fileParam.data).map((function (t) {
            n[t] = e.search_condition[t]
        })),
        Qt.queryParam.data = i,
        Qt.fileParam.data = n,
        document.title = e.seo_title,
        t.loading = !1,
        t.$root.$emit('fetchDataEnd', e),
        window.PageNp && window.PageNp.done()
    }))
}


*/

headers={
    "User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:80.0) Gecko/20100101 Firefox/80.0",
    "Accept": "application/json, text/javascript, */*; q=0.01",   /////不可少！！
    "Accept-Language": "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2",
    "X-Requested-With": "XMLHttpRequest",
    "Pragma": "no-cache",
    "Cache-Control": "no-cache",
}
// gbk
// json
// get(headers,u,q)

list=(keyword="java",page=2,)=>{
    //u="https://search.51job.com/list/040000,000000,0000,00,9,99,java,2,2.html"
    u="https://search.51job.com/list/040000,000000,0000,00,9,99," + keyword + ",2,"+page+".html"
    q={

        "lang" : "c",
        "postchannel" : "0000",
        "workyear" : "99",
        "cotype" : "99",
        "degreefrom" : "99",
        "jobterm" : "99",
        "companysize" : "99",
        "ord_field" : "0",
        "dibiaoid" : "0",
        "line" : "",
        "welfare" : "" ,
    }

    res={
      "top_ads": [],
      "auction_ads": [],
      "market_ads": [],
      "engine_search_result": [
        {
          "type": "engine_search_result",
          "jt": "0",
          "tags": [],
          "ad_track": "",
          "jobid": "112623672",
          "coid": "2282999",
          "effect": "1",
          "is_special_job": "",
          "job_href": "https://jobs.51job.com/shenzhen-ftq/112623672.html?s=01&t=0",
          "job_name": "Java开发工程师",
          "job_title": "Java开发工程师",
          "company_href": "https://jobs.51job.com/all/co2282999.html",
          "company_name": "盛视科技股份有限公司",
          "providesalary_text": "0.7-1.8万/月",
          "workarea": "040100",
          "workarea_text": "深圳-福田区",
          "updatedate": "07-31",
          "isIntern": "",
          "iscommunicate": "",
          "companytype_text": "上市公司",
          "degreefrom": "6",
          "workyear": "4",
          "issuedate": "2020-07-31 10:06:14",
          "isFromXyz": "",
          "jobwelf": "五险一金 员工旅游 年终奖金 绩效奖金 专业培训 餐饮补贴",
          "jobwelf_list": [
            "五险一金",
            "员工旅游",
            "年终奖金",
            "绩效奖金",
            "专业培训",
            "餐饮补贴"
          ],
          "attribute_text": [
            "深圳-福田区",
            "2年经验",
            "本科",
            "招4人"
          ],
          "companysize_text": "500-1000人",
          "companyind_text": "计算机硬件",
          "adid": ""
        }
      ],
      "jobid_count": "8725",
      "banner_ads": "<div class=\"mainleft s_search search_btm0\" id=\"banner_ads\">\r\n                                        <table border=0 cellspacing=0 cellpadding=4><tr>\n\t<td><a adid=\"33739135\" onmousedown=\"return AdsClick(33739135)\" href=\"https://companyadc.51job.com/companyads/ads/34/33740/33739028/index.htm\" title=\"多人行网络科技（上海）有限公司\" target=\"_blank\" onfocus=\"blur()\"><img src=\"//img05.51jobcdn.com/im/images/ads/34/33740/33739135/dr.gif\" border=\"0\" width=\"150\" height=\"60\"></a></td>\n\t<td><a adid=\"33929155\" onmousedown=\"return AdsClick(33929155)\" href=\"https://companyadc.51job.com/companyads/ads/34/33930/33929155/index.htm\" title=\"古驰（中国）贸易有限公司\" target=\"_blank\" onfocus=\"blur()\"><img src=\"//img05.51jobcdn.com/im/images/ads/34/33930/33929155/gc60.gif\" border=\"0\" width=\"150\" height=\"60\"></a></td>\n\t<td><a adid=\"33837189\" onmousedown=\"return AdsClick(33837189)\" href=\"https://companyadc.51job.com/companyads/ads/34/33838/33837032/index.htm\" title=\"施坦威钢琴亚太有限公司\" target=\"_blank\" onfocus=\"blur()\"><img src=\"//img05.51jobcdn.com/im/images/ads/34/33838/33837189/st6.gif\" border=\"0\" width=\"150\" height=\"60\"></a></td>\n\t<td><a adid=\"33929153\" onmousedown=\"return AdsClick(33929153)\" href=\"https://companyadc.51job.com/companyads/ads/34/33930/33929153/index.htm\" title=\"古驰（中国）贸易有限公司\" target=\"_blank\" onfocus=\"blur()\"><img src=\"//img05.51jobcdn.com/im/images/ads/34/33930/33929155/gc60.gif\" border=\"0\" width=\"150\" height=\"60\"></a></td>\n\t<td><a adid=\"33591245\" onmousedown=\"return AdsClick(33591245)\" href=\"https://companyadc.51job.com/companyads/ads/34/33592/33591223/index.htm\" title=\"上海亿通国际股份有限公司\" target=\"_blank\" onfocus=\"blur()\"><img src=\"//img05.51jobcdn.com/im/images/ads/34/33592/33591245/aa.gif\" border=\"0\" width=\"150\" height=\"60\"></a></td>\n\t<td><a adid=\"34122453\" onmousedown=\"return AdsClick(34122453)\" href=\"https://companyadc.51job.com/companyads/ads/35/34123/34122453/index.htm\" title=\"珑骧北方（北京）贸易有限公司上海第一分公司\" target=\"_blank\" onfocus=\"blur()\"><img src=\"//img05.51jobcdn.com/im/images/ads/35/34123/34122453/lj.gif\" border=\"0\" width=\"150\" height=\"60\"></a></td>\n</tr>\n</table>                    <table border=0 cellspacing=0 cellpadding=4>\r\n                        <tr>\r\n                            <td><a href=\"https://edu.51job.com\" target=\"_blank\" onfocus=\"blur()\"><img src=\"//img05.51jobcdn.com/im/mkt/zn/train/20200618/ad/eduad.png\" border=\"0\" width=\"150\" height=\"60\"></a></td>\r\n                        </tr>\r\n                    </table>\r\n                                </div>",
      "is_collapseexpansion": "1",
      "co_ads": [],
      "keyword_recommendation": {
        "title": "猜你喜欢",
        "data_type": "1",
        "keyword": "java",
        "data": [
          {
            "href": "https://search.51job.com/list/040000,000000,0000,00,9,99,Java%2B%25E5%25BC%2580%25E5%258F%2591,2,1.html?lang=c&postchannel=0000&workyear=99&cotype=99&degreefrom=99&jobterm=99&companysize=99&ord_field=0&dibiaoid=0&line=&welfare=",
            "text": "Java 开发",
            "click": "15\u001a1\u001a1596172916744\u001a27103578\u001a1e2f02d59dfc99476fca0a603a153ce3\u001a113.104.213.1\u001a1\u001ajava\u001aJava 开发,Java 高级,Java 实习,Java 初级,Java 中级,Java Web,Java 后台,Java 架构,J2EE,Javascript\u001aJava 开发"
          }
        ]
      },
      "search_condition": {
        "lang": "c",
        "keywordtype": "2",
        "ord_field": "0",
        "jobarea": "040000",
        "curr_page": "3",
        "district": "000000",
        "dibiaoid": "0",
        "postchannel": "0000",
        "reservechannel": "00000000",
        "issuedate": "9",
        "providesalary": "99",
        "degreefrom": "99",
        "companysize": "99",
        "cotype": "99",
        "workyear": "99",
        "industrytype": "00",
        "funtype": "0000",
        "jobterm": "99",
        "keyword": "java",
        "welfare": "",
        "address": "",
        "line": "",
        "confirmdate": "9",
        "radius": "-1",
        "lonlat": "0,0"
      },
      "searched_condition": "java(全文)+深圳",
      "curr_page": "3",
      "total_page": "175",
      "keyword_ads": [],
      "job_search_assistance": [
        {
          "url": "https://edu.51job.com/lesson_detail.php?lessonid=16201",
          "img": "//img01.51jobcdn.com/im/images/2020/mkt/0408/q1.png?1595844672",
          "txt": "HR教你优雅谈钱，3步定薪酬",
          "vtxt": "跳槽 薪资",
          "startdate": "2020-07-28",
          "enddate": "2020-12-31",
          "indexform": "",
          "isdefault": "0"
        }
      ],
      "seo_title": "【深圳,java招聘，求职】-前程无忧",
      "seo_description": "前程无忧为您提供最新最全的深圳,java招聘，求职信息。网聚全国各城市的人才信息，找好工作，找好员工，上前程无忧。",
      "seo_keywords": "找工作,求职,人才,招聘"
    }
}




//jsonp
default_choose_best=()=>{

    u="https://i.51job.com/userset/ajax/default_choose_best.php?lang=c&_="
    q={
        "jsoncallback":"",
        "lang":"c",
        "_":"1596170197096",
    }
    res="({...})"

    json= {
      "status": "1",
      "msg": "success",
      "data": [
        {
          "jobid": "121901014",
          "coid": "123056",
          "workarea": "040000",
          "cjobname": "xxx",
          "jobtype": "0",
          "ejobname": "专家",
          "cocname": "xxx公司",
          "coename": "xxx公司",
          "jobareaname": "深圳",
          "salary": "",
          "issuedate": "2020-07-29 07:03:10",
          "exercitation": "0",
          "jobwelf": "",
          "degreefrom": "本科",
          "workyear": "3-4年",
          "companytype": "06",
          "companysize": "7",
          "industrytype1": "32",
          "industrytype2": "",
          "functype1": "7201",
          "functype2": "",
          "effect": "1",
          "sValue": "01",
          "jobinfostr": "3-4年 | 本科 | 7-29发布"
        }
      ],
      "type": "1"
    }
}




logRecord=()=>{
    u='https://trace.51jingying.com/logRecord.php'
    d={
       "VerType" : "3",
        "webId" : "2",
        "logTime" : "1596170199201",
        "ip" : "1.1.2.1",
        "guid" : "1e2f02d59dfc99476fca0a603a153ce3",
        "domain" : "search.51job.com",
        "pageCode" : "10101",
        "cusParam" : "1\x1627103578\x1651job_web\x16\x160\x16",
        "vt" : "1596170200623",
        "logType" : "pageView" 
    }
}

```





