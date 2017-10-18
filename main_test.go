package main

import (
	"fmt"
	"testing"
	"net/http"
	"io/ioutil"

	"github.com/stretchr/testify/assert"
)

func TestCheckResponse(t *testing.T) {

	res, err := http.Get("http://localhost:8080/")
	if err != nil {
		fmt.Println(err)
	}

	resp, _ := ioutil.ReadAll(res.Body)
	assert.Equal(c.HTML(http.StatusOK, "hello.tmpl", map[string]string{"name": "world"}))
}