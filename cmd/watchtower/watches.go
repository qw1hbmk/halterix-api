package watchtower

type Watch struct {
	Id          string `json:"id"`
	RecordingId string `json:"recordingId"`
	Network     string `json:"network"`
	ClientTime  int64  `json:"clientTime"`
	ServerTime  int64  `json:"serverTime"`
	PatientId   string `json:"patientId"`
	Active      bool   `json:"active"`
}
