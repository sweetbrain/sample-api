package main

import (
	"net/http"
	"regexp"
)

var vaidUUID = regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")

func UserRegister(w http.ResponseWriter, r *http.Request) {
	if err := checkHandler(r); err != nil {
		newHTTPError(w, err)
		return
	}

	newUser := User{}
	err := requestBodyToStruct(r, &newUser)
}