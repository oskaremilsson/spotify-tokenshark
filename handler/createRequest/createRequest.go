package createRequest

import (
	"net/http"
	"strings"

	"github.com/oskaremilsson/spotify-tokenshark/database"
	"github.com/oskaremilsson/spotify-tokenshark/utils/infoJson"
	"github.com/oskaremilsson/spotify-tokenshark/utils/spotify"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	refresh_token := r.Form.Get("refresh_token")
	requesting := strings.ToLower(r.Form.Get("requesting"))

	me, err := spotify.WhoAmI(refresh_token)

	if err != nil || requesting == "" {
		info := infoJson.Parse("Could not get current username or requesting", false)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(info)
		return
	}

	if !database.IsServiceUser(requesting) {
		info := infoJson.Parse(requesting+" is not a user of service", false)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(info)
		return
	}

	if database.CreateRequest(me, requesting) {
		info := infoJson.Parse(me+" have requested "+requesting, true)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(info)
		return
	}

	info := infoJson.Parse("Could not save to database", false)
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write(info)
}
