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

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
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

	router.HandlerFunc("GET", "/", statusHandler)
	router.HandlerFunc("POST", "/codeExchange", codeExchange.Handler)
	router.HandlerFunc("POST", "/storeRefreshToken", storeRefreshToken.Handler)
	router.HandlerFunc("POST", "/giveConsent", giveConsent.Handler)
	router.HandlerFunc("POST", "/revokeConsent", revokeConsent.Handler)
	router.HandlerFunc("POST", "/createRequest", createRequest.Handler)
	router.HandlerFunc("POST", "/getRequests", getRequests.Handler)
	router.HandlerFunc("POST", "/getAccessToken", getAccessToken.Handler)

	router.HandleOPTIONS = true
	router.GlobalOPTIONS = http.HandlerFunc(Cors)

	fmt.Printf("Server is running...\n")
	log.Fatal(server.ListenAndServe())
}

func Cors(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Access-Control-Request-Method") != "" {
		header := w.Header()
		//header.Set("Access-Control-Allow-Headers", "CONTENT-TYPE")
		header.Set("Access-Control-Allow-Origin", "*")
	}

	w.WriteHeader(http.StatusNoContent)
}
