package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

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
	fmt.Printf("Server is running...\n")

	http.HandleFunc("/", handler)
	http.HandleFunc("/storeRefreshToken", storeRefreshToken.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
