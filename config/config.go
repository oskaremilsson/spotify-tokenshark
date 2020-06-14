package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var DatabaseFolder = os.Getenv("DATA_FOLDER")
var DatabaseFileName = DatabaseFolder + os.Getenv("DATABASE_NAME")
var SpotifyClientString = os.Getenv("SPOTIFY_CLIENT_ID") + ":" + os.Getenv("SPOTIFY_CLIENT_SECRET")
var SpotifyRedirectUri = os.Getenv("SPOTIFY_REDIRECT_URI")
var Port = os.Getenv("PORT")
var AllowOrigin = os.Getenv("ALLOW_ORIGIN")
var DataSecret = os.Getenv("DATA_SECRET")
