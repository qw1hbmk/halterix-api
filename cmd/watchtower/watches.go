package watchtower

type Watch struct {
	Id           string  `json:"id,omitempty"`
	RecordingId  *string `json:"recordingId,omitempty"`
	Network      *string `json:"network,omitempty"`
	LastPingTime int64   `json:"lastPingTime,omitempty"`
	PatientId    string  `json:"patientId,omitempty"`
	Active       bool    `json:"active,omitempty"`
}
