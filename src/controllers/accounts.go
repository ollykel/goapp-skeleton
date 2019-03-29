package controllers

import (
	"log"
	// framework imports
	"github.com/ollykel/webapp"
	// local imports
	"response"
)

type account struct {
	Username, Password string
}//-- end account struct

func CreateAccount (w http.ResponseWriter, r *http.Request,
		data webapp.ReqData) {
	output := response.Data{}
	response.Fmt(&output, data)
	acct := account{}
	response.ParseBody(&acct, r, data["Content-Type"])
	if acct.Username == "" || acct.Password == "" {
		response.Error(w, &output, http.StatusBadRequest,
			"need username and password")
		return
	}
	username, password := acct.Username, acct.Password
	err := users.Create(username, password)
	if err != nil {
		response.Error(w, &output, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(w, &output)
}//-- end func CreateAccount

