package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

const DatabaseFolder = "./data/"
const DatabaseFileName = DatabaseFolder + "/data.db"

var SpotifyClientString = os.Getenv("SPOTIFY_CLIENT_ID") + ":" + os.Getenv("SPOTIFY_CLIENT_SECRET")
