package acceptRequest

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
	allow_user := strings.ToLower(r.Form.Get("allow_user"))

	me, err := spotify.WhoAmI(refresh_token)

	if err != nil || allow_user == "" {
		info := infoJson.Parse("Can't get current user or missing allow_user", false)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(info)
		return
	}

	if database.StoreConsent(me, allow_user) {
		info := infoJson.Parse(me+" now allows "+allow_user, true)

		database.DeleteRequest(me, allow_user)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(info)
		return
	}

	info := infoJson.Parse("Could not save to database", false)
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write(info)
}
