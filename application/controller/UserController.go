/**
	Rahmat wahyu hadi
**/

package controller

import (
	"fmt"
	"net/http"

	"../../models"

	"../lib"
)

// GetAllUsersHandler ...
func GetAllUsers(w http.ResponseWriter, req *http.Request) {
	res := lib.Response{ResponseWriter: w}
	user := new(models.User)
	users := user.FetchAll()
	res.SendOK(users)
}

// CreateUserHandler ...
func CreateUser(w http.ResponseWriter, r *http.Request) {
	req := lib.Request{ResponseWriter: w, Request: r}
	res := lib.Response{ResponseWriter: w}

	user := new(models.User)
	req.GetJSONBody(user)
	fmt.Println(req)

	if err := user.Save(); err != nil {
		res.SendBadRequest(err.Error())
		return
	}

	res.SendCreated(user)
}

// GetUserByIdHandler ...
func GetUserById(w http.ResponseWriter, r *http.Request) {
	req := lib.Request{ResponseWriter: w, Request: r}
	res := lib.Response{ResponseWriter: w}

	id, _ := req.GetVarID()
	user := models.User{
		ID: id,
	}

	if err := user.FetchById(); err != nil {
		res.SendNotFound()
		return
	}

	res.SendOK(user)
}

// UpdateUserHandler ...
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	req := lib.Request{ResponseWriter: w, Request: r}
	res := lib.Response{ResponseWriter: w}

	id, _ := req.GetVarID()

	user := new(models.User)
	req.GetJSONBody(user)
	user.ID = id

	if err := user.Save(); err != nil {
		res.SendBadRequest(err.Error())
		return
	}

	res.SendOK(user)
}

// DeleteUserHandler ...
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	req := lib.Request{ResponseWriter: w, Request: r}
	res := lib.Response{ResponseWriter: w}

	id, _ := req.GetVarID()
	user := models.User{
		ID: id,
	}

	if err := user.Delete(); err != nil {
		res.SendNotFound()
		return
	}

	res.SendOK(user)
}
