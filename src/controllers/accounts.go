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
	fmtResponse(&output, data)
	acct := account{}
	parseBody(&acct, r, data["Content-Type"])
	if acct.Username == "" || acct.Password == "" {
		errorResponse(w, &output, http.StatusBadRequest,
			"need username and password")
		return
	}
	username, password := acct.Username, acct.Password
	err := users.Create(username, password)
	if err != nil {
		errorResponse(w, &output, http.StatusBadRequest, err.Error())
		return
	}
	successResponse(w, &output)
}//-- end func CreateAccount

