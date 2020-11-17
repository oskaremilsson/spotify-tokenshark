package codeExchange

import (
	"encoding/json"
	"net/http"

	"github.com/oskaremilsson/spotify-tokenshark/failure"
	"github.com/oskaremilsson/spotify-tokenshark/utils/infoJson"
	"github.com/oskaremilsson/spotify-tokenshark/utils/spotify"
)

type Tokens struct {
	Refresh_token string
	Access_token  string
	Success       bool
}

func Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	access_token, refresh_token, err := spotify.GetTokensFromCode(r.Form.Get("code"))

	if err != nil {
		info := infoJson.Parse(err.Error(), false)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(info)
		return
	}

	data := Tokens{
		Refresh_token: refresh_token,
		Access_token:  access_token,
		Success:       true,
	}
	js, err := json.Marshal(data)
	failure.Check(err)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(js)
}
