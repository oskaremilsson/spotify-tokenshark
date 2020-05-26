package storeRefreshToken

import (
	"net/http"

	"github.com/oskaremilsson/spotify-controller/database"
	"github.com/oskaremilsson/spotify-controller/utils/decodeSpotifyToken"
	"github.com/oskaremilsson/spotify-controller/utils/infoJson"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	token := r.Form.Get("token")

	username := decodeSpotifyToken.GetUsername(token)

	if username == "-" {
		info := infoJson.Parse("Bad token, missing username", false)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(info)
		return
	}

	if database.StoreToken(username, token) {
		info := infoJson.Parse(username+"'s token is stored!", true)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(info)
		return
	}

	info := infoJson.Parse("Could not save to database", false)
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write(info)
}
