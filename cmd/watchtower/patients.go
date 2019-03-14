package watchtower

type Patient struct {
	Id           string  `json:"id"`
	Active       bool    `json:"active"`
	RecordingId  *string `json:"recordingId"`
	Network      *string `json:"network"`
	LastPingTime int64   `json:"lastPingTime"`
	WatchId      string  `json:"watchId"`
}
