package test

import (
	"jin"
	"jin/log"
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

func TestSetCookie(t *testing.T) {
	r := jin.Default()

	r.GET("/cookie", func(c *jin.Context) {
		cookie, err := c.Cookie("jin_cookie")

		if err != nil {
			cookie = "NotSet"
			c.SetCookie("jin_cookie", "test", 3600, "/", "localhost", false, true)
		}

		log.Infof("Cookie value: %s\n", cookie)
	})

	r.Run(":8080")
}

func TestUploadSingleFile(t *testing.T) {
	const dst string = "../static/file/"

	r := jin.Default()
	r.POST("/upload", func(c *jin.Context) {
		file, _ := c.FormFile("file")
		log.Infof("file name: %s", file.Filename)
		c.SaveUploadFile(file, dst, file.Filename)
		c.String(http.StatusOK, "Upload succeed!")
	})

	r.Run(":8080")
}

func TestRedirect(t *testing.T) {
	r := jin.Default()

	r.GET("/", func(c *jin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/index")
	})

	r.GET("/index", func(c *jin.Context) {
		c.String(http.StatusOK, "Welcome to Index Page!")
	})

	r.Run(":8080")
}
