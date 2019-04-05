package views

import (
	"log"
	"net/http"
	webapp "gopkg.in/ollykel/webapp.v0"
	"response"
	"models/users"
)

func User(w http.ResponseWriter, r *http.Request, data webapp.ReqData) {
	output := response.Data{Type: data["Content-Type"]}
	username := data["tag"]
	if username == "" {
		log.Print("Error! No username provided")
		response.Error(w, &output, http.StatusBadRequest, "no username")
		return
	}
	usr, err := users.GetUserByName(username)
	if err != nil {
		response.Error(w, &output, http.StatusNotFound, err.Error())
		return
	}
	output.Data = usr
	output.Write(w)
}//-- end func User

