package watchtower

type Watch struct {
	Id           string  `json:"id"`
	RecordingId  *string `json:"recordingId"`
	Network      *string `json:"network"`
	LastPingTime int64   `json:"lastPingTime"`
	PatientId    string  `json:"patientId"`
	Active       bool    `json:"active"`
}
