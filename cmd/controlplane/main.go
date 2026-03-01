package main

import (
	"fmt"
	"log"
	"net/http"
	)
func handleRoot(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status": "MeshGuard ControlPlane is Online", "version": "0.1.0"}`)
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", handleRoot)


	port := ":8080"
	fmt.Printf("Starting MeshGuard ControlPlane on Port %s \n", port)

	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
