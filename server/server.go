package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
)

var products = make(map[string]Product)

// Product single product
type Product struct {
	Number           string `json:"number"`
	Category         string `json:"category"`
	Price            string `json:"price"`
	ValidationPeriod string `json:"validationperiod"`
}

// Products dictionary struct
type Products struct {
	AllProducts map[string]Product `json:"allprods"`
}

// Rfq requets struct
type Rfq struct {
	ID              string `json:"id"`
	AccountID       string `json:"accountid"`
	ProductNumber   string `json:"productnumber"`
	ProductCategory string `json:"productcategory"`
	Quantity        string `json:"quantity"`
}

// Reply json struct
type Reply struct {
	Message               string `json:"message"`
	UnitPrice             string `json:"unitprice"`
	PriceValidationPeriod string `json:"pricevalidationperiod"`
}

// GetQuote gets the quote
func GetQuote(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)

	u, _ := url.Parse(fmt.Sprint(r.URL))
	m, _ := url.ParseQuery(u.RawQuery)

	params := mux.Vars(r)

	rep := Reply{}

	val, ok1 := products[params["id"]]
	ok2 := val.Category == m["category"][0]

	if ok1 && ok2 {
		rep.Message = "Found"
		rep.UnitPrice = val.Price
		rep.PriceValidationPeriod = val.ValidationPeriod

	} else {
		rep.Message = "Product not found"
	}

	json.NewEncoder(w).Encode(rep)
}

// GetProd get a single product
func GetProd(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Method, r.URL)
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(products[params["id"]])
}

// PostProd create new product entry
func PostProd(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)

	var prod Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&prod)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	products[prod.Number] = prod

	json.NewEncoder(w).Encode(products)
}

// GetProducts gets all products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)
	json.NewEncoder(w).Encode(Products{products})
}

// GetProtoQuote returns Reply type
func GetProtoQuote(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)

	u, _ := url.Parse(fmt.Sprint(r.URL))
	m, _ := url.ParseQuery(u.RawQuery)

	params := mux.Vars(r)

	rep := ProtoResponse{}

	val, ok1 := products[params["id"]]
	ok2 := val.Category == m["category"][0]

	if ok1 && ok2 {
		rep.Message = "Found"
		rep.UnitPrice = val.Price
		rep.PriceValidationPeriod = val.ValidationPeriod

	} else {
		rep.Message = "Product not found"
	}

	data, _ := proto.Marshal(&rep)

	w.Write(data)
}

func main() {
	products["1"] = Product{"1", "fruits", "1.99", "4 days"}
	products["2"] = Product{"2", "fruits", "5.99", "10 Days"}
	products["3"] = Product{"3", "electronics", "100.00", "2 months"}
	products["4"] = Product{"4", "electronics", "250.00", "3 months"}

	router := mux.NewRouter()
	router.HandleFunc("/json/products", GetProducts).Methods("GET")
	router.HandleFunc("/json/product/{id}", GetProd).Methods("GET")
	router.HandleFunc("/json/products", PostProd).Methods("POST")
	router.HandleFunc("/json/quote/{id}", GetQuote).Methods("GET")

	router.HandleFunc("/proto/products", GetProducts).Methods("GET")
	router.HandleFunc("/proto/product/{id}", GetProd).Methods("GET")
	router.HandleFunc("/proto/products", PostProd).Methods("POST")
	router.HandleFunc("/proto/quote/{id}", GetProtoQuote).Methods("GET")

	addr := ":8000"
	log.Println("listen on", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
