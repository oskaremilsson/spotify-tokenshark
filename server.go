package main

import (
  "fmt"
  "net/http"
  "log"
  "encoding/json"
  "time"
  "github.com/dgrijalva/jwt-go"
  "strings"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)

type HealthCheck struct {
	Time time.Time
}

type Response struct {
  Message string
  Success bool
}

type User struct {
	Username string `json:username`
}

func handler(w http.ResponseWriter, r *http.Request) {
  hk := HealthCheck{time.Now()}

  js, err := json.Marshal(hk)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

  w.WriteHeader(http.StatusOK) // Respond 200
	_, _ = w.Write(js)
}

func StoreRefreshToken(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
	token := r.Form.Get("token")

	tokenParts := strings.Split(token, ".")

	if len(tokenParts) > 2 {
		data, err := jwt.DecodeSegment(tokenParts[1])
    checkErr(err)
    
    user := User{}
    json.Unmarshal([]byte(data), &user)

    if user.Username == "" {
      info := infoJSON("Bad token, missing username", false)
      w.WriteHeader(http.StatusBadRequest)
      _, _ = w.Write(info)
      return
    }
    
    if storeToDb(user.Username, token) {
      info := infoJSON(user.Username + "'s token is stored!", true)
      w.WriteHeader(http.StatusOK)
      _, _ = w.Write(info)
      return
    }

    info := infoJSON("Could not save to database", false)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(info)
	} else {
    info := infoJSON("Bad request", false)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(info)
  }
}

func storeToDb(username string, token string) bool {
  db, err := sql.Open("sqlite3", "./db/data.db")
  checkErr(err)

  stmt, err := db.Prepare("INSERT INTO tokens(username, token) values(?,?)")
  checkErr(err)

  _, err = stmt.Exec(username, token)
  if err != nil {
    stmt, err = db.Prepare("UPDATE tokens SET token = ? WHERE username = ?")
    _, err = stmt.Exec(token, username)
    if err != nil {
      return false
    }
    return true
  }

  return true
}

func main() {
  fmt.Printf("Server is running...\n")
  
  http.HandleFunc("/", handler)
  http.HandleFunc("/storeRefreshToken", StoreRefreshToken)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func infoJSON(message string, success bool) []byte {
  info := Response{Message: message, Success: success}
  js, err := json.Marshal(info)
  checkErr(err)

  return js
}

func checkErr(err error) {
  if err != nil {
      panic(err)
  }
}
