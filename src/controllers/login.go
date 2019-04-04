package controllers

import (
	"log"
	"net/http"
	"models/sessions"
	webapp "gopkg.in/ollykel/webapp.v0.1"
	resp "gopkg.in/ollykel/webapp.v0.1/resp"
	"response"
)

type LoginStatus response.Protocol

type authData struct {
	Authorization string
}//-- end type authData

func Login(w http.ResponseWriter, r *http.Request, data webapp.ReqData) {
	output := resp.Data{Type: data["Content-Type"]}
	creds := sessions.Credentials{}
	response.ParseBody(&creds, r, data["Content-Type"])
	var err error
	creds.Hash, err = sessions.Login(creds.Username, creds.Password)
	if err != nil || creds.Hash == "" {
		log.Printf("Login failed for %s", creds.Username)
		if err != nil { log.Print(err.Error()) }
		output.Code = http.StatusBadRequest
		output.Data = &LoginStatus{
			Authorized: false, Success: false, Error: "login failed"}
		output.Write(w)
		return
	}
	output.Data = &LoginStatus{Authorized: true, Success: true}
	sessions.SetAuth(w, &creds)
	log.Printf("Login success for %s", creds.Username)
	output.Write(w)
}//-- end func Login

func Logout (w http.ResponseWriter, r *http.Request, data webapp.ReqData) {
	output := response.Data{}
	response.Fmt(&output, data)
	username := data["Username"]
	err := sessions.Logout(username)
	if err != nil {
		log.Print(err.Error());
		response.Error(w, &output, http.StatusInternalServerError,
			err.Error())
		return
	}
	response.Success(w, &output)
}//-- end func Logout

