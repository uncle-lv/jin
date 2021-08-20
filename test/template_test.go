package test

import (
	"jin"
	"net/http"
	"testing"
)

func TestTemplate(t *testing.T) {
	r := jin.New()

	r.LoadHTMLGlob("../templates/*")
	r.Static("/assets", "../static")

	r.GET("/", func(c *jin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", jin.H{
			"title":   "Jin",
			"content": "Welcome",
		})
	})

	r.Run(":8080")
}
