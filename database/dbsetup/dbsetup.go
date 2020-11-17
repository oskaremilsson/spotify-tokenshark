package dbsetup

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/oskaremilsson/spotify-tokenshark/config"
	"github.com/oskaremilsson/spotify-tokenshark/failure"
)

const createTokensTable = `
CREATE TABLE IF NOT EXISTS tokens (
	username TEXT PRIMARY KEY,
  token TEXT NOT NULL
);
`

const createConsentsTable = `
CREATE TABLE IF NOT EXISTS consents (
	username TEXT NOT NULL,
	allow_user TEXT NOT NULL,
	UNIQUE(username,allow_user)
);
`

const createRequestsTable = `
CREATE TABLE IF NOT EXISTS requests (
	username TEXT NOT NULL,
	requesting TEXT NOT NULL,
	UNIQUE(username,requesting)
);
`

func Init() {
	os.MkdirAll(config.DatabaseFolder, 0755)
	db, err := sql.Open("sqlite3", createDatabaseFile())

	_, err = db.Exec(createTokensTable)
	failure.Check(err)

	_, err = db.Exec(createConsentsTable)
	failure.Check(err)

	_, err = db.Exec(createRequestsTable)
	failure.Check(err)

	db.Close()
}

func createDatabaseFile() string {
	_, err := os.Stat(config.DatabaseFileName)
	if os.IsNotExist(err) {
		os.Create(config.DatabaseFileName)
	}
	return config.DatabaseFileName
}
