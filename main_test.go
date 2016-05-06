package main

import (
	"net/http"
	"testing"
)

type TestResponseWriter struct {
}

func (t TestResponseWriter) Header() (h http.Header) {
	return
}

func (t TestResponseWriter) Write([]byte) (int, error) {
	return 0, nil
}

func (t TestResponseWriter) WriteHeader(int) {
	return
}

func TestLoadCSV(t *testing.T) {
	businessDB, err := loadCSV()
	if err != nil {
		t.Errorf("error was not nil: %#v", err)
	}
	if businessDB == nil {
		t.Error("nil db")
	}
}

func TestBasicAuth(t *testing.T) {
	canaryHandler := func(http.ResponseWriter, *http.Request) {
		t.Error("Shouldn't be called")
	}

	request, err := http.NewRequest("GET", "foo.com/businesses", TestReader{})
	if err != nil {
		t.Errorf("error was not nil: %#v", err)
	}

	MyMiddleware(TestResponseWriter{}, request, canaryHandler)
}
