package test

import (
	"jin"
	"testing"
)

func TestRouterGroup(t *testing.T) {
	r := jin.New()
	/*
		r.GET("/", func(c *jin.Context) {
			c.HTML(http.StatusOK, "<h1>Index Page</h1>")
		})
		v1 := r.Group("/v1")
		{
			v1.GET("/", func(c *jin.Context) {
				c.HTML(http.StatusOK, "<h1>Welcomt to RouterGroup V1</h1>")
			})
		}
		v2 := r.Group("/v2")
		{
			v2.GET("/", func(c *jin.Context) {
				c.HTML(http.StatusOK, "<h1>Welcomt to RouterGroup V2</h1>")
			})
		}
	*/

	r.Run(":8080")
}
