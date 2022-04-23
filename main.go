package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/devsquared/godevblog/templates"
)

func main() {
	fmt.Println("Welcome to godevblog. Let's get to writing!")

	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse() // parse the flags

	http.Handle("/", templates.NewTemplateHandler("tailwindHTMLBase.html"))

	// start the web server
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
