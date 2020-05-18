package StoreRefreshToken

import (
  "fmt"
  "github.com/dgrijalva/jwt-go"
  "net/http"
  "strings"
)

func run(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	token := r.Form.Get("token")

	tokenParts := strings.Split(token, ".")

	if len(tokenParts) > 2 {
		data, err := jwt.DecodeSegment(tokenParts[1])
		if err != nil {
		fmt.Println("error:", err)
		return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(data)

		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Bad request"))

		return
	}
}
