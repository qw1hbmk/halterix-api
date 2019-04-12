package authenticator

import (
	"errors"
	"fmt"
	"time"

	"github.com/qw1hbmk/halterix-api/kit"
)

type AccountAuth struct {
	APIKey string     `json:"apiKey"`
	Expiry *time.Time `json:"expiry"`
	Role   string     `json:"role"`
}

func Validate(apiKey string, db *database) bool {

	var aa *AccountAuth
	aa, err := db.GetAuthDetails(apiKey)
	if aa.Expiry == nil {
		if aa.Role != "admin" {
			kit.LogInfo(aa.Role)
			kit.LogInfo(aa.Role == "admin")
			kit.LogError(errors.New("Auth failed due to malicious access attempt by role " + aa.Role))
			return false
		}
	} else {
		if aa.Expiry.Sub(time.Now()) < time.Duration(0) {
			kit.LogError(errors.New("Auth failed due to expired keys."))
			return false
		}
	}

	if err != nil || aa == nil {
		kit.LogError(errors.New(fmt.Sprintf("Auth failed. %v", err)))
		return false
	}

	return true
}
