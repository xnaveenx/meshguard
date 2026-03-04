package api

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/xnaveenx/meshguard/internal/controlplane/database"
	)

type APIServer struct {
	DB *database.Database
}

func (s *APIServer) HandleRegister(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var payload DeviceRegistration
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "Invalid JSON format"}`)
		return
	}
	if payload.MachineKey == "" || payload.Hostname == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "No MachineKey/Hostname found"}`)
		return
	}

	dbErr := s.DB.SaveDevice(payload.MachineKey, payload.Hostname, payload.OS)
	if dbErr != nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error" : "Failed to save device to database"}`)
		fmt.Printf("Database error: %v\n", dbErr)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"status": "success", "message": "Device Registered Successfully"}`)
}
