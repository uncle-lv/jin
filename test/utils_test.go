package test

import (
	"jin"
	"net/http"
	"testing"
)

func TestResolveAddr(t *testing.T) {
	r := jin.Default()

	r.Run()
}

func TestAssertPath(t *testing.T) {
	r := jin.Default()

	r.GET("/", nil)

	r.GET("index", func(c *jin.Context) {
		c.String(http.StatusOK, "Welcome to Index Page")
	})

	r.Run()
}
