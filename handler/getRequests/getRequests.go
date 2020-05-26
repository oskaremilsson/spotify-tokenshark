package getRequests

import (
	"encoding/json"
	"net/http"

	"github.com/oskaremilsson/spotify-controller/database"
	"github.com/oskaremilsson/spotify-controller/failure"
	"github.com/oskaremilsson/spotify-controller/utils/decodeSpotifyToken"
	"github.com/oskaremilsson/spotify-controller/utils/infoJson"
)

type Requests struct {
	Requests []string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	token := r.Form.Get("token")

	username := decodeSpotifyToken.GetUsername(token)

	if username == "bad_token" {
		info := infoJson.Parse("Bad token", false)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(info)
		return
	}

	requests := Requests{Requests: database.GetRequests(username)}

	json, err := json.Marshal(requests)
	failure.Check(err)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(json)
}
