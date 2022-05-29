package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
)

func TestIndexPageHandler(t *testing.T) {
	mockResponse := `{"message":"This service lists all Nigerian companies, their logos, address, type and categories"}`
	r := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	fmt.Println(string(responseData))
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestAddNewCompany(t *testing.T) {
	r := setupRouter()
	// r.POST("/companies", nil)
	companyId := xid.New().String()
	company := Company{
		ID:        companyId,
		Name:      "Mekio-Tech",
		Sector:    "IT",
		Category:  "small",
		IsStartup: true,
		CEO:       "Bright Edwin",
		Revenue:   "",
	}

	jsonValue, _ := json.Marshal(company)
	req, _ := http.NewRequest("POST", "/companies", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

}

func TestGetCompanies(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest("GET", "/companies", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var companies []Company
	json.Unmarshal(w.Body.Bytes(), &companies)

	fmt.Println("Companies", companies)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, companies)
}
