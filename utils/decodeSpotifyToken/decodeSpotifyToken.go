package decodeSpotifyToken

import (
	"encoding/json"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/oskaremilsson/spotify-controller/failure"
)

type User struct {
	Username string `json:"username"`
}

func GetUsername(token string) string {
	//info := infoJson.Parse("Bad token, missing username", false)
	tokenParts := strings.Split(token, ".")

	if len(tokenParts) > 2 {
		data, err := jwt.DecodeSegment(tokenParts[1])
		failure.Check(err)

		user := User{}
		json.Unmarshal([]byte(data), &user)
		return user.Username
	} else {
		return "bad_token"
	}

}
