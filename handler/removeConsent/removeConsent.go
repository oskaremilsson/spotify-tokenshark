package removeConsent

import (
	"net/http"
	"strings"

	"github.com/oskaremilsson/spotify-controller/database"
	"github.com/oskaremilsson/spotify-controller/utils/infoJson"
	"github.com/oskaremilsson/spotify-controller/utils/spotify"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	refresh_token := r.Form.Get("refresh_token")
	remove_user := strings.ToLower(r.Form.Get("remove_user"))

	me, err := spotify.WhoAmI(refresh_token)

	if err != nil || remove_user == "" {
		info := infoJson.Parse("Can't get current user or missing remove_user", false)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(info)
		return
	}

	if database.DeleteConsent(remove_user, me) {
		info := infoJson.Parse(me+" removed consent to "+remove_user, true)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(info)
		return
	}

	info := infoJson.Parse("Could not remove consent in database", false)
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write(info)
}
