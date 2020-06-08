package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/oskaremilsson/spotify-controller/config"
	"github.com/oskaremilsson/spotify-controller/database/dbsetup"
	"github.com/oskaremilsson/spotify-controller/handler/codeExchange"
	"github.com/oskaremilsson/spotify-controller/handler/createRequest"
	"github.com/oskaremilsson/spotify-controller/handler/getAccessToken"
	"github.com/oskaremilsson/spotify-controller/handler/getRequests"
	"github.com/oskaremilsson/spotify-controller/handler/giveConsent"
	"github.com/oskaremilsson/spotify-controller/handler/revokeConsent"
	"github.com/oskaremilsson/spotify-controller/handler/storeRefreshToken"
)

type HealthCheck struct {
	Time time.Time
}

func handler(w http.ResponseWriter, r *http.Request) {
	hk := HealthCheck{time.Now()}

	js, err := json.Marshal(hk)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK) // Respond 200
	_, _ = w.Write(js)
}

func main() {
	fmt.Printf("Setting up database...\n")
	dbsetup.Init()

	fmt.Printf("Registering handlers...\n")
	http.HandleFunc("/", handler)
	http.HandleFunc("/codeExchange", codeExchange.Handler)
	http.HandleFunc("/storeRefreshToken", storeRefreshToken.Handler)
	http.HandleFunc("/giveConsent", giveConsent.Handler)
	http.HandleFunc("/revokeConsent", revokeConsent.Handler)
	http.HandleFunc("/createRequest", createRequest.Handler)
	http.HandleFunc("/getRequests", getRequests.Handler)
	http.HandleFunc("/getAccessToken", getAccessToken.Handler)

	fmt.Printf("Server is running...\n")
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
