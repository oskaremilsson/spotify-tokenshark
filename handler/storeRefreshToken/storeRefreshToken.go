package storeRefreshToken

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/oskaremilsson/spotify-controller/config"
	"github.com/oskaremilsson/spotify-controller/database"
	"github.com/oskaremilsson/spotify-controller/failure"
	"github.com/oskaremilsson/spotify-controller/utils/infoJson"
	"github.com/oskaremilsson/spotify-controller/utils/spotify"
)

type Response struct {
	Refresh_token     string `json:"refresh_token"`
	Access_token      string `json:"access_token"`
	Error             string `json:"error"`
	Error_description string `json:"error_description"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", r.Form.Get("code"))
	data.Set("redirect_uri", os.Getenv("REDIRECT_URI"))

	client := &http.Client{}
	res, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	res.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(config.SpotifyClientString)))

	resp, err := client.Do(res)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	failure.Check(err)

	response := Response{}
	err = json.Unmarshal(body, &response)
	failure.Check(err)

	if resp.StatusCode != 200 || response.Refresh_token == "" {
		info := infoJson.Parse(response.Error, false)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(info)
		return
	}

	username := spotify.GetCurrentUsername(response.Access_token)

	if username != "" && database.StoreToken(username, response.Refresh_token) {
		info := infoJson.Parse(username+"'s refresh_token is stored!", true)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(info)
		return
	}

	info := infoJson.Parse("Could not store refresh token", false)
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write(info)
}
