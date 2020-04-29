package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/zasdaym/gophercises/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the server")
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()

	f, err := os.Open(*filename)
	if err != nil {
		log.Fatalln(err)
	}

	story, err := cyoa.JSONStory(f)
	if err != nil {
		log.Println(err)
	}

	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
