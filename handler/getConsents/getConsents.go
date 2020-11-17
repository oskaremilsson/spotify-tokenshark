package getConsents

import (
	"encoding/json"
	"net/http"

	"github.com/oskaremilsson/spotify-tokenshark/database"
	"github.com/oskaremilsson/spotify-tokenshark/failure"
	"github.com/oskaremilsson/spotify-tokenshark/utils/infoJson"
	"github.com/oskaremilsson/spotify-tokenshark/utils/spotify"
)

type Consents struct {
	Consents []string
	Success  bool
}

func Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	refresh_token := r.Form.Get("refresh_token")

	me, err := spotify.WhoAmI(refresh_token)

	if err != nil {
		info := infoJson.Parse("Bad token", false)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(info)
		return
	}

	consents := Consents{Consents: database.GetConsents(me), Success: true}

	json, err := json.Marshal(consents)
	failure.Check(err)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(json)
}
