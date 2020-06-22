package removeRequest

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
	username := strings.ToLower(r.Form.Get("username"))

	me, err := spotify.WhoAmI(refresh_token)

	if err != nil || username == "" {
		info := infoJson.Parse("Can't get current user or missing username", false)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(info)
		return
	}

	if database.DeleteRequest(me, username) {
		info := infoJson.Parse(me+" nolonger requests "+username, true)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(info)
		return
	}

	info := infoJson.Parse("Could not remove in database", false)
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write(info)
}
