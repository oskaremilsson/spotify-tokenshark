package spotify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/oskaremilsson/spotify-controller/failure"
)

type User struct {
	Username string `json:"id"`
}

func GetCurrentUsername(access_token string) string {
	authHeader := "Bearer " + access_token

	client := &http.Client{}
	r, err := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", authHeader)

	resp, err := client.Do(r)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	failure.Check(err)

	user := User{}
	json.Unmarshal(body, &user)
	return user.Username
}
