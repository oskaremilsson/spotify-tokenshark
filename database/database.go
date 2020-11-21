package database

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/oskaremilsson/spotify-tokenshark/config"
	"github.com/oskaremilsson/spotify-tokenshark/failure"
)

func StoreToken(username string, token string) bool {
	db, err := sql.Open("postgres", config.DatabaseUrl)
	failure.Check(err)

	stmt, err := db.Prepare("INSERT INTO tokens (username, token) values ($1, $2)")
	failure.Check(err)

	_, err = stmt.Exec(username, token)
	if err != nil {
		stmt, err = db.Prepare("UPDATE tokens SET token = $1 WHERE username = $2")
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
	db, err := sql.Open("postgres", config.DatabaseUrl)
	failure.Check(err)

	stmt, err := db.Prepare("INSERT INTO consents(username, allow_user) values($1, $2)")
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
	db, err := sql.Open("postgres", config.DatabaseUrl)
	failure.Check(err)

	stmt, err := db.Prepare("DELETE FROM consents WHERE username = $1 AND allow_user = $2")
	failure.Check(err)

	_, err = stmt.Exec(username, disallow_user)
	if err != nil {
		db.Close()
		return false
	}

	db.Close()
	return true
}

func DeleteRequest(username string, requesting string) bool {
	db, err := sql.Open("postgres", config.DatabaseUrl)
	failure.Check(err)

	stmt, err := db.Prepare("DELETE FROM requests WHERE (username = $1 AND requesting = $2) OR (requesting = $3 AND username = $4)")
	failure.Check(err)

	_, err = stmt.Exec(username, requesting, username, requesting)
	if err != nil {
		db.Close()
		return false
	}

	db.Close()
	return true
}

func CreateRequest(username string, requesting string) bool {
	db, err := sql.Open("postgres", config.DatabaseUrl)
	failure.Check(err)

	stmt, err := db.Prepare("INSERT INTO requests(username, requesting) values($1, $2)")
	failure.Check(err)

	_, err = stmt.Exec(username, requesting)
	if err != nil {
		db.Close()
		return false
	}

	db.Close()
	return true
}

func IsServiceUser(username string) bool {
	db, err := sql.Open("postgres", config.DatabaseUrl)
	failure.Check(err)

	sqlStmt := "SELECT username FROM tokens WHERE username = $1"

	err = db.QueryRow(sqlStmt, username).Scan(&username)
	db.Close()

	return err == nil
}

func GetRequests(requesting string) []string {
	db, err := sql.Open("postgres", config.DatabaseUrl)
	failure.Check(err)

	stmt, err := db.Prepare("SELECT * FROM requests WHERE requesting = $1")
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

func GetMyRequests(username string) []string {
	db, err := sql.Open("postgres", config.DatabaseUrl)
	failure.Check(err)

	stmt, err := db.Prepare("SELECT * FROM requests WHERE username = $1")
	failure.Check(err)

	requests := []string{}

	rows, err := stmt.Query(username)
	if err != nil {
		db.Close()
		return requests
	}

	defer rows.Close()
	for rows.Next() {
		var username string
		var requesting string
		err = rows.Scan(&username, &requesting)
		failure.Check(err)
		requests = append(requests, requesting)
	}

	db.Close()
	return requests
}

func GetConsents(username string) []string {
	db, err := sql.Open("postgres", config.DatabaseUrl)
	failure.Check(err)

	stmt, err := db.Prepare("SELECT * FROM consents WHERE allow_user = $1")
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

func GetMyConsents(username string) []string {
	db, err := sql.Open("postgres", config.DatabaseUrl)
	failure.Check(err)

	stmt, err := db.Prepare("SELECT * FROM consents WHERE username = $1")
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
		usernames = append(usernames, allow_user)
	}

	db.Close()
	return usernames
}

func ValidateConsent(me string, username string) bool {
	db, err := sql.Open("postgres", config.DatabaseUrl)
	failure.Check(err)

	sqlStmt := "SELECT * FROM consents WHERE username = $1 AND allow_user = $2"

	err = db.QueryRow(sqlStmt, username, me).Scan(&username, &me)
	db.Close()

	return err == nil
}

func GetRefreshToken(username string) (string, error) {
	db, err := sql.Open("postgres", config.DatabaseUrl)
	failure.Check(err)

	stmt, err := db.Prepare("SELECT token FROM tokens WHERE username = $1")
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

func DeleteMyData(username string) bool {
	db, err := sql.Open("postgres", config.DatabaseUrl)
	failure.Check(err)

	stmt, err := db.Prepare("DELETE FROM requests WHERE username = $1 OR requesting = $1")
	failure.Check(err)

	_, err = stmt.Exec(username)
	if err != nil {
		db.Close()
		return false
	}

	stmt, err = db.Prepare("DELETE FROM consents WHERE username = $1 OR allow_user = $1")
	failure.Check(err)

	_, err = stmt.Exec(username)
	if err != nil {
		db.Close()
		return false
	}

	stmt, err = db.Prepare("DELETE FROM tokens WHERE username = $1")
	failure.Check(err)

	_, err = stmt.Exec(username)
	if err != nil {
		db.Close()
		return false
	}

	db.Close()
	return true
}

func CreateGdprConsent() string {
	var id string
	db, err := sql.Open("postgres", config.DatabaseUrl)
	failure.Check(err)

	stmt, err := db.Prepare("INSERT INTO gdpr_consents DEFAULT VALUES RETURNING id")
	failure.Check(err)

	err = stmt.QueryRow().Scan(&id)
	if err != nil {
		db.Close()
		return ""
	}

	db.Close()
	return id
}
