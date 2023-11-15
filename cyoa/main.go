package main

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"cyoa/parseStory"
)

const (
	rootPath   = "/"
	serverPort = ":8080"
	introPath  = "/intro"
)

var tmpl = template.Must(template.ParseFiles("layout.html"))
var fileData, err = parseStory.JSON("data.json")

func main() {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	registerHandlers()

	fmt.Printf("Starting the server on %s\n", serverPort)
	http.ListenAndServe(serverPort, nil)

}

func registerHandlers() {
	http.HandleFunc(rootPath, handler)

	for key := range fileData {
		path := "/" + key
		http.HandleFunc(path, handler)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	currentPath := r.URL.Path
	chapterTitle := strings.TrimSpace(currentPath[1:])

	if currentPath == rootPath {
		http.Redirect(w, r, introPath, http.StatusFound)
	}

	_, exists := fileData[chapterTitle]
	if !exists {
		http.Error(w, "Chapter not found.", http.StatusNotFound)
		return
	}

	layoutData := parseStory.Chapter{
		Title:      fileData[chapterTitle].Title,
		Paragraphs: fileData[chapterTitle].Paragraphs,
		Options:    fileData[chapterTitle].Options,
	}

	tmpl.Execute(w, layoutData)
}
