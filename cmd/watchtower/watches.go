package watchtower

type Watch struct {
	Id        string `json:"id"`
	PatientId string `json:"patientId"`
	Active    bool   `json:"active"`
}
