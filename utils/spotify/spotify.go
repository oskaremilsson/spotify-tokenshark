package spotify

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/oskaremilsson/spotify-tokenshark/config"
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

func WhoAmI(refresh_token string) (string, error) {
	access_token, err := GetAccessToken(refresh_token)
	if err != nil {
		return "", err
	}

	body, err := callApi(access_token, "GET", "https://api.spotify.com/v1/me", nil)
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

func callApi(access_token string, method string, url string, data io.Reader) ([]byte, error) {
	client := &http.Client{}
	r, err := http.NewRequest(method, url, data)
	if err != nil {
		return []byte(""), err
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Bearer "+access_token)

	resp, err := client.Do(r)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func GetAccessToken(refresh_token string) (string, error) {
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refresh_token)

	exchange, err := callExchange(data)

	return exchange.Access_token, err
}

func GetTokensFromCode(code string) (string, string, error) {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", config.SpotifyRedirectUri)

	exchange, err := callExchange(data)

	return exchange.Access_token, exchange.Refresh_token, err
}

func callExchange(data url.Values) (Exchange, error) {
	exchange := Exchange{}

	client := &http.Client{}
	res, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	res.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(config.SpotifyClientString)))

	resp, err := client.Do(res)
	if err != nil {
		return exchange, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return exchange, err
	}

	err = json.Unmarshal(body, &exchange)
	if err != nil {
		return exchange, err
	}

	if resp.StatusCode != 200 {
		fmt.Println(exchange.Error_description)
		return exchange, fmt.Errorf(exchange.Error)
	}

	return exchange, nil
}
