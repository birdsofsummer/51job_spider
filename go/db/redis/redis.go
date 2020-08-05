package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

var  RC redis.Conn

func Conn()   (redis.Conn, error){
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return c,err
	}
	//defer c.Close()
	return c,err
}

func init (){
	var err error
	RC,err=Conn()
	if err != nil {
		fmt.Println("Connect to redis error", err)
	}
}

func newPool(server, password string) *redis.Pool {
    return &redis.Pool{
            MaxIdle: 64,
            MaxActive: 1000,
            IdleTimeout: 240 * time.Second,
            Dial: func() (redis.Conn, error) {
                c, err := redis.Dial("tcp", server)
                if err != nil {
                    return nil, err
                }
               /*
                if _, err := c.Do("AUTH", password); err != nil {
                    c.Close()
                    return nil, err
                }*/
            return c, err
        },
        TestOnBorrow: func(c redis.Conn, t time.Time) error {
        if time.Since(t) < time.Minute {
            return nil
        }
            _, err := c.Do("PING")
 
            return err
        },
    }
}

//k:="job"
func SMEMBERS(k string) []string{
	var all []string
	cursor := "0"
	for {
		res, err := redis.MultiBulk(RC.Do("SSCAN", k, cursor, "count", "1000"))
		if err != nil {
			break;
		}
		cursor, err = redis.String(res[0], err)
		val, err := redis.Strings(res[1], err)
		all = append(all, val...)
		if cursor == "0" {
			break
		}
	}
	fmt.Println("cached",all)
	return all
}







func Test() {
	c,err:=Conn()
	_, err = c.Do("SET", "mykey", "superWang")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}


    _, err = c.Do("HSet", "books", "abc", 100)
    if err != nil {
        fmt.Println(err)
        return
    }
    r, err := redis.Int(c.Do("HGet", "books", "abc"))
    if err != nil {
        fmt.Println("get abc failed,", err)
        return
    }
	fmt.Println(r)

    _, err = c.Do("MSet", "abc", 100, "efg", 300)
    if err != nil {
        fmt.Println(err)
        return
    }
	r1, err := redis.Ints(c.Do("MGet", "abc", "efg"))
    if err != nil {
        fmt.Println("get abc failed,", err)
        return
    }
    for _, v := range r1 {
        fmt.Println(v)
    }



    _, err = c.Do("expire", "abc", 10)
    if err != nil {
        fmt.Println(err)
        return
    }

    _, err = c.Do("lpush", "book_list", "abc", "ceg", 300)
    if err != nil {
        fmt.Println(err)
        return
    }
	r2, err := redis.String(c.Do("lpop", "book_list"))
    if err != nil {
        fmt.Println("get abc failed,", err)
        return
    }
    fmt.Println(r2)

}

func Test1(){

	server:="127.0.0.1:6379"
	password:=""
	p:=newPool(server, password)
	c:= p.Get()
	defer c.Close()


	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
}

func Test2(){
	c,err:=Conn()

    j:=[]string{"11","22","33"}
    k:="job"

	fmt.Println("save",j)
	for _,v:=range(j){
		_, err = c.Do("SADD", k, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	//SMEMBERS k
	var all []string
	cursor := "0"
	for {
		res, err := redis.MultiBulk(c.Do("SSCAN", k, cursor, "count", "1000"))
		if err != nil {
			break;
		}
		cursor, err = redis.String(res[0], err)
		val, err := redis.Strings(res[1], err)
		all = append(all, val...)
		if cursor == "0" {
			break
		}
	}
	fmt.Println("cached",all)



	value, err := redis.Values(c.Do("smembers", "job"))
	if err != nil {
		fmt.Println("set get members failed", err.Error())
	} else {
		fmt.Printf("cached jobs")
		for _, v := range value {
			fmt.Printf("%s ", v.([]byte))
		}
		fmt.Printf("\n")
	}
    
	for _,v:=range(all) {
		res, err := c.Do("SREM",k,v)
		if err != nil {
		}
		fmt.Println("clean",res)
	}
}





func main(){
	//Test()
	//Test1()
	//Test2()
}
