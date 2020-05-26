package createRequest

import (
	"fmt"
	"net/http"

	"github.com/oskaremilsson/spotify-controller/database"
	"github.com/oskaremilsson/spotify-controller/utils/decodeSpotifyToken"
	"github.com/oskaremilsson/spotify-controller/utils/infoJson"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	token := r.Form.Get("token")
	requesting := r.Form.Get("requesting")

	username := decodeSpotifyToken.GetUsername(token)

	fmt.Printf(username + "\n")
	fmt.Printf(requesting + "\n")

	if username == "bad_token" || requesting == "" {
		info := infoJson.Parse("Missing username or requesting", false)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(info)
		return
	}

	if database.CreateRequest(username, requesting) {
		info := infoJson.Parse(username+" have requested "+requesting, true)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(info)
		return
	}

	info := infoJson.Parse("Could not save to database", false)
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write(info)
}
