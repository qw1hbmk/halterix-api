package watchtower

import (
	"context"
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
	if p.Network != "" {
		fp.Network = p.Network
	}

	_, err = db.store.Collection("patients").Doc(p.Id).Update(db.ctx, []firestore.Update{
		{Path: "recordingId", Value: fp.RecordingId},
		{Path: "network", Value: fp.Network},
		{Path: "lastPingTime", Value: fp.LastPingTime},
	})
	if err != nil {
		return Patient{}, err
	}

	// Return firebase object
	fp.Id = p.Id
	return fp, nil
}

func (db *database) GetPatient(patientId string) (Patient, error) {

	dsnap, err := db.store.Collection("patients").Doc(patientId).Get(db.ctx)
	if err != nil {
		return Patient{}, err
	}
	var p Patient
	err = dsnap.DataTo(&p)
	if err != nil {
		return Patient{}, err
	}
	p.Id = patientId
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

func (db *database) PostWearLog(wl WearLog) (WearLog, error) {

	// Get Unix time in millis
	wl.ServerTime = time.Now()

	_, err := db.store.Collection("wear-logs").Doc(wl.Id).Create(db.ctx, WearLog{
		WatchId:    wl.WatchId,
		PatientId:  wl.PatientId,
		Message:    wl.Message,
		Code:       wl.Code,
		ServerTime: wl.ServerTime,
	})
	if err != nil {
		return WearLog{}, err
	}

	// Return updated object
	return wl, nil
}
