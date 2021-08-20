package test

import (
	"jin"
	"net/http"
	"testing"
)

func TestRecovery(t *testing.T) {
	r := jin.Default()

	r.GET("/recovery", func(c *jin.Context) {
		name := []string{"Jack"}
		c.String(http.StatusOK, name[1])
	})

	r.Run(":8080")
}
