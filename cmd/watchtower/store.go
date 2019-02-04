package watchtower

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
)

type database struct {
	store *firestore.Client
	ctx   context.Context
}

func NewDatabase(store *firestore.Client, ctx context.Context) *database {
	return &database{store, ctx}
}

func (db *database) Update(w Watch) (Watch, error) {

	w.Updated = time.Now().Unix()
	_, err := db.store.Collection("watches").Doc(w.Id).Set(db.ctx, map[string]interface{}{
		"id":         w.Id,
		"recordId":   w.RecordingId,
		"network":    w.Network,
		"lastUpdate": w.Updated,
	})
	if err != nil {
		log.Fatalf("Failed adding watch: %v", err)
		return Watch{}, err
	}
	return w, nil
}
