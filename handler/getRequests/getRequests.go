package getRequests

import (
	"encoding/json"
	"net/http"

	"github.com/oskaremilsson/spotify-controller/database"
	"github.com/oskaremilsson/spotify-controller/failure"
	"github.com/oskaremilsson/spotify-controller/utils/infoJson"
	"github.com/oskaremilsson/spotify-controller/utils/spotify"
)

type Requests struct {
	Requests []string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	refresh_token := r.Form.Get("refresh_token")

	username, err := spotify.WhoAmI(refresh_token)
	failure.Check(err)

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
