package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type billing struct {
	firstName string
	lastName  string
	address1  string
	address2  string
	city      string
	state     string
	postcode  string
	country   string
	email     string
	phone     string
}

type shipping struct {
	firstName string
	lastName  string
	address1  string
	address2  string
	city      string
	state     string
	postcode  string
	country   string
}

type order struct {
	paymentMethod      string
	paymentMethodTitle string
	createdVia         string
	setPaid            string
	status             string
	total              string
	orderType          string
	parentID           string
	billing            billing
	shipping           shipping
	lineItems          string
	shippingLines      string
}

func incomingOrder(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var o order
	err = json.Unmarshal(body, &o)
	if err != nil {
		panic(err)
	}
	log.Println(o)
}

func main() {

	http.HandleFunc("/", incomingOrder)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
