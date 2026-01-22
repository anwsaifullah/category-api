package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server is running..")
	http.ListenAndServe(":8080", nil)
}
