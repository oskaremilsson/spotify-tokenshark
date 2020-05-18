package main

import (
  "fmt"
  "net/http"
  "log"
  "encoding/json"
  "time"
  "github.com/dgrijalva/jwt-go"
  "strings"
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

func StoreRefreshToken(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
	token := r.Form.Get("token")

	tokenParts := strings.Split(token, ".")

	if len(tokenParts) > 2 {
		data, err := jwt.DecodeSegment(tokenParts[1])
		if err != nil {
		fmt.Println("error:", err)
		return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(data)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Bad request"))
  }
}

func main() {
  fmt.Printf("booting server...\n")
  
  http.HandleFunc("/", handler)
  http.HandleFunc("/storeRefreshToken", StoreRefreshToken)
  log.Fatal(http.ListenAndServe(":8080", nil))
  fmt.Printf("server running! \n")
}
