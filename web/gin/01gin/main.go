package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Name     string `json:"u"`
	Password string `json:"p"`
	Email    string `json:"email" binding:"required,email,googleEmail"`
}

var googleEmail validator.Func = func(fl validator.FieldLevel) bool {
	email, ok := fl.Field().Interface().(string)
	if ok {
		if !strings.Contains(email, "gmail") {
			return false
		}
	}
	return true
}

func runTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		// t := time.Now()
		// 请求前
		fmt.Println("before")
		c.Set("cost", "1111")
		c.Next()
		fmt.Println("after")
		// 请求后
		// latency := time.Since(t)

	}
}

func main() {
	router := gin.Default()

	// hello,go
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "hello,%s", "gin")
	})
	// group
	grp1 := router.Group("/testGroup1")
	grp1.GET("/grp1", func(ctx *gin.Context) {
		ctx.String(200, "grp1")
	})
	grp2 := router.Group("/testGroup2")
	grp2.GET("/grp2", func(ctx *gin.Context) {
		ctx.String(200, "grp2")
	})
	// restful
	router.POST("/update", func(ctx *gin.Context) {

	})
	router.PUT("/put", func(ctx *gin.Context) {

	})
	router.DELETE("/del", func(ctx *gin.Context) {

	})

	// 重定向
	// 重定向到外部
	router.GET("/testOut", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	})

	// 重定向到内部
	router.POST("/testIn", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/foo")
	})

	router.GET("/testOut2", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		router.HandleContext(c)
	})
	router.GET("/testJson", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	// 静态文件
	router.Static("/static", "./static")
	// router.StaticFS("/static", http.Dir("static"))
	router.StaticFile("/f1", "./static/1.txt")

	// html
	router.LoadHTMLGlob("template/**/*")
	router.GET("/html/testTemplate", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "template/user/user.tmpl", gin.H{
			"title": "user",
		})
	})
	// bindParam
	router.POST("/testParams/:u/:p", func(ctx *gin.Context) {
		u := ctx.Param("u")
		p := ctx.Param("p")
		ctx.JSON(http.StatusOK, &User{Name: u, Password: p})
	})

	router.POST("/testParams", func(ctx *gin.Context) {
		u := ctx.Query("u")
		p := ctx.Query("p")
		ctx.JSON(http.StatusOK, &User{Name: u, Password: p})
	})

	router.POST("/testParam", func(ctx *gin.Context) {
		u := ctx.PostForm("u")
		p := ctx.PostForm("p")
		ctx.JSON(http.StatusOK, &User{Name: u, Password: p})
	})

	router.POST("/testParamJSON", func(ctx *gin.Context) {
		var user User
		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, &User{Name: user.Name, Password: user.Password})
	})

	// valid
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("googleEmail", googleEmail)
	}
	router.POST("/testVaild", func(ctx *gin.Context) {
		var user User
		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, user)
	})

	// midware
	// router.Use(runTime())
	// grp1.Use(runTime())
	router.POST("/testMidware", runTime(), func(ctx *gin.Context) {
		var user User
		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		cost := ctx.MustGet("cost").(string)
		ctx.JSON(http.StatusOK, gin.H{"cost": cost})
	})
	// https
	// router.RunTLS("","","")

	err := router.Run()
	if err != nil {
		panic(err)
	}
}
