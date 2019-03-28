/**
 * @author Oliver Kelton, oakelton@gmail.com
 * @date Jan 7, 2019
 * Session data for app
 * @dependencies users.go
 */

package sessions

import (
	"fmt"
	"log"
	"time"
	"crypto/rand"
	"net/http"
	"github.com/ollykel/webapp/model"
	"models/users"
)

const (
	DefaultSessionLength = 300//-- 5 minutes
	DefaultLogoutInterval = 20//-- 20 seconds
	AuthStringFmt = "%s : %s"
	AuthHeaderName = "Authorization"
	hashLength = 16
)

type Credentials struct {
	Username, Password, Hash string
}//-- end Credentials struct

func SetAuth (w http.ResponseWriter, creds *Credentials) {
	authString := fmt.Sprintf(AuthStringFmt, creds.Username, creds.Hash)
	w.Header().Set("Authorization", authString)
	http.SetCookie(w, &http.Cookie{
		Name: AuthHeaderName,
		Value: authString,
		Path: "/api",
		MaxAge: DefaultSessionLength})
}//-- end func SetAuthCookie

func GetAuth (r *http.Request) *Credentials {
	authString := r.Header.Get(AuthHeaderName)
	if authString == "" {
		ck, err := r.Cookie(AuthHeaderName)
		if err != nil || ck == nil { return nil }
		authString = ck.Value
	}
	creds := Credentials{}
	num, err := fmt.Sscanf(authString, AuthStringFmt, &creds.Username,
		&creds.Hash)
	if num != 2 || err != nil { return nil }
	return &creds
}//-- end func GetAuth

var (
	sessionLength int = DefaultSessionLength//-- in milliseconds
)//-- end vars

type Session struct {
	Id int
	Hash string
	Expiration int//-- unix time
}//-- end Session struct

type SessionData []Session

func (sd *SessionData) Append(row model.Scannable) error {
	arr := []Session(*sd)
	arr = append(arr, Session{})
	sess := &arr[len(arr) - 1]
	user := 0
	err := row.Scan(&sess.Id, &user, &sess.Hash, &sess.Expiration)
	*sd = SessionData(arr)
	return err
}//-- end SessionData.Append

func (data *SessionData) Clear () error {
	*data = make(SessionData, 0)
	return nil
}

func Tablename () string { return "sessions" }

func Fields () []model.Field {
	return []model.Field{
		model.Field{Name: "user", Reference: "users"},
		model.Field{Name: "hash", Type: model.Char, Length: hashLength},
		model.Field{Name: "expiration", Type: model.BigInt}}//-- end return
}//-- end Sessions.Fields

func Constraints () map[string]string {
	return map[string]string {
		"FOREIGN KEY (user)": "REFERENCES users (id) ON DELETE CASCADE"}
}//-- end func Constraints

var (
	login model.SqlQuery
	logout model.SqlQuery
	updateSession model.SqlQuery
	getSession model.SqlQuery
	expireSessions model.SqlQuery
)//-- end vars	

func Init (db model.Database) (err error) {
	login, err = db.MakeQuery(`INSERT INTO %TABLE%
		(user, hash, expiration) VALUES ( ? , ? , ? )`, Define())
	if err != nil { return }
	logout, err = db.MakeQuery(`DELETE s.* FROM %TABLE% s INNER JOIN users
		ON s.user = users.id WHERE users.username = ?`, Define())
	if err != nil { return }
	getSession, err = db.MakeQuery(`SELECT %FIELDS% FROM %TABLE% INNER
		JOIN users ON %TABLE%.user = users.id WHERE users.username = ?
		AND %TABLE%.hash = ? LIMIT 1`, Define())
	if err != nil { return }
	expireSessions, err = db.MakeQuery(`DELETE FROM %TABLE% WHERE
		expiration < ?`, Define())
	if err != nil { return }
	updateSession, err = db.MakeQuery(`UPDATE %TABLE% SET hash = ?,
		expiration = ? WHERE id = ? LIMIT 1`, Define())
	if err != nil { return }
	ExpireSessionsInterval(-1)
	return
}//-- end Sessions.Init

func makeHash(length int) string {
	if length < 1 { return "" }
	output := make([]byte, length)
	rand.Read(output)
	for i, ch := range output {
		output[i] = (ch % 26) + 'A'
	}
	return string(output)
}//-- end func makeHash

func Login (username, password string) (string, error) {
	shouldLogin, err := users.VerifyCredentials(username, password)
	if err != nil {
		log.Print(err.Error())
		return "", err
	}
	if !shouldLogin { return "", nil }
	user, err := users.GetUserByName(username)
	if err != nil {
		log.Print(err.Error())
		return "", err
	}
	hash := makeHash(hashLength)
	err = login(nil, user.Id, hash, int(time.Now().Unix()) + sessionLength)
	if err != nil {
		log.Print(err.Error())
		return "", err
	}
	return hash, nil
}//-- end Sessions.Login

func GetSession(username, hash string) (*Session, error) {
	data := make(SessionData, 0)
	err := getSession(&data, username, hash)
	if err != nil { return nil, err }
	if len(data) == 0 {
		return nil, fmt.Errorf("session %s @ %s not found", username, hash)
	}
	return &data[0], nil
}//-- end func GetSession

func CheckLogin (username, hash string) (bool, error) {
	sess, err := GetSession(username, hash)
	return sess != nil, err
}//-- end func CheckLogin

func Logout (username string) error {
	return logout(nil, username)
}//-- end func Logout

func ExpireSessions () error {
	log.Print("Expiring sessions...")
	return expireSessions(nil, time.Now().Unix())
}//-- end func ExpireSessions

func ExpireSessionsInterval (interval int) (halter func()) {
	shouldContinue := true
	interv := interval
	if interv < 1 { interv = DefaultLogoutInterval }
	go func() {
		var err error
		for shouldContinue {
			err = ExpireSessions()
			if err != nil { return }
			time.Sleep(time.Duration(interv) * time.Second)
		}//-- end for shouldContinue
	}()
	halter = func() { shouldContinue = false }
	return
}//-- end func ExpireSessionsInterval

func (sess *Session) Update () error {
	return updateSession(nil, sess.Hash, sess.Expiration, sess.Id)
}//-- end Session.Update

func (sess *Session) Extend (secs int) error {
	sess.Expiration = int(time.Now().Unix()) + secs
	return sess.Update()
}//-- end Session.Extend

func Define () *model.Definition {
	return &model.Definition{
		Tablename: "sessions",
		Fields: Fields(),
		Init: Init}//-- end return
}//-- end func Define

