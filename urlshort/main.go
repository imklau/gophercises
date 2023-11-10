package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"urlshort/handler"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := handler.MapHandler(pathsToUrls, mux)

	yamlFile, err := ioutil.ReadFile("paths.yaml")

	if err != nil {
        log.Printf("Unable to open the yaml file:  #%v ", err)
    }

	yamlHandler, err := handler.YAMLHandler([]byte(yamlFile), mapHandler)
	
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}