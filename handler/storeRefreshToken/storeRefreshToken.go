package storeRefreshToken

import (
	"net/http"

	"github.com/oskaremilsson/spotify-tokenshark/database"
	"github.com/oskaremilsson/spotify-tokenshark/utils/crypto"
	"github.com/oskaremilsson/spotify-tokenshark/utils/infoJson"
	"github.com/oskaremilsson/spotify-tokenshark/utils/spotify"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	refresh_token := r.Form.Get("refresh_token")
	gdpr_consent := r.Form.Get("gdpr_consent")
	encrypted_token := crypto.Encrypt([]byte(refresh_token))

	me, err := spotify.WhoAmI(refresh_token)

	if err != nil {
		info := infoJson.Parse("Could not get current user", false)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(info)
		return
	}

	if !database.ConnectGdprConsent(me, gdpr_consent) {
		info := infoJson.Parse("Could not connect user to gdpr consent", false)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(info)
		return
	}

	if database.StoreToken(me, string(encrypted_token)) {
		info := infoJson.Parse(me+"'s refresh_token is stored encrypted!", true)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(info)
		return
	}

	info := infoJson.Parse("Could not store refresh token", false)
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write(info)
}
