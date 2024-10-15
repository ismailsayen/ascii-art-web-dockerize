package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	asciiartweb "asciiartweb/handlers"
)

const port string = ":8080"

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Please enter only the program name.")
		return
	}
	
	http.HandleFunc("/assets/images/", asciiartweb.ImagesHandler)
	http.HandleFunc("/assets/css/", asciiartweb.CssHandler)
	http.HandleFunc("/", asciiartweb.IndexPage)
	http.HandleFunc("/ascii-art", asciiartweb.AsciiArtPage)
	fmt.Println("http://localhost:8080/")
	log.Fatal(http.ListenAndServe(port, nil))
}
