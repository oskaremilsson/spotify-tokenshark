package revokeConsent

import (
	"net/http"

	"github.com/oskaremilsson/spotify-controller/database"
	"github.com/oskaremilsson/spotify-controller/utils/infoJson"
	"github.com/oskaremilsson/spotify-controller/utils/spotify"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	refresh_token := r.Form.Get("refresh_token")
	disallow_user := r.Form.Get("disallow_user")

	me, err := spotify.WhoAmI(refresh_token)

	if err != nil || disallow_user == "" {
		info := infoJson.Parse("Can't get current user or missing disallow_user", false)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(info)
		return
	}

	if database.DeleteConsent(me, disallow_user) {
		info := infoJson.Parse(me+" disallows "+disallow_user, true)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(info)
		return
	}

	info := infoJson.Parse("Could not revoke in database", false)
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write(info)
}
