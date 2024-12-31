package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ports := []string{"8080", "8081", "8082"}

	for i, port := range ports {
		wg.Add(1)
		go func(p string, instance int) {
			defer wg.Done()
			log.Printf("Starting server on port %s\n", p)
			if err := http.ListenAndServe(":"+p, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				response := fmt.Sprintf("Servidor %d na porta %s", instance+1, p)
				w.Write([]byte(response))
			})); err != nil {
				log.Fatalf("Could not start server on port %s: %v\n", p, err)
			}
		}(port, i)
	}

	wg.Wait()
}
