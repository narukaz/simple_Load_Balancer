package main

import (
	"fmt"
	"net/http"

	"github.com/narukaz/simple_Load_Balancer/servers"
)

func main() {

	//creating main server
	var mux = http.NewServeMux()
	// mux.HandleFunc("/",func)
	go servers.RaiseServers()
	fmt.Println("server live on port 4000")
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
	if err := http.ListenAndServe("localhost:4000", mux); err != nil {
		fmt.Println("server crashed")
	}

}
