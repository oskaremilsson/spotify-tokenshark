package getAccessToken

import (
	"encoding/json"
	"net/http"

	"github.com/oskaremilsson/spotify-controller/database"
	"github.com/oskaremilsson/spotify-controller/failure"
	"github.com/oskaremilsson/spotify-controller/utils/crypto"
	"github.com/oskaremilsson/spotify-controller/utils/infoJson"
	"github.com/oskaremilsson/spotify-controller/utils/spotify"
)

type Response struct {
	Access_token string
	Success      bool
}

func Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	refresh_token := r.Form.Get("refresh_token")
	username := r.Form.Get("username")

	me, err := spotify.WhoAmI(refresh_token)

	if err != nil {
		info := infoJson.Parse("Could not find current user", false)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(info)
		return
	}

	if !database.ValidateConsent(me, username) && me != username {
		info := infoJson.Parse("You don't have the proper consent", false)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(info)
		return
	}

	user_token, err := database.GetRefreshToken(username)
	if err != nil {
		info := infoJson.Parse("Could not find users refresh token", false)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(info)
		return
	}

	user_token = string(crypto.Decrypt([]byte(user_token)))

	access_token, err := spotify.GetAccessToken(user_token)

	if err != nil {
		info := infoJson.Parse("Could not get access token", false)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(info)
		return
	}

	response := Response{Access_token: access_token, Success: true}

	json, err := json.Marshal(response)
	failure.Check(err)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(json)
}
