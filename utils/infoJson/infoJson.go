package infoJson

import (
	"encoding/json"

	"github.com/oskaremilsson/spotify-controller/failure"
)

type Info struct {
	Message string
	Success bool
}

func Parse(message string, success bool) []byte {
	info := Info{Message: message, Success: success}
	js, err := json.Marshal(info)
	failure.Check(err)

	return js
}
