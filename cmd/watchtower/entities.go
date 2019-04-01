package watchtower

import "time"

type Patient struct {
	Id           string  `json:"id" firestore:"id,omitempty"`
	Active       bool    `json:"active" firestore:"active"`
	RecordingId  *string `json:"recordingId" firestore:"recordingId"`
	Network      *string `json:"network" firestore:"network"`
	LastPingTime int64   `json:"lastPingTime" firestore:"lastPingTime"`
	WatchId      string  `json:"watchId" firestore:"watchId"`
}

type Watch struct {
	Id        string `json:"id" firestore:":id,omitempty"`
	PatientId string `json:"patientId" firestore:"patientId"`
	Active    bool   `json:"active" firestore:"active"`
}

type WearLog struct {
	Id         string    `json:"id" firestore:"id,omitempty"`
	WatchId    string    `json:"watchId" firestore:"watchId"`
	PatientId  string    `json:"patientId" firestore:"patientId"`
	Message    string    `json:"message" firestore:"message"`
	Code       int       `json:"code" firestore:"code"`
	ClientTime time.Time `json:"clientTime" firestore:"clientTime"`
	ServerTime time.Time `json:"serverTime" firestore:"serverTime"`
}
