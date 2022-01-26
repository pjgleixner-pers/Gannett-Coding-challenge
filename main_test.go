package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gannett.com/api.grocery/model"
	"gannett.com/api.grocery/views"
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

func TestCreateItem(t *testing.T) {
	var itemPayload = []byte(`{"ID":"QWER-1234-TYUI-5678","Name":"Apple","Price":"1.99"}`)

	req, err := http.NewRequest("POST", "/items", bytes.NewBuffer(itemPayload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(model.PostItems)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	var m []views.Item
	json.Unmarshal(rr.Body.Bytes(), &m)

	if m[0].ID != "QWER-1234-TYUI-5678" {
		t.Errorf("Expected ProduceCode to be 'QWER-1234-TYUI-5678'. Got '%v'", m[0].ID)
	}
	if m[0].Name != "Apple" {
		t.Errorf("Expected Name to be 'Cucumber'. Got '%v'", m[0].Name)
	}
	if m[0].Price != "1.99" {
		t.Errorf("Expected Price to be '1.99'. Got '%v'", m[0].Price)
	}
}

func TestCreateItems(t *testing.T) {
	var itemPayload = []byte(`[{"ID":"QWER-1234-TYUI-5678","Name":"Apple","Price":"1.99"},{"ID":"ZXCV-5678-VBNM-9012","Name":"Tomato","Price":"0.99"}]`)

	req, err := http.NewRequest("POST", "/items", bytes.NewBuffer(itemPayload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(model.PostItems)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	var m []views.Item
	json.Unmarshal(rr.Body.Bytes(), &m)

	if m[0].ID != "QWER-1234-TYUI-5678" {
		t.Errorf("Expected ProduceCode to be 'QWER-1234-TYUI-5678'. Got '%v'", m[0].ID)
	}
	if m[0].Name != "Apple" {
		t.Errorf("Expected Name to be 'Cucumber'. Got '%v'", m[0].Name)
	}
	if m[0].Price != "1.99" {
		t.Errorf("Expected Price to be '1.99'. Got '%v'", m[0].Price)
	}

	if m[1].ID != "ZXCV-5678-VBNM-9012" {
		t.Errorf("Expected ProduceCode to be 'ZXCV-5678-VBNM-9012'. Got '%v'", m[1].ID)
	}
	if m[1].Name != "Tomato" {
		t.Errorf("Expected Name to be 'Cucumber'. Got '%v'", m[1].Name)
	}
	if m[1].Price != "0.99" {
		t.Errorf("Expected Price to be '1.99'. Got '%v'", m[1].Price)
	}
}
