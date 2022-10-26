package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Yashsharma1911/mongoapi/router"
)

func main() {
	fmt.Println("Server is getting started! 🚀")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Server is started 🎉")
}
