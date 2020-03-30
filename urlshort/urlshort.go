package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func main() {
	var jsonFile, yamlFile string
	flag.StringVar(&jsonFile, "json", "", "path to json file for url map.")
	flag.StringVar(&yamlFile, "yaml", "", "path to yaml file for url map.")

	flag.Parse()

	mux := defaultMux()

	var handler http.Handler
	if jsonFile != "" {
		jsonData, err := ioutil.ReadFile(jsonFile)
		if err != nil {
			log.Fatalf("failed to read JSON file, error: %v", err)
		}

		handler, err = JSONHandler(jsonData, mux)
		if err != nil {
			log.Fatalf("failed to create handler from JSON file, error: %v", err)
		}
	} else if yamlFile != "" {
		yamlData, err := ioutil.ReadFile(yamlFile)
		if err != nil {
			log.Fatalf("failed to read YAML file, error: %v", err)
		}

		handler, err = YAMLHandler(yamlData, mux)
		if err != nil {
			log.Printf("failed to create handler from YAML file, error :%v", err)
		}
	} else {
		log.Printf("no JSON/YAML backend file given, fallback to default handler /hello.\n")
		handler = mux
	}

	fmt.Println("Starting the url shortener server on :8080")
	http.ListenAndServe(":8080", handler)
}
