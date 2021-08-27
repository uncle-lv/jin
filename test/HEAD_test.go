package test

import (
	"jin"
	"testing"
)

func TestHEAD(t *testing.T) {
	r := jin.New()

	r.GET("/", func(c *jin.Context) {
		c.JSON(200, jin.H{
			"msg": "Hello World!",
		})
	})

	r.HEAD("/", func(c *jin.Context) {
		c.JSON(200, jin.H{
			"msg": "Hello World!",
		})
	})

	r.Run(":8080")
}
