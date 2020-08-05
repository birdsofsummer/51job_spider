USE `test`;
DROP TABLE IF EXISTS `job`;
DROP TABLE IF EXISTS `history`;

--  id
--  url
--  title
--  address
--  company_title
--  company_url
--  money
--  date

--	{
--		"id": "123340325",
--		"url": "https://jobs.51job.com/shenzhen-nsq/123340325.html?s=01\u0026t=0",
--		"title": "Java开发工程师",
--		"address": "深圳-南山区",
--		"company_title": "深圳市搜了网络科技股份有限公司",
--		"company_url": "https://jobs.51job.com/all/co2235461.html",
--		"money": "1.2-1.8万/月",
--		"date": "07-21"
--	}

-- job + history

--	{
--		"id": "123340325",
--		"url": "https://jobs.51job.com/shenzhen-nsq/123340325.html?s=01\u0026t=0",
--		"title": "Java开发工程师",
--		"address": "深圳-南山区",
--		"company_title": "深圳市搜了网络科技股份有限公司",
--		"company_url": "https://jobs.51job.com/all/co2235461.html",
--		"money": "1.2-1.8万/月",
--		"date": "07-21"
--		"date": "07-1,07-2,07-3,07-4,07-5,07-6,07-7,07-8,07-9"
--	}



CREATE TABLE `job` (
  `id` varchar(100) NOT NULL,
  `url` varchar(100) ,
  `title` varchar(100),
  `address` varchar(100) ,
  `company_title` varchar(100),
  `company_url` varchar(100) ,
  `money` varchar(100) ,
  `date` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8;

CREATE TABLE `history` (
  `id` varchar(100) NOT NULL,
  `date` varchar(100) NOT NULL,
  PRIMARY KEY (`id`,`date`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8;


--INSERT INTO job ( id, url, title, address, company_title, company_url, money, date) VALUES ( "1","http://","java","xxx street","a","http://a","1-2w","1");
--INSERT INTO job ( id, url, title, address, company_title, company_url, money, date) VALUES ( "2","http://","c++","xxx street","b","http://a","1-2w","1");
--INSERT INTO job ( id, url, title, address, company_title, company_url, money, date) VALUES ( "3","http://","c#","xxx street","c","http://a","1-2w","1");
--INSERT INTO job ( id, url, title, address, company_title, company_url, money, date) VALUES ( "4","http://","node","xxx street","d","http://a","1-2w","1");
--INSERT INTO job ( id, url, title, address, company_title, company_url, money, date) VALUES ( "5","http://","ios","xxx street","e","http://a","1-2w","1");
--INSERT INTO job ( id, url, title, address, company_title, company_url, money, date) VALUES ( "6","http://","android","xxx street","f","http://a","1-2w","1");
--
--INSERT INTO history ( id,date) VALUES ( "1",1);
--INSERT INTO history ( id,date) VALUES ( "1",2);
--INSERT INTO history ( id,date) VALUES ( "1",3);
--INSERT INTO history ( id,date) VALUES ( "1",4);
--INSERT INTO history ( id,date) VALUES ( "2",1);
--INSERT INTO history ( id,date) VALUES ( "3",2);
--INSERT INTO history ( id,date) VALUES ( "4",3);
--INSERT INTO history ( id,date) VALUES ( "5",4);
--INSERT INTO history ( id,date) VALUES ( "6",4);
--
--SELECT a.*, group_concat(b.date) as day 
--FROM job a 
--LEFT JOIN history b 
--ON a.id = b.id
--GROUP BY id;


----   +----+---------+---------+------------+---------------+-------------+-------+------+---------+
----   | id | url     | title   | address    | company_title | company_url | money | date | day     |
----   +----+---------+---------+------------+---------------+-------------+-------+------+---------+
----   | 1  | http:// | java    | xxx street | a             | http://a    | 1-2w  | 1    | 1,2,3,4 |
----   | 2  | http:// | c++     | xxx street | b             | http://a    | 1-2w  | 1    | 1       |
----   | 3  | http:// | c#      | xxx street | c             | http://a    | 1-2w  | 1    | 2       |
----   | 4  | http:// | node    | xxx street | d             | http://a    | 1-2w  | 1    | 3       |
----   | 5  | http:// | ios     | xxx street | e             | http://a    | 1-2w  | 1    | 4       |
----   | 6  | http:// | android | xxx street | f             | http://a    | 1-2w  | 1    | 4       |
----   +----+---------+---------+------------+---------------+-------------+-------+------+---------+
----   6 rows in set (0.00 sec)

