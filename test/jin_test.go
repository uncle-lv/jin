package test

import (
	"fmt"
	"jin"
	"net/http"
	"testing"
)

func TestJin(t *testing.T) {
	r := jin.New()

	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world!")
	})

	r.POST("/post", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "POST request has handled.")
	})

	r.Run(":8080")
}
