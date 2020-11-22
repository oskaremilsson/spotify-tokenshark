package dbsetup

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/oskaremilsson/spotify-tokenshark/config"
	"github.com/oskaremilsson/spotify-tokenshark/failure"
)

const createTokensTable = `
CREATE TABLE IF NOT EXISTS tokens (
	username TEXT PRIMARY KEY,
  token bytea NOT NULL
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

const createGdprConsentTable = `
CREATE TABLE IF NOT EXISTS gdpr_consents (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	username TEXT DEFAULT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
`

func Init() {
	db, err := sql.Open("postgres", config.DatabaseUrl)

	_, err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
	failure.Check(err)

	_, err = db.Exec(createTokensTable)
	failure.Check(err)

	_, err = db.Exec(createConsentsTable)
	failure.Check(err)

	_, err = db.Exec(createRequestsTable)
	failure.Check(err)

	_, err = db.Exec(createGdprConsentTable)
	failure.Check(err)

	db.Close()
}
