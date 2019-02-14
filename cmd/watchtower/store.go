package watchtower

import (
	"context"
	"errors"
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
	// Get Unix time in millis
	w.ServerTime = time.Now().UnixNano() / 1000000

	// Check if document exists before trying to update
	// TO DO: In future, api may allow this to be simplified to one call

	_, err := db.store.Collection("watches").Doc(w.Id).Get(db.ctx)
	if err != nil {
		return w, errors.New("Watch does not exist in databse.")
	}

	_, err = db.store.Collection("watches").Doc(w.Id).Set(db.ctx, map[string]interface{}{
		"recordingId": w.RecordingId,
		"network":     w.Network,
		"clientTime":  w.ClientTime,
		"serverTime":  w.ServerTime,
	})
	if err != nil {
		log.Fatalf("Failed adding watch: %v", err)
		return Watch{}, err
	}
	return w, nil
}

func (db *database) Get(watchId string) (Watch, error) {

	dsnap, err := db.store.Collection("watches").Doc(watchId).Get(db.ctx)
	if err != nil {
		return Watch{}, err
	}
	var w Watch
	dsnap.DataTo(&w)
	w.Id = watchId
	return w, nil
}
