package watchtower

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/qw1hbmk/halterix-api/kit"
	httpwriter "github.com/qw1hbmk/halterix-api/kit/http"
	"github.com/qw1hbmk/halterix-api/kit/util"

	"github.com/julienschmidt/httprouter"
)

func (s *server) PatientsPatchHandler(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	wr := httpwriter.NewVerboseResponseWriter(w)

	// Should not be possible, but add check anyway
	patientId := params.ByName("id")
	if len(patientId) == 0 {
		err := errors.New("No patient ID")
		kit.LogBadRequestError(req, err)
		wr.WriteBadRequest(req, "", errors.New("No patient ID"))
		return
	}

	decoder := json.NewDecoder(req.Body)
	var p Patient
	if err := decoder.Decode(&p); err != nil {
		kit.LogBadRequestError(req, err)
		wr.WriteDecodeError(req, err)
		return
	}

	// patientId is an optional body parameter
	if p.Id != "" {
		if patientId != p.Id {
			err := errors.New("Incorrect patient ID")
			kit.LogBadRequestError(req, err)
			wr.WriteBadRequest(req, "", err)
			return
		}
	} else {
		p.Id = patientId
	}

	// Update patient in firebase store
	// Patch method is only available on recordingId and network fields
	// No update will take place is watch is inactive
	p, err := s.db.PatchPatient(p)
	if err != nil {
		msg := "Unable to update patient with id: " + p.Id
		kit.LogInternalServerError(req, msg, err)
		wr.WriteInternalServerError(req, msg, err)
		return
	}
	// We want to keep the patient information confidential, so return original request
	wr.WriteResponse(req.Body)
}

func (s *server) WatchesGetHandler(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	wr := httpwriter.NewVerboseResponseWriter(w)

	watchId := params.ByName("id")
	if len(watchId) == 0 {
		err := errors.New("No watch ID")
		kit.LogBadRequestError(req, err)
		wr.WriteBadRequest(req, "", err)
		return
	}

	// Save watch to firebase store
	watch, err := s.db.GetWatch(watchId)
	if err != nil {
		msg := "Unable to get watch with id: " + watchId
		kit.LogInternalServerError(req, msg, err)
		wr.WriteInternalServerError(req, msg, err)
		return
	}
	wr.WriteResponse(watch)

}

func (s *server) WearLogPostHandler(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	wr := httpwriter.NewVerboseResponseWriter(w)
	decoder := json.NewDecoder(req.Body)
	var wl WearLog
	if err := decoder.Decode(&wl); err != nil {
		kit.LogBadRequestError(req, err)
		wr.WriteDecodeError(req, err)
		return
	}

	uuid, err := util.Uuid()
	if err != nil {
		msg := "Could not generate UUID"
		kit.LogInternalServerError(req, msg, err)
		wr.WriteInternalServerError(req, msg, err)
		return
	}

	wl.Id = uuid

	// Update patient in firebase store
	// Patch method is only available on recordingId and network fields
	// No update will take place is watch is inactive
	wl, err = s.db.PostWearLog(wl)
	if err != nil {
		msg := "Unable to update wearlog with id: " + wl.Id
		kit.LogInternalServerError(req, msg, err)
		wr.WriteInternalServerError(req, msg, err)
		return
	}
	// Return wearlog
	wr.WriteResponse(wl)
}
