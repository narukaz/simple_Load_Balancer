package servers

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/narukaz/simple_Load_Balancer/config"
	"github.com/narukaz/simple_Load_Balancer/operations"
)

func RaiseServers() {
	var wg sync.WaitGroup
	var CustomClient operations.Mongo
	temp, err := operations.ConnectToMongo("mongodb://localhost:27017/")

	if err != nil {
		log.Fatal("failed to connect to mongo client")
	}
	CustomClient.Client = temp

	for _, value := range config.ServerSlice {
		wg.Add(1)
		go func(value *config.Server, wg *sync.WaitGroup) {
			mux := http.NewServeMux()
			fmt.Println("connecting on port", value.Port)
			mux.HandleFunc("/get", CustomClient.Get)
			mux.HandleFunc("/delete", CustomClient.Delete)
			mux.HandleFunc("/add", CustomClient.Add)
			if err := http.ListenAndServe(value.Host+":"+value.Port, mux); err != nil {
				wg.Done()
			}

		}(&value, &wg)

	}
	wg.Wait()
}
