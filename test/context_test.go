package test

import (
	"jin"
	"net/http"
	"testing"
)

func TestContext(t *testing.T) {
	r := jin.New()

	/*
		r.GET("/", func(c *jin.Context) {
			c.HTML(http.StatusOK, "<h1>Welcome!</h1>")
		})
	*/

	r.GET("/hello", func(c *jin.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *jin.Context) {
		c.JSON(http.StatusOK, jin.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":8080")
}
