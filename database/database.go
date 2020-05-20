package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/oskaremilsson/spotify-controller/config"
	"github.com/oskaremilsson/spotify-controller/failure"
)

func StoreToken(username string, token string) bool {
	db, err := sql.Open("sqlite3", config.DatabaseFileName)
	failure.Check(err)

	stmt, err := db.Prepare("INSERT INTO tokens(username, token) values(?,?)")
	failure.Check(err)

	_, err = stmt.Exec(username, token)
	if err != nil {
		stmt, err = db.Prepare("UPDATE tokens SET token = ? WHERE username = ?")
		_, err = stmt.Exec(token, username)
		if err != nil {
			db.Close()
			return false
		}
		db.Close()
		return true
	}

	db.Close()
	return true
}

func StoreConsent(username string, allowed string) bool {
	db, err := sql.Open("sqlite3", config.DatabaseFileName)
	failure.Check(err)

	stmt, err := db.Prepare("INSERT INTO consent(username, allow_user) values(?,?)")
	failure.Check(err)

	_, err = stmt.Exec(username, allowed)
	if err != nil {
		db.Close()
		return false
	}

	db.Close()
	return true
}
