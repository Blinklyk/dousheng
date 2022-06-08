package utils

import (
	"errors"
	"os"
)

// PathExists check if the path exists
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, errors.New("error: path" + err.Error())
}
