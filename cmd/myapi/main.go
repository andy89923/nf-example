package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"status":"ok"}`)
	})

	mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		defer r.Body.Close()
		var v any
		body, _ := io.ReadAll(r.Body)
		if err := json.Unmarshal(body, &v); err != nil {
			v = map[string]string{"echo": string(body)}
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(v)
	})

	addr := ":8080"
	log.Println("myapi listening on", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
