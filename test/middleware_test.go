package test

import (
	"jin"
	"jin/log"
	"net/http"
	"testing"
)

func TestMiddleware(t *testing.T) {
	r := jin.New()
	r.Use(globalMiddleware())
	r.GET("/", func(c *jin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello</h1>")
	})

	v1 := r.Group("/v1")
	v1.Use(forV1())
	{
		v1.GET("/", func(c *jin.Context) {
			c.HTML(http.StatusOK, "<h1>Welcome to V1</h1>")
		})
	}

	r.Run(":8080")
}

func forV1() jin.HandlerFunc {
	return func(c *jin.Context) {
		log.Info("Before from forV1")
		c.Next()
		log.Info("After from forV1")
	}
}

func globalMiddleware() jin.HandlerFunc {
	return func(c *jin.Context) {
		log.Info("Before from globalMiddleware")
		c.Next()
		log.Info("After from globalMiddleware")
	}
}
