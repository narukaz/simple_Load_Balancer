package main

import (
	"fmt"
	"net/http"
)

func main() {

	//creating main server
	var mux = http.NewServeMux()

	mux.HandleFunc("/")

	fmt.Println("server live on port 4000")
	if err := http.ListenAndServe("localhost:4000", mux); err != nil {
		fmt.Println("server crashed")
	}

}
