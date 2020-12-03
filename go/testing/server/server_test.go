package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoute(t *testing.T) {
	ts := httptest.NewServer(Route())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/hello?name=Bob")
	if err != nil {
		t.Fatalf("http.Get %s failed: %s", ts.URL, err)
	}
	hello, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatalf("read from HTTP Response Body failed: %s", err)
	}
	want := "Hi, Bob"
	if string(hello) != want {
		t.Fatalf("response of /hello?name=Bob returns %s want %s", string(hello), want)
	}
}
