package main

import (
	"net/http"
	"testing"
)

type TestResponseWriter struct {
	headers http.Header
	body    []byte
}

func (t TestResponseWriter) Header() (h http.Header) {
	return t.headers
}

func (t TestResponseWriter) Write(contents []byte) (int, error) {
	t.body = append(t.body, contents...)

	return len(contents), nil
}

func (t TestResponseWriter) WriteHeader(int) {
	return
}

func TestLoadCSV(t *testing.T) {
	businessDB, err := loadCSV()
	if err != nil {
		t.Error(err)
	}
	if businessDB == nil {
		t.Error("nil db")
	}
}

func TestBasicAuthSuccess(t *testing.T) {
	canaryHandler := func(http.ResponseWriter, *http.Request) {
	}

	request, err := http.NewRequest("GET", "foo.com/businesses", TestReader{})
	if err != nil {
		t.Errorf("error was not nil: %#v", err)
	}
	request.SetBasicAuth("", "token")

	MyMiddleware(TestResponseWriter{headers: http.Header{}}, request, canaryHandler)
}

func TestBasicAuthFail(t *testing.T) {
	canaryHandler := func(http.ResponseWriter, *http.Request) {
		t.Error("Shouldn't be called")
	}

	request, err := http.NewRequest("GET", "foo.com/businesses", TestReader{})
	if err != nil {
		t.Errorf("error was not nil: %#v", err)
	}
	request.SetBasicAuth("", "fake_token")

	MyMiddleware(TestResponseWriter{headers: http.Header{}}, request, canaryHandler)
}
