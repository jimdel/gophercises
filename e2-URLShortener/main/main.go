package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"jamesdelay.com/urlshort"
)

func main() {
	mux := defaultMux()

	//Accept an optional CLI arg for YAML/JSON file
	yamlSource := flag.String("ymlSrc", "./urlpaths.yaml", "path to YAML source file containing url path mappings")
	jsonSource := flag.String("jsonSrc", "./urlpaths.json", "path to JSON source file containing url path mappings")

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the fallback
	yaml := readFile(*yamlSource)
	jsonFile := readFile(*jsonSource)

	yamlHandler, yamlErr := urlshort.YAMLHandler(yaml, mapHandler)
	jsonHandler, jsonErr := urlshort.JSONHandler(jsonFile, mapHandler)

	if yamlErr != nil {
		panic(yamlErr)
	}

	if jsonErr != nil {
		panic(jsonErr)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
	http.ListenAndServe(":8080", jsonHandler)

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/help", hello)
	return mux
}

func readFile(path string) []byte {
	contents, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return contents
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
