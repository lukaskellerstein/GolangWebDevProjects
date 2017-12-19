package main

import (
	"html/template"
	"net/http"
)

//***********************************************
// TEMPLATE CACHING
//***********************************************
var templates = template.Must(template.ParseFiles("index.gohtml", "contact.gohtml", "catalog.gohtml", "aboutus.gohtml"))

//***********************************************
//***********************************************

type IndexViewModel struct {
	Title string
	Text1 []byte
	Text2 []byte
	Text3 []byte
}

//USE AS > http://localhost:8080/
func indexHandler(w http.ResponseWriter, r *http.Request) {
	//View model data
	vm := &IndexViewModel{
		Title: "Index",
		Text1: []byte("AAAAA"),
		Text2: []byte("BBBBBBB"),
		Text3: []byte("CCCCC"),
	}
	//Render template
	err := templates.ExecuteTemplate(w, "index.gohtml", vm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type Product struct {
	Title string
	Done  bool
}
type CatalogViewModel struct {
	Title    string
	Products []Product
}

//USE AS > http://localhost:8080/catalog
func catalogHandler(w http.ResponseWriter, r *http.Request) {
	//View model data
	vm := &CatalogViewModel{
		Title: "Catalog",
		Products: []Product{
			{Title: "Product 1", Done: false},
			{Title: "Product 2", Done: true},
			{Title: "Product 3", Done: true},
		},
	}
	//Render template
	err := templates.ExecuteTemplate(w, "catalog.gohtml", vm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//USE AS > http://localhost:8080/contact
func contactHandler(w http.ResponseWriter, r *http.Request) {
	//Render template
	err := templates.ExecuteTemplate(w, "contact.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//USE AS > http://localhost:8080/aboutus
func aboutusHandler(w http.ResponseWriter, r *http.Request) {
	//Render template
	err := templates.ExecuteTemplate(w, "aboutus.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/catalog/", catalogHandler)
	http.HandleFunc("/contact/", contactHandler)
	http.HandleFunc("/aboutus/", aboutusHandler)
	http.ListenAndServe(":8080", nil)
}
