package dbsetup

import (
	"database/sql"
	"os"

	"github.com/oskaremilsson/spotify-controller/config"
	"github.com/oskaremilsson/spotify-controller/failure"
)

const createTokensTable = `
CREATE TABLE IF NOT EXISTS tokens (
	username string PRIMARY KEY,
   	token json NOT NULL
) WITHOUT ROWID;
`

//const DatabaseFileName = config.DatabaseFolder + "/" + config.DatabaseFileName

func Init() {
	os.MkdirAll(config.DatabaseFolder, 0755)
	os.Create(config.DatabaseFileName)

	db, err := sql.Open("sqlite3", config.DatabaseFileName)
	_, err = db.Exec(createTokensTable)
	failure.Check(err)

	db.Close()
}
