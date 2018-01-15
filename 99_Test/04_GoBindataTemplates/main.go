package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	template "github.com/arschles/go-bindata-html-template"
	"github.com/gorilla/mux"
)

type someDTO struct {
	ExceptionText string      `json:"exceptionText"`
	Data          interface{} `json:"data"`
}

var layoutDir = "frontend/layout"
var testtemplate *template.Template
var err error

func main() {

	files := append(layoutFiles(), "frontend/index.gohtml")
	testtemplate, err = template.New("index", Asset).ParseFiles(files...)
	if err != nil {
		fmt.Printf("error parsing template: %s", err)
	}

	fmt.Println(testtemplate)

	r := mux.NewRouter()
	r.Handle("/test", http.HandlerFunc(testHandler))
	http.ListenAndServe(":3000", r)
}

func testHandler(w http.ResponseWriter, r *http.Request) {

	testtemplate.ExecuteTemplate(w, "layouttemplate", nil)
}

//-------------------------------------
//HELPERS
//-------------------------------------

func layoutFiles() []string {
	files, err := filepath.Glob(layoutDir + "/*.gohtml")
	if err != nil {
		//low-level exception logging
		fmt.Println(err.Error())
	}
	return files
}
