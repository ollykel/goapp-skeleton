package views

import (
	"net/http"
	"models/sessions"
	"models/users"
	webapp "github.com/ollykel/webapp"
	"github.com/ollykel/webapp/resp"
	"response"
)

func Account (w http.ResponseWriter, r *http.Request,
		data webapp.ReqData) {
	output := resp.Data{Type: data["Content-Type"]}
	protocol := response.Protocol{}
	output.Data = &protocol
	creds := sessions.GetAuth(r)
	if creds == nil {
		protocol.Error = "unauthorized"
		output.Code = http.StatusUnauthorized
		output.Write(w)
		return
	}
	protocol.Authorized = true
	usr, err := users.GetUserByName(creds.Username)
	if err != nil || usr == nil {
		protocol.Error = "account not found"
		output.Code = http.StatusInternalServerError
		output.Write(w)
		return
	}
	usr.FetchPosts()
	protocol.Success, protocol.Data = true, usr
	output.Write(w)
}//-- end func Account

