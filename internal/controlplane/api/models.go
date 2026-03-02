package api

type DeviceRegistration struct {
	MachineKey string `json:"machine_key"`
	Hostname   string `json:"hostname"`
	OS	   string `json:"os"`
}
