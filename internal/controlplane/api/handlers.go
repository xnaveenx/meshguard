package api

import (
	"encoding/json"
	"net/http"
	"fmt"
	)
func HandleRegister(w http.ResponseWriter, r *http.Request){
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
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"status": "success", "message": "Device Registered Successfully"}`)
}
