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

func (db *database) PatchPatient(p Patient) (Patient, error) {

	// Check if document exists before trying to update it.
	// Todo: Fail and alert/alarm that we attempt to create a new watch.
	dsnap, err := db.store.Collection("patients").Doc(p.Id).Get(db.ctx)
	if err != nil {
		return p, err
	}

	var fp Patient
	dsnap.DataTo(&fp)

	// Get Unix time in millis
	fp.LastPingTime = time.Now().UnixNano() / 1000000
	// Only allow callers to update recordingId and network
	if p.RecordingId != nil {
		fp.RecordingId = p.RecordingId
	}
	if p.Network != nil {
		fp.Network = p.Network
	}

	_, err = db.store.Collection("patients").Doc(p.Id).Update(db.ctx, []firestore.Update{
		{Path: "recordingId", Value: fp.RecordingId},
		{Path: "network", Value: fp.Network},
		{Path: "lastPingTime", Value: fp.LastPingTime},
	})
	if err != nil {
		log.Fatalf("Failed updating watch: %v", err)
		return Patient{}, err
	}

	// Return original object
	return p, nil
}

func (db *database) GetWatch(watchId string) (Watch, error) {

	dsnap, err := db.store.Collection("watches").Doc(watchId).Get(db.ctx)
	if err != nil {
		return Watch{}, err
	}
	var w Watch
	err = dsnap.DataTo(&w)
	if err != nil {
		return Watch{}, err
	}
	w.Id = watchId
	return w, nil
}
