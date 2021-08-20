package test

import (
	"jin"
	"net/http"
	"testing"
)

func TestTrieRouter(t *testing.T) {
	r := jin.New()

	r.GET("/hello/:name", func(c *jin.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *jin.Context) {
		c.JSON(http.StatusOK, jin.H{
			"filepath": c.Param("filepath"),
		})
	})

	r.Run(":8080")
}
