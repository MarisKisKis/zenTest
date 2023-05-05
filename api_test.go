package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostIncr(t *testing.T) {
	var jsonStr = []byte(`{"key": "age","value": 19}`)

	req, err := http.NewRequest("POST", "localhost:8080/redis/incr", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestPostSha512(t *testing.T) {
	var jsonStr = []byte(`{"text": "test","key": "test123"}`)

	req, err := http.NewRequest("POST", "localhost:8080/sign/hmacsha512", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestPostUser(t *testing.T) {
	var jsonStr = []byte(`{"name":"Alex","age":20}`)

	req, err := http.NewRequest("POST", "localhost:8080/postgres/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
