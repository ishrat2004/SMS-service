package main

import (
	"fmt"
	"goserver/router"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting a new go server")

	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("listening on port 4000")
}
