package getMyConsents

import (
	"encoding/json"
	"net/http"

	"github.com/oskaremilsson/spotify-controller/database"
	"github.com/oskaremilsson/spotify-controller/failure"
	"github.com/oskaremilsson/spotify-controller/utils/infoJson"
	"github.com/oskaremilsson/spotify-controller/utils/spotify"
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

	consents := Consents{Consents: database.GetMyConsents(me), Success: true}

	json, err := json.Marshal(consents)
	failure.Check(err)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(json)
}
