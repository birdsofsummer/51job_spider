package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/satori/go.uuid"
	"fmt"
	"time"
	//"io/ioutil"
	//"encoding/json"
	"../db/models"
)


type Myheader struct {
    Rate int `header:"Rate"`
    Domain string `header:"Domain"`
    Token string `header:"Domain"`
}

//os.Getenv
var (
	//CORS=cors.Default() 
    CORS=cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost"},
		AllowMethods:     []string{"PUT", "PATCH","GET","POST","HEAD"},
		AllowHeaders:     []string{"Origin","Token"},
		ExposeHeaders:    []string{"Content-Length","Token"},
		AllowCredentials: false,
		AllowOriginFunc: func(origin string) bool {
			//return origin == "https://github.com"
			return true
		},
		MaxAge: 12 * time.Hour,
	})
)




func get_history(c *gin.Context){
	id:= c.Params.ByName("id")
	h:=models.History{
		Id:id,
		Date:"08-01",
	}
	c.JSON(http.StatusOK, h)
}

func list_history(c *gin.Context){
	h:=models.History{
		Id:"123",
		Date:"08-01",
	}
	h1:=make([]models.History,0)
	h1=append(h1,h)
	c.JSON(http.StatusOK, h1)
}

//echo
func add_history(c *gin.Context){
	var r models.History
	err := c.ShouldBindJSON(&r)
	if err!=nil {
		//c.JSON(http.StatusOK, r)
		c.JSON(http.StatusBadRequest,gin.H{
			"error_msg": err.Error(),
			"data":r,
		})
	}else{
		c.JSON(http.StatusOK, r)
	}
}

func token_middleware() gin.HandlerFunc {
	return func (c *gin.Context){
				//	roomid := c.Param("roomid")
				//	nick := c.Query("nick")
				//	message := c.PostForm("message")
				//  body,_ := ioutil.ReadAll(c.Request.Body)
				//  s:=string(body) 
				//  fmt.Println("bbbbb",s)
				//	err := json.Unmarshal(body, &r)

				for k,v :=range c.Request.Header {
					fmt.Println(k,v)
				}

				ip:=c.ClientIP()
				fmt.Println("client",ip)

				///t := c.Request.FormValue("token")
				t := c.GetHeader("token")
				if t == "" {
				   c.JSON(401,gin.H{
					   "error_message":"token?",
				   })
				   c.Abort()
				   return
				}

				u1 := uuid.Must(uuid.NewV4())
				u2:=fmt.Sprintf("%s", u1)

				c.Header("token",t)
				c.Writer.Header().Set("X-Revision", "111")
				c.Writer.Header().Set("X-Request-Id", u2)
				c.Next()
			}
}

//curl localhost:8080/upload -i -H "token:123ww"  -X POST -H Content-Type:multipart/form-data  -F file=@main.go
func upload(context *gin.Context) {
      header, _ := context.FormFile("file")
      path := "/tmp/" + header.Filename 
      fmt.Println(path)
      err := context.SaveUploadedFile(header, path)
      if err != nil {
         fmt.Println(err.Error())
      }
      fmt.Println(header.Filename," ",header.Size," ",header.Header)
      context.JSON(200,gin.H{
         "fileName":path,
      })
    }

func uploads(c *gin.Context) {
      form, _ := c.MultipartForm()
      files := form.File["file[]"]
      for _,file := range files {
         fmt.Println(file.Filename)
         c.SaveUploadedFile(file,"/tmp/"+file.Filename)
      }
}
func ping (c *gin.Context) {
		c.String(http.StatusOK, "pong")
}



func log_format(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[FORMATTER TEST] %v | %3d | %13v | %15s | %-7s %s\n%s",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Path,
			param.ErrorMessage,
		)
}




var db = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()
	//r.MaxMultipartMemory = 8 << 20 // 8 MiB
    r.MaxMultipartMemory = 100 << 20 // 设置最大上传大小为100M
	r.Use(CORS)
	r.Use(token_middleware())
	r.Use(gin.LoggerWithFormatter(log_format))


	//r.Static("/", "./public")
	r.GET("/ping", ping)
	r.GET("/h/:id", get_history)
	r.GET("/h", list_history)
	r.POST("/h", add_history)
    r.POST("/upload", upload)
    r.POST("/uploads", uploads)



	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})





	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
