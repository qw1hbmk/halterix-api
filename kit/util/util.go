package util

import gouuid "github.com/nu7hatch/gouuid"

func Uuid() (string, error) {
	uuid, err := gouuid.NewV4()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}
