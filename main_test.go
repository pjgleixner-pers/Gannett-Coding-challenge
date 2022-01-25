package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gannett.com/api.grocery/model"
)

func TestShowItems(t *testing.T) {
	req, err := http.NewRequest("GET", "/item", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(model.GetItems)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestShowItemById(t *testing.T) {
	req, err := http.NewRequest("GET", "/item/123-ABC", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(model.GetItemByID)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestDeleteItem(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/items/123-ABC", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(model.DeleteItems)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
