package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var SpotifyClientString = os.Getenv("SPOTIFY_CLIENT_ID") + ":" + os.Getenv("SPOTIFY_CLIENT_SECRET")
var SpotifyRedirectUri = os.Getenv("SPOTIFY_REDIRECT_URI")
var Port = os.Getenv("PORT")
var AllowOrigin = os.Getenv("ALLOW_ORIGIN")
var EncryptionSecretKey = os.Getenv("ENCRYPTION_SECRET_KEY")
var DatabaseUrl = os.Getenv("DATABASE_URL") + "?sslmode=disable"
