package storeRefreshToken

import (
	"net/http"

	"github.com/oskaremilsson/spotify-controller/database"
	"github.com/oskaremilsson/spotify-controller/utils/infoJson"
	"github.com/oskaremilsson/spotify-controller/utils/spotify"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	refresh_token := r.Form.Get("refresh_token")

	// validate that refresh token is valid by getting access token from it
	access_token, err := spotify.GetAccessTokenFromRefreshToken(refresh_token)
	if err != nil {
		info := infoJson.Parse(err.Error(), false)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(info)
		return
	}

	username, err := spotify.GetCurrentUsername(access_token)

	if err != nil {
		info := infoJson.Parse("Could not get current user", false)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(info)
		return
	}

	if username != "" && database.StoreToken(username, refresh_token) {
		info := infoJson.Parse(username+"'s refresh_token is stored!", true)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(info)
		return
	}

	info := infoJson.Parse("Could not store refresh token", false)
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write(info)
}
