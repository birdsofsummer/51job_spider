## gin

```go
func (engine *Engine) allocateContext() *Context 
func (engine *Engine) Delims(left, right string) *Engine 
func (engine *Engine) SecureJsonPrefix(prefix string) *Engine 
func (engine *Engine) LoadHTMLGlob(pattern string) 
func (engine *Engine) LoadHTMLFiles(files ...string) 
func (engine *Engine) SetHTMLTemplate(templ *template.Template) 
func (engine *Engine) SetFuncMap(funcMap template.FuncMap) 
func (engine *Engine) NoRoute(handlers ...HandlerFunc) 
func (engine *Engine) NoMethod(handlers ...HandlerFunc) 
func (engine *Engine) Use(middleware ...HandlerFunc) IRoutes 
func (engine *Engine) rebuild404Handlers() 
func (engine *Engine) rebuild405Handlers() 
func (engine *Engine) addRoute(method, path string, handlers HandlersChain) 
func (engine *Engine) Routes() (routes RoutesInfo) 
func (engine *Engine) Run(addr ...string) (err error) 
func (engine *Engine) RunTLS(addr, certFile, keyFile string) (err error) 
func (engine *Engine) RunUnix(file string) (err error) 
func (engine *Engine) RunFd(fd int) (err error) 
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) 
func (engine *Engine) HandleContext(c *Context) 
func (engine *Engine) handleHTTPRequest(c *Context) 


func Bind(val interface{}) HandlerFunc 
	return func(c *Context) 
// WrapF is a helper function for wrapping http.HandlerFunc and returns a Gin middleware.
func WrapF(f http.HandlerFunc) HandlerFunc 
	return func(c *Context) 
// WrapH is a helper function for wrapping http.Handler and returns a Gin middleware.
func WrapH(h http.Handler) HandlerFunc 
	return func(c *Context) 
func (h H) MarshalXML(e *xml.Encoder, start xml.StartElement) error 



func assert1(guard bool, text string) 
func filterFlags(content string) string 
func chooseData(custom, wildcard interface{}) interface{} 
func parseAccept(acceptHeader string) []string 
func lastChar(str string) uint8 
func nameOfFunction(f interface{}) string 
func joinPaths(absolutePath, relativePath string) string 
func resolveAddress(addr []string) string 
```




## router

```go

func LoadHTMLGlob(pattern string) 
func LoadHTMLFiles(files ...string) 
func SetHTMLTemplate(templ *template.Template) 
func NoRoute(handlers ...gin.HandlerFunc) 
func NoMethod(handlers ...gin.HandlerFunc) 
func Group(relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup 
func Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes 
func POST(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes 
func GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes 
func DELETE(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes 
func PATCH(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes 
func PUT(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes 
func OPTIONS(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes 
func HEAD(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes 
func Any(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes 
func StaticFile(relativePath, filepath string) gin.IRoutes 
func Static(relativePath, root string) gin.IRoutes 
func StaticFS(relativePath string, fs http.FileSystem) gin.IRoutes 
func Use(middlewares ...gin.HandlerFunc) gin.IRoutes 
func Routes() gin.RoutesInfo 
func Run(addr ...string) (err error) 
func RunTLS(addr, certFile, keyFile string) (err error) 
func RunUnix(file string) (err error) 
func RunFd(fd int) (err error) 

```




## context

