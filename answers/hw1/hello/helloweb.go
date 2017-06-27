package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/gin-contrib/multitemplate"
	"net/http"
	"log"
	"os"
	"io"
	"fmt"
	"hello/app"
	"time"
)

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Passwd   string `form:"passwd" json:"passwd" binding:"required"`
	Age      int    `form:"age" json:"age"`
}

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// 設置example變數到Context的Key中, 通過MustGet等函數可以取得
		c.Set("Example", "12345")
		// 發送req之前
		c.Next() // 取出所有註冊的handler function都執行過一次, 再回到此func
		// 發送req之後
		latency := time.Since(t)
		log.Println(latency)
		// c.Writer是ResponseWriter, 可以透過她取得response status
		status := c.Writer.Status()
		log.Println(status)
	}
}

func createMyRender() multitemplate.Render {
	render := multitemplate.New()
	render.AddFromFiles("base", "hello/template/base.html", "hello/template/content.html", "hello/template/article.html")

	return render
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("hello/template/*")
	router.Use(logger())
	//router.HTMLRender = createMyRender()

	// 分組
	v1 := router.Group("/v1")
	v1.GET("/", func(c *gin.Context) {
		example := c.MustGet("Example")
		log.Println(example)

		//c.HTML(http.StatusOK, "base", gin.H{
		//	"title" : "Main webSite",
		//})
		c.String(http.StatusOK, app.HelloMessage())
	})

	v1.GET("/hw1", func(c *gin.Context){
		c.HTML(http.StatusOK, "hw1.html", gin.H{})
	})

	router.StaticFS("/resources", gin.Dir("hello/resources", false))

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		nation := c.DefaultQuery("nation", "taiwan")
		c.String(http.StatusOK, "HelloWorld! %s, %s", name, nation)
	})

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status":  gin.H{
				"status_code": http.StatusOK,
				"status":      "ok",
			},
			"message": message,
			"nick":    nick,
		})
	})

	router.POST("/upload", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("upload")
		if err != nil {
			c.String(http.StatusBadRequest, "Bad request")
			return
		}
		filename := header.Filename
		fmt.Println(file, err, filename)

		out, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}

		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}
		c.String(http.StatusCreated, "upload successful")
	})

	router.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", gin.H{
			"title":"測試加載HTML template!",
		})
	})

	router.GET("/login", func(c *gin.Context) {

	})

	// context.Bind可以把收到的content (json或x-www-form-urlencoded)透過struct做綁定
	router.POST("/login", func(c *gin.Context) {
		var user User

		err := c.Bind(&user)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"username":   user.Username,
			"passwd":     user.Passwd,
			"age":        user.Age,
		})

	})

	router.GET("/render", func(c *gin.Context) {
		contentType := c.DefaultQuery("ctype", "json")
		if contentType == "json" {
			c.JSON(http.StatusOK, gin.H{
				"user" : "chiakie",
				"password" : "1234",
			})
		} else if contentType == "xml" {
			c.XML(http.StatusOK, gin.H{
				"user" : "chiakie",
				"password" : "1234",
			})
		}
	})

	router.GET("/redirect/google", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://google.com")
	})

	router.Run(":8080")
}