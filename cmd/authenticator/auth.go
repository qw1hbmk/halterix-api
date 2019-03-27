package authenticator

type AccountAuth struct {
	APIKey   string `json:"apiKey"`
	Expiry   string `json:"expiry"`
	Identity string `json:"identity"`
}

func Validate(apiKey string, db *database) bool {

	var aa *AccountAuth
	aa, err := db.GetAuthDetails(apiKey)

	if err != nil || aa == nil {
		return false
	}
	// Check expiry
	// Check Identity

	return true
}
