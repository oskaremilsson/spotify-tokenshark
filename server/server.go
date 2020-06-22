package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"

	"github.com/oskaremilsson/spotify-controller/config"
	"github.com/oskaremilsson/spotify-controller/database/dbsetup"
	"github.com/oskaremilsson/spotify-controller/handler/acceptRequest"
	"github.com/oskaremilsson/spotify-controller/handler/codeExchange"
	"github.com/oskaremilsson/spotify-controller/handler/createRequest"
	"github.com/oskaremilsson/spotify-controller/handler/getAccessToken"
	"github.com/oskaremilsson/spotify-controller/handler/getConsents"
	"github.com/oskaremilsson/spotify-controller/handler/getRequests"
	"github.com/oskaremilsson/spotify-controller/handler/giveConsent"
	"github.com/oskaremilsson/spotify-controller/handler/removeRequest"
	"github.com/oskaremilsson/spotify-controller/handler/revokeConsent"
	"github.com/oskaremilsson/spotify-controller/handler/storeRefreshToken"
)

type HealthCheck struct {
	Time time.Time
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
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

	addr := fmt.Sprintf(":%s", config.Port)
	router := httprouter.New()
	handler := negroni.New()
	handler.UseHandler(router)

	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	router.HandlerFunc("GET", "/", cors(statusHandler))
	router.HandlerFunc("POST", "/codeExchange", cors(codeExchange.Handler))
	router.HandlerFunc("POST", "/storeRefreshToken", cors(storeRefreshToken.Handler))
	router.HandlerFunc("POST", "/giveConsent", cors(giveConsent.Handler))
	router.HandlerFunc("POST", "/revokeConsent", cors(revokeConsent.Handler))
	router.HandlerFunc("POST", "/createRequest", cors(createRequest.Handler))
	router.HandlerFunc("POST", "/getRequests", cors(getRequests.Handler))
	router.HandlerFunc("POST", "/getConsents", cors(getConsents.Handler))
	router.HandlerFunc("POST", "/getAccessToken", cors(getAccessToken.Handler))
	router.HandlerFunc("POST", "/removeRequest", cors(removeRequest.Handler))
	router.HandlerFunc("POST", "/acceptRequest", cors(acceptRequest.Handler))

	router.HandleOPTIONS = true
	router.GlobalOPTIONS = http.HandlerFunc(optionsCors)

	fmt.Printf("Server is running...\n")
	log.Fatal(server.ListenAndServe())
}

func cors(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", config.AllowOrigin)
		handler(w, r)
	}
}

func optionsCors(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Access-Control-Request-Method") != "" {
		header := w.Header()
		header.Set("Access-Control-Allow-Origin", config.AllowOrigin)
	}

	w.WriteHeader(http.StatusNoContent)
}
