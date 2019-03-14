package watchtower

import (
	"encoding/json"
	"errors"
	"net/http"

	httpwriter "github.com/qw1hbmk/halterix-api/kit/http"

	"github.com/julienschmidt/httprouter"
)

func (s *server) PatientsPatchHandler(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	wr := httpwriter.NewVerboseResponseWriter(w)

	patientId := params.ByName("id")
	if len(patientId) == 0 {
		wr.WriteBadRequest(req, "", errors.New("No patient ID"))
		return
	}

	decoder := json.NewDecoder(req.Body)
	var p Patient
	if err := decoder.Decode(&p); err != nil {
		wr.WriteDecodeError(req, err)
		return
	}

	// patientId is an optional body parameter
	if p.Id != "" {
		if patientId != p.Id {
			wr.WriteBadRequest(req, "", errors.New("Incorrect patient ID"))
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
		wr.WriteInternalServerError(req, "Unable to update patient with id: "+p.Id, err)
		return
	}
	// We want to keep the patient information confidential, so return original request
	wr.WriteResponse(req.Body)
}

func (s *server) WatchesGetHandler(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	wr := httpwriter.NewVerboseResponseWriter(w)

	watchId := params.ByName("id")
	if len(watchId) == 0 {
		wr.WriteBadRequest(req, "", errors.New("No watch ID"))
		return
	}

	// Save watch to firebase store
	watch, err := s.db.GetWatch(watchId)
	if err != nil {
		wr.WriteInternalServerError(req, "Unable to get watch with id: "+watchId, err)
		return
	}
	wr.WriteResponse(watch)

}
