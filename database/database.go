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

func StoreConsent(username string, allow_user string) bool {
	db, err := sql.Open("sqlite3", config.DatabaseFileName)
	failure.Check(err)

	stmt, err := db.Prepare("INSERT INTO consents(username, allow_user) values(?,?)")
	failure.Check(err)

	_, err = stmt.Exec(username, allow_user)
	if err != nil {
		db.Close()
		return false
	}

	db.Close()
	return true
}

func CreateRequest(username string, requesting string) bool {
	db, err := sql.Open("sqlite3", config.DatabaseFileName)
	failure.Check(err)

	stmt, err := db.Prepare("INSERT INTO requests(username, requesting) values(?,?)")
	failure.Check(err)

	_, err = stmt.Exec(username, requesting)
	if err != nil {
		db.Close()
		return false
	}

	db.Close()
	return true
}
