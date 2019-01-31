package watchtower

import (
	"encoding/json"
	"net/http"
	"errors"

	"github.com/julienschmidt/httprouter"
	httpwriter "github.com/qw1hbmk/halterix-api/kit/http"
	
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
	err := s.db.Update(watch)
	if err != nil {
		wr.WriteInternalServerError(req, "Unable to update watch with id: " + watchId, err)
		return
	}
	wr.WriteResponse(watch)

}

// // Get all the watches. Able to filter by active watches, watches on a specific network, minutes since last active
// func (s *server) WatchesGetHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
// 	wr := httpwriter.NewVerboseResponseWriter(w)

// 	id := ps.ByName("id")
// 	if len(id) == 0 {
// 		// Unclear how we would ever get here, but just in case...
// 		ew.BadRequest(r, "Bad rule id.")
// 		return
// 	}

// 	callerID := "123"	
// 	params := req.URL.Query()
// 	aw := params["active"]
// 	net := params["netowrk"]
// 	from := params["from"]
// 	// Get all patient data from distrubuted store

// 	ps := store.GetWatches(clinician, organization, callerId string)

// 	js, err := json.Marshal(ps)
// 	if err != nil {
// 		wr.WriteJSONMarshalError(req, ps, err)
// 		return
// 	}

// 	w.WriteResponse(js)
// })

