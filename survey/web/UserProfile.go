package web

import (
	"fmt"
	"net/http"
)

const userProfile = "user_profile.html"

func UserProfile(w http.ResponseWriter, r *http.Request) {
	t, ok := templates[userProfile]
	if !ok {
		// TODO handle error
		return
	}

	//	status := <-global.LogginStatus
	data := make(map[string]interface{})
	//data["status"] = status
	data["Name"] = "John Doe"
	data["Email"] = "johndoe@email.com"
	data["Address"] = "Fake Street, 123"
	data["PhoneNumber"] = "654123987"

	if err := t.Execute(w, data); err != nil {
		// TODO handle error
		fmt.Println(err)
	}
}
