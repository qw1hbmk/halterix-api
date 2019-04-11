package kit

import (
	"net/http"

	httpwriter "github.com/qw1hbmk/halterix-api/kit/http"
	log "github.com/sirupsen/logrus"
)

func LogBadRequestError(req *http.Request, err error) {
	log.WithFields(log.Fields{
		"req":  httpwriter.FormatRequest(req),
		"code": 400,
	}).Error(err)
}

func LogInternalServerError(req *http.Request, msg string, err error) {
	log.WithFields(log.Fields{
		"req":     httpwriter.FormatRequest(req),
		"code":    500,
		"details": msg,
	}).Error(err)
}

func LogUnauthorizedError(req *http.Request, err error) {
	log.WithFields(log.Fields{
		"req":  httpwriter.FormatRequest(req),
		"code": 401,
	}).Error(err)
}

func LogInfo(msg interface{}) {
	log.Info(msg)
}

func LogError(err error) {
	log.Error(err)
}
