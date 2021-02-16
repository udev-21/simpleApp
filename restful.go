package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	validator "github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type RestFullApi interface {
	apiGetAll(http.ResponseWriter, *http.Request)
	apiGetOne(http.ResponseWriter, *http.Request)
	apiCreate(http.ResponseWriter, *http.Request)
	apiUpdate(http.ResponseWriter, *http.Request)
	apiDelete(http.ResponseWriter, *http.Request)
}

func getAll(r RestFullApi) func(http.ResponseWriter, *http.Request) {
	return r.apiGetAll
}

func getOne(r RestFullApi) func(http.ResponseWriter, *http.Request) {
	return r.apiGetOne
}

func create(r RestFullApi) func(http.ResponseWriter, *http.Request) {
	return r.apiCreate
}

func update(r RestFullApi) func(http.ResponseWriter, *http.Request) {
	return r.apiUpdate
}

func delete(r RestFullApi) func(http.ResponseWriter, *http.Request) {
	return r.apiDelete
}

func (book Book) apiGetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db.Where("1=1").Model(&Book{})

	fmt.Println("Book apiGetAll()")

	tmpLimit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	tmpOffset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	if tmpLimit <= 0 {
		tmpLimit = 10
	}

	if tmpLimit > 100 {
		tmpLimit = 100
	}

	if tmpOffset <= 0 {
		tmpOffset = 0
	}

	result := Response{}
	books := []Book{}
	db.Limit(tmpLimit).Offset(tmpOffset).Find(&books)

	db.Statement.SQL.Reset()
	result.Result = books
	result.Ok = true
	checkErr(json.NewEncoder(w).Encode(result))
}

func (book Book) apiGetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var tmpBook = Book{}
	fmt.Println("Book apiGetOne()")

	db.Model(Book{}).Where("id = ?", params["id"]).First(&tmpBook)
	if tmpBook.ID <= 0 {

	}
	var Response Response
	Response.Ok = true
	Response.Result = tmpBook
	json.NewEncoder(w).Encode(Response)
}

func (book Book) apiCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Book apiCreate()")
	tmpBook := Book{}
	json.NewDecoder(r.Body).Decode(&tmpBook)

	var validate = validator.New()
	var errors []Error
	err := validate.Struct(tmpBook)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, Error{Field: err.StructField(), Value: err.Value()})
		}
		json.NewEncoder(w).Encode(Response{Errors: errors, Ok: false})
		return
	}

	db.Model(&Book{}).Create(&tmpBook)
	json.NewEncoder(w).Encode(tmpBook)
}

func (book Book) apiUpdate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Book apiUpdate()")
}

func (book Book) apiDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Book apiDelete()")
}
