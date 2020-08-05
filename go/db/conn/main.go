package conn

import (
	_ "github.com/go-sql-driver/mysql"
	// "github.com/lib/pq"
    "xorm.io/xorm"
	"math/rand"
	"time"
	"fmt"
)
var (
    userName  string = "root"
    password  string = "123456"
    ipAddrees string = "127.0.0.1"
    port      int    = 3306
    dbName    string = "test"
    charset   string = "utf8"
)

var Engine *xorm.EngineGroup 
var println = fmt.Println


func Random() string{
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%v", rand.Intn(254))
}
func Conn()(error, *xorm.EngineGroup){
    var err error
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddrees, port, dbName, charset)
	conns := []string{
		//"postgres://postgres:postgres@localhost:5432/chat1?sslmode=disable;", 
		dsn,
	}
    Engine, err = xorm.NewEngineGroup("mysql", conns, xorm.RoundRobinPolicy())

	if err != nil {
		println(err.Error())
		return err,Engine
	}
	//Engine.SetMapper(names.SameMapper{})
	//Engine.SetTableMapper(names.SameMapper{})
	//Engine.SetColumnMapper(names.SnakeMapper{})
	return err,Engine
}

func init(){
	Conn()
}
