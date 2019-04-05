package middleware

/**
 * @author Oliver Kelton, oakelton@gmail.com
 * @date Mar 21, 2019
 * @dependencies models/sessions, response
 * Authentication middleware; determines whether a user is authorized, responds
 * to request if they are not authorized.
 */

import (
	webapp "gopkg.in/ollykel/webapp.v0"
	"gopkg.in/ollykel/webapp.v0/resp"
	"response"
	"models/sessions"
	"net/http"
	"fmt"
	"log"
)

const (
	authHeaderName = sessions.AuthHeaderName
	sessionCookieName = sessions.AuthHeaderName
	AuthStringFmt = sessions.AuthStringFmt
)

var (
	ignoreLogin = map[string]bool {
		"/api/login/": true,
		"/api/account/": true}//-- allows user to login/create account
)

type authorization response.Protocol

func handleShouldLogin(w http.ResponseWriter, data *resp.Data, ans bool) {
	data.Code = http.StatusUnauthorized
	data.Data = &authorization{Authorized: !ans}
	data.Write(w)
}//-- end func handleShouldLogin

func checkAuth (w http.ResponseWriter, res *resp.Data, data webapp.ReqData,
		authString string) bool {
	if authString == "" {
		handleShouldLogin(w, res, true)
		return false
	}
	var username, hash string
	num, err := fmt.Sscanf(authString, AuthStringFmt, &username, &hash)
	if num != 2 || err != nil {
		handleShouldLogin(w, res, true)
		return false
	}
	sess, err := sessions.GetSession(username, hash)
	if sess == nil || err != nil {
		log.Printf("Unauthorized for %s : %s", username, hash)
		if err != nil { log.Print(err.Error()) }
		handleShouldLogin(w, res, true)
		return false
	}
	sess.Extend(sessions.DefaultSessionLength)
	sessions.SetAuth(w, &sessions.Credentials{
		Username: username, Hash: hash})
	data["Username"] = username
	return true
}//-- end func checkAuth

/**
 * 
 */
func CheckLogin (w http.ResponseWriter, r *http.Request,
		data webapp.ReqData) bool {
	if ignoreLogin[r.URL.Path] { return true }
	log.Printf("Checking login at %s...\n", r.URL.Path)
	res := resp.Data{Type: data["Content-Type"]}
	authString := r.Header.Get(authHeaderName)
	if authString != "" { return checkAuth(w, &res, data, authString) }
	cookie, err := r.Cookie(sessionCookieName)
	if err != nil || cookie == nil {
		log.Print("Could not find auth cookie")
		if err != nil { log.Print(err.Error()) }
		handleShouldLogin(w, &res, true)
		return false
	}
	return checkAuth(w, &res, data, cookie.Value)
}//-- end func CheckLogin
