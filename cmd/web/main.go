package main

import (
	"fmt"
	"net/http"

	"github.com/KanurkarPrateek/Golang_web_app/pkg/handlers"
)

const PortNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("starting app on port %s", PortNumber))
	_ = http.ListenAndServe(PortNumber, nil)
}
