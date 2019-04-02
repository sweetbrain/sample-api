package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"../common"
	"../model"
)

var validUUID = regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")

func UserRegister(w http.ResponseWriter, r *http.Request) {
	if err := checkHeader(r); err != nil {
		newHTTPError(w, err)
		return
	}

	newUser := model.User{}
	err := requestBodyToStruct(r, &newUser)
	if err != nil {
		newHTTPError(w, err)
		return
	}
	user, err := model.RegistUser(newUser)

	if err != nil {
		newHTTPError(w, err)
		return
	}

	newHTTPResponse(w, http.StatusOK, user)
}

func UserUpdater(w http.ResponseWriter, r *http.Request) {
	if err := checkHeader(r); err != nil {
		newHTTPError(w, err)
		return
	}

	q := r.URL.Query()
	id := q.Get("id")

	newUser := model.User{}
	err := requestBodyToStruct(r, &newUser)
	newUser.ID = id

	if err != nil {
		newHTTPError(w, err)
		return
	}

	newUser, err = model.UpdateUser(newUser)

	if err != nil {
		newHTTPError(w, err)
		return
	}

	newHTTPResponse(w, http.StatusOK, newUser)
}

func UserReader(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	id := q.Get("id")

	if id != "" {
		user, err := model.ReadUser(id)

		if err != nil {
			newHTTPError(w, err)
			return
		}

		newHTTPResponse(w, http.StatusOK, user)
		return
	}

	users, err := model.ListUser()
	if err != nil {
		 newHTTPError(w, err)
		 return
	}

	newHTTPResponse(w, http.StatusOK, users)
}

func UserDeleter(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	id := q.Get("id")

	err := model.DeleteUser(id)
	if err != nil {
		newHTTPError(w, err)
		return
	}

	successMessage := struct {
		code 	int 	`json:"code"`
		Message string	`json:"message"`
	}{
		http.StatusOK,
		fmt.Sprintf("Success delete user=[%s]", id),
	}

	newHTTPResponse(w, http.StatusOK, successMessage)
}


func NotFoundResources(w http.ResponseWriter, r * http.Request) {
	newHTTPError(w, common.NewError(http.StatusNotFound, "Not found resources"))
}


func newHTTPError(w http.ResponseWriter, err error) {
	errMessage := err.(common.ErrorMessage)
	w.WriteHeader(errMessage.Code)
	w.Write(structToResponseBody(err))
	return
}

func newHTTPResponse(w http.ResponseWriter, code int, body interface{}) {
	w.WriteHeader(code)
	w.Write(structToResponseBody(body))
	return
}

func getUserID(r *http.Request) (string, error) {
	query := r.URL.Query()
	id := query.Get("id")

	if id != "" {
		if !validUUID.MatchString(id) {
			return id, common.NewError(http.StatusBadRequest, "invalid user id format")
		}
	}

	return id, nil
}


func structToResponseBody(data interface{}) []byte {
	json, err := json.Marshal(&data)

	if err != nil {
		return []byte(err.Error())
	}

	return json
}

func requestBodyToStruct(r *http.Request, data interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		return common.NewError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func checkHeader(r *http.Request) error {
	if r.Method == http.MethodPost || r.Method == http.MethodPut {
		code := http.StatusBadRequest
		if r.Header.Get("Content-type") != "application/json" {
			return common.NewError(code, "Content-type is not application/json")
			}

		if r.ContentLength == 0 {
			return common.NewError(code, "Request body length is 0")
		}
	}
	return nil
}