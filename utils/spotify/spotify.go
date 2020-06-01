package spotify

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/oskaremilsson/spotify-controller/config"
)

type User struct {
	Username string `json:"id"`
}

type Exchange struct {
	Refresh_token     string `json:"refresh_token"`
	Access_token      string `json:"access_token"`
	Error             string `json:"error"`
	Error_description string `json:"error_description"`
}

func GetCurrentUsername(access_token string) (string, error) {
	authHeader := "Bearer " + access_token

	client := &http.Client{}
	r, err := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)
	if err != nil {
		return "", err
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", authHeader)

	resp, err := client.Do(r)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	user := User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return "", err
	}

	return user.Username, nil
}

func GetAccessTokenFromRefreshToken(refresh_token string) (string, error) {
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refresh_token)

	client := &http.Client{}
	res, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	res.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(config.SpotifyClientString)))

	resp, err := client.Do(res)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	exchange := Exchange{}
	err = json.Unmarshal(body, &exchange)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 || exchange.Access_token == "" {
		return "", fmt.Errorf(exchange.Error)
	}

	return exchange.Access_token, nil
}

func GetTokensFromCode(code string) (string, string, error) {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", os.Getenv("SPOTIFY_REDIRECT_URI"))

	client := &http.Client{}
	res, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	res.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(config.SpotifyClientString)))

	resp, err := client.Do(res)
	if err != nil {
		return "", "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	exchange := Exchange{}
	err = json.Unmarshal(body, &exchange)
	if err != nil {
		return "", "", err
	}

	if resp.StatusCode != 200 || exchange.Refresh_token == "" {
		return "", "", fmt.Errorf(exchange.Error)
	}

	return exchange.Access_token, exchange.Refresh_token, nil
}
