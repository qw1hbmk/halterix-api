package watchtower

type Watch struct {
	Id          string `json:"id"`
	RecordingId string `json:"recordingId"`
	Network     string `json:"network"`
	ClientTime  int64  `json:"clientTime"`
	ServerTime  int64  `json:"serverTime"`
	PatientId   string `json:"patientId"`
	active      bool   `json:"patientId"`

	// watchId should be unique (not in db on watch creation)
}

// on replacement of patient watch, set watch record to inactive before overwriting patient's watchId
