package createGdprConsent

import (
	"net/http"

	"github.com/oskaremilsson/spotify-tokenshark/database"
	"github.com/oskaremilsson/spotify-tokenshark/utils/infoJson"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	id := database.CreateGdprConsent()
	if id != "" {
		info := infoJson.Parse(id, true)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(info)
		return
	}

	info := infoJson.Parse("Could not create GDPR consent", false)
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write(info)
}
