package createGdprConsent

import (
	"net/http"
	"time"

	"github.com/oskaremilsson/spotify-tokenshark/database"
	"github.com/oskaremilsson/spotify-tokenshark/utils/infoJson"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	id := database.CreateGdprConsent()
	if id != "" {
		info := infoJson.Parse(id, true)
		addCookie(w, "gdpr_consent", id)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(info)
		return
	}

	info := infoJson.Parse("Could not create GDPR consent", false)
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write(info)
}

func addCookie(w http.ResponseWriter, name, value string) {
	expire := time.Now().AddDate(1, 0, 0)
	cookie := http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expire,
	}
	http.SetCookie(w, &cookie)
}