```go

func (c *Context) reset() 
func (c *Context) Copy() *Context 
// this function will return "main.handleGetUsers".
func (c *Context) HandlerName() string 
func (c *Context) HandlerNames() []string 
func (c *Context) Handler() HandlerFunc 
func (c *Context) Next() 
func (c *Context) IsAborted() bool 
func (c *Context) Abort() 
func (c *Context) AbortWithStatus(code int) 
func (c *Context) AbortWithStatusJSON(code int, jsonObj interface{}) 
func (c *Context) AbortWithError(code int, err error) *Error 
func (c *Context) Error(err error) *Error 

//挂全局变量
func (c *Context) Set(key string, value interface{}) 
func (c *Context) Get(key string) (value interface{}, exists bool) 
func (c *Context) MustGet(key string) interface{} 
func (c *Context) GetString(key string) (s string) 
func (c *Context) GetBool(key string) (b bool) 
func (c *Context) GetInt(key string) (i int) 
func (c *Context) GetInt64(key string) (i64 int64) 
func (c *Context) GetFloat64(key string) (f64 float64) 
func (c *Context) GetTime(key string) (t time.Time) 
func (c *Context) GetDuration(key string) (d time.Duration) 
func (c *Context) GetStringSlice(key string) (ss []string) 
func (c *Context) GetStringMap(key string) (sm map[string]interface{}) 
func (c *Context) GetStringMapString(key string) (sms map[string]string) 
func (c *Context) GetStringMapStringSlice(key string) (smss map[string][]string) 

// 解析input params|query|body
//     router.GET("/user/:id", func(c *gin.Context) 
func (c *Context) Param(key string) string 
func (c *Context) Query(key string) string 
func (c *Context) DefaultQuery(key, defaultValue string) string 
func (c *Context) GetQuery(key string) (string, bool) 
func (c *Context) QueryArray(key string) []string 
func (c *Context) GetQueryArray(key string) ([]string, bool) 
func (c *Context) QueryMap(key string) map[string]string 
func (c *Context) GetQueryMap(key string) (map[string]string, bool) 
func (c *Context) PostForm(key string) string 
func (c *Context) DefaultPostForm(key, defaultValue string) string 
func (c *Context) GetPostForm(key string) (string, bool) 
func (c *Context) PostFormArray(key string) []string 
func (c *Context) GetPostFormArray(key string) ([]string, bool) 
func (c *Context) PostFormMap(key string) map[string]string 
func (c *Context) GetPostFormMap(key string) (map[string]string, bool) 
func (c *Context) get(m map[string][]string, key string) (map[string]string, bool) 
func (c *Context) FormFile(name string) (*multipart.FileHeader, error) 
func (c *Context) MultipartForm() (*multipart.Form, error) 
func (c *Context) SaveUploadedFile(file *multipart.FileHeader, dst string) error 

func (c *Context) Bind(obj interface{}) error 
func (c *Context) BindJSON(obj interface{}) error 
func (c *Context) BindXML(obj interface{}) error 
func (c *Context) BindQuery(obj interface{}) error 
func (c *Context) BindYAML(obj interface{}) error 
func (c *Context) BindUri(obj interface{}) error 
func (c *Context) MustBindWith(obj interface{}, b binding.Binding) error 

func (c *Context) ShouldBind(obj interface{}) error 
func (c *Context) ShouldBindJSON(obj interface{}) error 
func (c *Context) ShouldBindXML(obj interface{}) error 
func (c *Context) ShouldBindQuery(obj interface{}) error 
func (c *Context) ShouldBindYAML(obj interface{}) error 
func (c *Context) ShouldBindUri(obj interface{}) error 
func (c *Context) ShouldBindWith(obj interface{}, b binding.Binding) error 
func (c *Context) ShouldBindBodyWith(obj interface{}, bb binding.BindingBody) (err error) 


func (c *Context) ClientIP() string 
func (c *Context) ContentType() string 
func (c *Context) IsWebsocket() bool 
func (c *Context) requestHeader(key string) string 
// bodyAllowedForStatus is a copy of http.bodyAllowedForStatus non-exported function.
func bodyAllowedForStatus(status int) bool 
func (c *Context) Status(code int) 

func (c *Context) Header(key, value string) 
func (c *Context) GetHeader(key string) string 

func (c *Context) GetRawData() ([]byte, error) 
func (c *Context) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool) 
func (c *Context) Cookie(name string) (string, error) 

// response
func (c *Context) Render(code int, r render.Render) 
func (c *Context) HTML(code int, name string, obj interface{}) 
func (c *Context) IndentedJSON(code int, obj interface{}) 
func (c *Context) SecureJSON(code int, obj interface{}) 
func (c *Context) JSONP(code int, obj interface{}) 
func (c *Context) JSON(code int, obj interface{}) 
func (c *Context) AsciiJSON(code int, obj interface{}) 
func (c *Context) PureJSON(code int, obj interface{}) 
func (c *Context) XML(code int, obj interface{}) 
func (c *Context) YAML(code int, obj interface{}) 
func (c *Context) ProtoBuf(code int, obj interface{}) 
func (c *Context) String(code int, format string, values ...interface{}) 
func (c *Context) Redirect(code int, location string) 

func (c *Context) Data(code int, contentType string, data []byte) 
func (c *Context) DataFromReader(code int, contentLength int64, contentType string, reader io.Reader, extraHeaders map[string]string) 
func (c *Context) File(filepath string) 
func (c *Context) FileAttachment(filepath, filename string) 
func (c *Context) SSEvent(name string, message interface{}) 
func (c *Context) Stream(step func(w io.Writer) bool) bool 
func (c *Context) Negotiate(code int, config Negotiate) 
func (c *Context) NegotiateFormat(offered ...string) string 
func (c *Context) SetAccepted(formats ...string) 
func (c *Context) Deadline() (deadline time.Time, ok bool) 
func (c *Context) Done() <-chan struct{} 
func (c *Context) Err() error 
func (c *Context) Value(key interface{}) interface{} 

```
