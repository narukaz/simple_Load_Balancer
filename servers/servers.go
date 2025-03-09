package servers

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/narukaz/simple_Load_Balancer/config"
)

func RaiseServers() {
	var wg sync.WaitGroup
	for _, value := range config.ServerSlice {
		wg.Add(1)
		go func(value *config.Server, wg *sync.WaitGroup) {
			mux := http.NewServeMux()
			fmt.Println("connecting on port", value.Port)
			mux.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("I am running perfectly fine"))
			})
			if err := http.ListenAndServe(value.Host+":"+value.Port, mux); err != nil {
				wg.Done()
			}

		}(&value, &wg)

	}
	wg.Wait()
}
