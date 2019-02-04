package watchtower

type Watch struct {
	Id          string `json:"id"`
	RecordingId string `json:"recording_id"`
	Network     string `json:"network"`
	Updated     int64  `json:"updated"`
}
