package watchtower

import (
	"encoding/json"
	"errors"
	"net/http"

	httpwriter "github.com/qw1hbmk/halterix-api/kit/http"

	"github.com/julienschmidt/httprouter"
)

func (s *server) WatchesPatchHandler(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	wr := httpwriter.NewVerboseResponseWriter(w)

	watchId := params.ByName("id")
	if len(watchId) == 0 {
		wr.WriteBadRequest(req, "", errors.New("No watch ID"))
		return
	}

	decoder := json.NewDecoder(req.Body)
	var watch Watch
	if err := decoder.Decode(&watch); err != nil {
		wr.WriteDecodeError(req, err)
		return
	}

	// watchId is an optional body parameter
	if watch.Id != "" {
		if watchId != watch.Id {
			wr.WriteBadRequest(req, "", errors.New("Incorrect watch ID"))
			return
		}
	} else {
		watch.Id = watchId
	}

	// Update watch in firebase store
	// Patch method is only available on recordingId and network fields
	// No update will take place is watch is inactive
	watch, err := s.db.Patch(watch)
	if err != nil {
		wr.WriteInternalServerError(req, "Unable to update watch with id: "+watchId, err)
		return
	}
	wr.WriteResponse(watch)

}

func (s *server) WatchesGetHandler(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	wr := httpwriter.NewVerboseResponseWriter(w)

	watchId := params.ByName("id")
	if len(watchId) == 0 {
		wr.WriteBadRequest(req, "", errors.New("No watch ID"))
		return
	}

	// Save watch to firebase store
	watch, err := s.db.Get(watchId)
	if err != nil {
		wr.WriteInternalServerError(req, "Unable to get watch with id: "+watchId, err)
		return
	}
	wr.WriteResponse(watch)

}
