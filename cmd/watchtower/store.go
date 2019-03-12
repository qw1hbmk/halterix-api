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

func (db *database) Patch(w Watch) (Watch, error) {

	// Check if document exists before trying to update it.
	// Todo: Fail and alert/alarm that we attempt to create a new watch.
	dsnap, err := db.store.Collection("watches").Doc(w.Id).Get(db.ctx)
	if err != nil {
		return w, errors.New("Watch does not exist in database.")
	}

	var fw Watch
	dsnap.DataTo(&fw)
	if fw.Active == false {
		return fw, nil
	}
	// Get Unix time in millis
	fw.LastPingTime = time.Now().UnixNano() / 1000000
	// Only allow callers to update recordingId and network
	if w.RecordingId != nil {
		fw.RecordingId = w.RecordingId
	}
	if w.Network != nil {
		fw.Network = w.Network
	}

	_, err = db.store.Collection("watches").Doc(w.Id).Update(db.ctx, []firestore.Update{
		{Path: "recordingId", Value: fw.RecordingId},
		{Path: "network", Value: fw.Network},
		{Path: "lastPing", Value: fw.LastPingTime},
	})
	if err != nil {
		log.Fatalf("Failed updating watch: %v", err)
		return Watch{}, err
	}
	// Todo: Should use response from firestore
	return fw, nil
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
