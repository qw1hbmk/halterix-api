package watchtower

import (
	"time"
)

type Watch struct {
	Id	 string `json:"id"`
	RecordId   string    `json:"record_id"`
	Active       bool    `json:"active"`
	Network		string `json:"network"`
	LastUpdate    time.Time    `json:"last_update"` 
}