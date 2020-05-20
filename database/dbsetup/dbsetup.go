package dbsetup

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/oskaremilsson/spotify-controller/config"
	"github.com/oskaremilsson/spotify-controller/failure"
)

const createTokensTable = `
CREATE TABLE IF NOT EXISTS tokens (
	username string PRIMARY KEY,
   	token json NOT NULL
) WITHOUT ROWID;
`

func Init() {
	os.MkdirAll(config.DatabaseFolder, 0755)

	db, err := sql.Open("sqlite3", createDatabaseFile())
	_, err = db.Exec(createTokensTable)
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
