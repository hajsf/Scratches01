package web

import (
	"net/http"
)

func Handlers() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/user-profile", UserProfile)
}
