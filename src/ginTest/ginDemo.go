package ginTest

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//GinTest 测试
func GinTest() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

	// r.POST("/xxxpost", getting)
	// r.PUT("/xxxput")

	//api参数
	//http://127.0.0.1:8080/apiParame/tinamei/video
	r.GET("/apiParame/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		action = strings.Trim(action, "/")
		log.Println("name=" + name)     //name=tinamei
		log.Println("action=" + action) //action=video

	})

	//URL参数
	// URL参数可以通过DefaultQuery()或Query()方法获取
	// DefaultQuery()若参数不村则，返回默认值，Query()若不存在，返回空串
	// API ? name=zs
	//http://127.0.0.1:8080/urlParame?name=sm
	r.GET("/urlParame", func(c *gin.Context) {
		name := c.DefaultQuery("name", "老王")
		age := c.Query("age")
		log.Println("name：" + name)
		log.Println("age:" + age)
	})

	// 表单传输为post请求，http常见的传输格式为四种：
	// application/json
	// application/x-www-form-urlencoded
	// application/xml
	// multipart/form-data
	// 表单参数可以通过PostForm()方法获取，该方法默认解析的是x-www-form-urlencoded或from-data格式的参数
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("userpassword")
		// c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
		c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	})

	//上传单个文件
	// multipart/form-data格式用于文件上传
	// gin文件上传与原生的net/http方法类似，不同在于gin把原生的request封装到c.Request中
	//限制上传最大尺寸
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(500, "上传图片出错")
		}
		// c.JSON(200, gin.H{"message": file.Header.Context})
		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, file.Filename)
	})

	//上传多个文件
	// 限制表单上传大小 8MB，默认为32MB
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload1", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		}
		// 获取所有图片
		files := form.File["files"]
		// 遍历所有图片
		for _, file := range files {
			// 逐个存
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
				return
			}
		}
		c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
	})

	// 路由组1 ，处理GET请求
	v1 := r.Group("/v1")
	// {} 是书写规范
	{
		v1.GET("/login", login)
		v1.GET("submit", submit)
	}
	v2 := r.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}


	// 服务端要给客户端cookie
	r.GET("cookie", func(c *gin.Context) {
		// 获取客户端是否携带cookie
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
		   cookie = "NotSet"
		   // 给客户端设置cookie
		   //  maxAge int, 单位为秒
		   // path,cookie所在目录
		   // domain string,域名
		   //   secure 是否智能通过https访问
		   // httpOnly bool  是否允许别人通过js获取自己的cookie
		   c.SetCookie("key_cookie", "value_cookie", 60, "/",
			  "localhost", false, true)
		}

		log.Fatal(cookie)
	})


	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8080")
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}
