package watchtower

import (
	"encoding/json"
	"errors"
	"net/http"

	httpwriter "github.com/qw1hbmk/halterix-api/kit/http"

	"github.com/julienschmidt/httprouter"
)

func (s *server) WatchesPutHandler(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
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

	if watchId != watch.Id {
		wr.WriteBadRequest(req, "", errors.New("Incorrect watch ID"))
		return
	}

	// Save watch to firebase store
	watch, err := s.db.Update(watch)
	if err != nil {
		wr.WriteInternalServerError(req, "Unable to update watch with id: "+watchId, err)
		return
	}
	wr.WriteResponse(watch)

}
