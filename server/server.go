package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/oskaremilsson/spotify-controller/database"
	"github.com/oskaremilsson/spotify-controller/failure"
)

type HealthCheck struct {
	Time time.Time
}

type Info struct {
	Message string
	Success bool
}

type User struct {
	Username string `json:"username"`
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
		failure.Check(err)

		user := User{}
		json.Unmarshal([]byte(data), &user)

		if user.Username == "" {
			info := infoJSON("Bad token, missing username", false)
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write(info)
			return
		}

		if database.StoreToken(user.Username, token) {
			info := infoJSON(user.Username+"'s token is stored!", true)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(info)
			return
		}

		info := infoJSON("Could not save to database", false)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(info)
	} else {
		info := infoJSON("Bad request", false)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(info)
	}
}

func main() {
	fmt.Printf("Server is running...\n")

	http.HandleFunc("/", handler)
	http.HandleFunc("/storeRefreshToken", StoreRefreshToken)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func infoJSON(message string, success bool) []byte {
	info := Info{Message: message, Success: success}
	js, err := json.Marshal(info)
	failure.Check(err)

	return js
}
