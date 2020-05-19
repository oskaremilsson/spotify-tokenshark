package storeRefreshToken

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/oskaremilsson/spotify-controller/database"
	"github.com/oskaremilsson/spotify-controller/failure"
	"github.com/oskaremilsson/spotify-controller/utils/infoJson"
)

type User struct {
	Username string `json:"username"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	token := r.Form.Get("token")

	tokenParts := strings.Split(token, ".")

	if len(tokenParts) > 2 {
		data, err := jwt.DecodeSegment(tokenParts[1])
		failure.Check(err)

		user := User{}
		json.Unmarshal([]byte(data), &user)

		if user.Username == "" {
			info := infoJson.Parse("Bad token, missing username", false)
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write(info)
			return
		}

		if database.StoreToken(user.Username, token) {
			info := infoJson.Parse(user.Username+"'s token is stored!", true)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(info)
			return
		}

		info := infoJson.Parse("Could not save to database", false)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(info)
	} else {
		info := infoJson.Parse("Bad request", false)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(info)
	}
}
