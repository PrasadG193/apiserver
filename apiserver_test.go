package main

import (
	"fmt"
	"strings"
	"testing"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var loginUrl string

func init() {
	server := httptest.NewServer(SetHandlers())
	loginUrl = fmt.Sprintf("%s/login", server.URL)
	InitRedis()
	redisClient.FlushDB()
}

func TestClientCount(t *testing.T) {
	for i := 1; i < 10; i++ {
		resp, err := http.Get(loginUrl)
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != 200 {
			t.Fatalf("Response: %d\n", resp.StatusCode)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		expected := fmt.Sprintf("{\"Count\":%d}\n", i)
		if strings.Compare(string(body), expected) != 0 {
			t.Errorf("Invalid resp %s", string(body))
		}
	}
}
