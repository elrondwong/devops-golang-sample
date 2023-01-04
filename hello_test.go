package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorldHandle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()

	router.GET("/", HelloWorldHandle)

	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, 200)

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	assert.Equal(t, string(body), "Hello World")
}
