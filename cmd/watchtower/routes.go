package watchtower

import (
	"github.com/julienschmidt/httprouter"
)

type server struct {
	router *httprouter.Router
	db     *database
}

func NewServer(r *httprouter.Router, s *database) *server {
	return &server{r, s}
}

func (s *server) RegisterRoutes() {
	s.router.PATCH("/watches/:id", s.WatchesPatchHandler)
	s.router.GET("/watches/:id", s.WatchesGetHandler)
}
