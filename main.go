package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/verify", handleRequest)

	httpPort := os.Getenv("PORT")

	if len(httpPort) == 0 {
		httpPort = "8080"
	}

	log.Println("Starting server at http://localhost:" + httpPort)

	err := http.ListenAndServe(":"+httpPort, logRequest(http.DefaultServeMux))

	if err != nil {
		panic(err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	email := r.URL.Query().Get("email")

	if email == "" {
		http.Error(w, "You need to pass and email address to verify.", 500)
		return
	}

	verify_result := VerifyResult{Email: email}

	verify_result.Verify()

	json.NewEncoder(w).Encode(verify_result)
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
