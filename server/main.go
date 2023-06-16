package main

import (
	"fmt"
	"log"
	"net/http"
	"todo/routes"
)

func main() {
	r := routes.Routes()
	fmt.Println("server is going to staert on 2020")
	log.Fatal(http.ListenAndServe(":2020",r))
}
