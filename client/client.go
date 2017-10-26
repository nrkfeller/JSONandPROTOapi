package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

// Rfq is what is send to server
type Rfq struct {
	ID              string `json:"id"`
	AccountID       string `json:"accountid"`
	ProductNumber   string `json:"productnumber"`
	ProductCategory string `json:"productcategory"`
	Quantity        string `json:"quantity"`
}

// JSONResponse of server
type JSONResponse struct {
	Message               string `json:"message"`
	UnitPrice             string `json:"unitprice"`
	PriceValidationPeriod string `json:"pricevalidationperiod"`
}

func jsonTemplate(w http.ResponseWriter, r *http.Request) {

	quoteJSONResponse := RequestJSONQuote(Rfq{
		ID:              r.FormValue("id"),
		AccountID:       r.FormValue("accountid"),
		ProductNumber:   r.FormValue("productnumber"),
		ProductCategory: r.FormValue("productcategory"),
		Quantity:        r.FormValue("quantity"),
	})

	err := tpl.ExecuteTemplate(w, "index.gohtml", quoteJSONResponse)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}

// RequestJSONQuote sends the request
func RequestJSONQuote(rfq Rfq) JSONResponse {

	jsonstring, _ := json.Marshal(rfq)

	req, _ := http.NewRequest(
		"GET",
		"http://localhost:8000/json/quote/"+rfq.ProductNumber+"?category="+rfq.ProductCategory,
		bytes.NewBuffer(jsonstring),
	)
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, _ := httpClient.Do(req)

	defer resp.Body.Close()

	var quoteJSONResponse JSONResponse

	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &quoteJSONResponse)

	return quoteJSONResponse
}

func protoTemplate(w http.ResponseWriter, r *http.Request) {

	quoteProtoResponse := ReqProtoQuote(RequestProtoQuote{
		ID:              r.FormValue("id"),
		AccountID:       r.FormValue("accountid"),
		ProductNumber:   r.FormValue("productnumber"),
		ProductCategory: r.FormValue("productcategory"),
		Quantity:        r.FormValue("quantity"),
	})

	err := tpl.ExecuteTemplate(w, "index.gohtml", quoteProtoResponse)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}

// ReqProtoQuote sends the request
func ReqProtoQuote(rfq RequestProtoQuote) ProtoResponse {

	_, err := proto.Marshal(&rfq)
	if err != nil {
		fmt.Println(err)
	}

	resp, _ := http.Get(
		"http://localhost:8000/proto/quote/" + rfq.ProductNumber + "?category=" + rfq.ProductCategory,
	)

	body, _ := ioutil.ReadAll(resp.Body)

	var quoteProtoResponse ProtoResponse

	_ = proto.Unmarshal(body, &quoteProtoResponse)

	return quoteProtoResponse
}

func main() {
	http.HandleFunc("/json/quote", jsonTemplate)
	http.HandleFunc("/proto/quote", protoTemplate)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
