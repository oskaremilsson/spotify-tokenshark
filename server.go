package main

import (
  "fmt"
  "github.com/dgrijalva/jwt-go"
  "net/http"
  "log"
  "encoding/json"
  "strings"
  "time"
  "github.com/oskaremilsson/spotify-controller/StoreRereshToken"
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
  fmt.Printf("booting server...\n")
  
  http.HandleFunc("/", handler)
  http.HandleFunc("/storeRefreshToken", StoreRereshToken.run)
  log.Fatal(http.ListenAndServe(":8080", nil))
  fmt.Printf("server running! \n")
}
