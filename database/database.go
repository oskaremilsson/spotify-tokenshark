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

func DeleteConsent(username string, disallow_user string) bool {
	db, err := sql.Open("sqlite3", config.DatabaseFileName)
	failure.Check(err)

	stmt, err := db.Prepare("DELETE FROM consents WHERE username = ? AND allow_user = ?")
	failure.Check(err)

	_, err = stmt.Exec(username, disallow_user)
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

func GetRequests(requesting string) []string {
	db, err := sql.Open("sqlite3", config.DatabaseFileName)
	failure.Check(err)

	stmt, err := db.Prepare("SELECT * FROM requests WHERE requesting = ?")
	failure.Check(err)

	usernames := []string{}

	rows, err := stmt.Query(requesting)
	if err != nil {
		db.Close()
		return usernames
	}

	defer rows.Close()
	for rows.Next() {
		var username string
		var requesting string
		err = rows.Scan(&username, &requesting)
		failure.Check(err)
		usernames = append(usernames, username)
	}

	db.Close()
	return usernames
}

func GetConsents(username string) []string {
	db, err := sql.Open("sqlite3", config.DatabaseFileName)
	failure.Check(err)

	stmt, err := db.Prepare("SELECT * FROM consents WHERE allow_user = ?")
	failure.Check(err)

	usernames := []string{}

	rows, err := stmt.Query(username)
	if err != nil {
		db.Close()
		return usernames
	}

	defer rows.Close()
	for rows.Next() {
		var username string
		var allow_user string
		err = rows.Scan(&username, &allow_user)
		failure.Check(err)
		usernames = append(usernames, username)
	}

	db.Close()
	return usernames
}

func ValidateConsent(me string, username string) bool {
	db, err := sql.Open("sqlite3", config.DatabaseFileName)
	failure.Check(err)

	sqlStmt := "SELECT * FROM consents WHERE username = ? AND allow_user = ?"

	err = db.QueryRow(sqlStmt, username, me).Scan(&username, &me)
	db.Close()

	return err == nil
}

func GetRefreshToken(username string) (string, error) {
	db, err := sql.Open("sqlite3", config.DatabaseFileName)
	failure.Check(err)

	stmt, err := db.Prepare("SELECT token FROM tokens WHERE username = ?")
	failure.Check(err)

	refresh_token := ""

	err = stmt.QueryRow(username).Scan(&refresh_token)
	if err != nil {
		db.Close()
		return "", err
	}

	db.Close()
	return refresh_token, nil
}
