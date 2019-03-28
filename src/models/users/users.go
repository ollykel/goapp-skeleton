package users

import (
	"fmt"
	"log"
	"crypto/rand"
	"database/sql"
	"github.com/ollykel/webapp/model"
	"models/posts"
)

type User struct {
	Id int
	Username string
	PassHash string `json:"-" xml:"-"`
	PassSalt string `json:"-" xml:"-"`
	NumPosts int
	Posts posts.PostData
}//-- end User struct

func (usr *User) GoString () string {
	return fmt.Sprintf("{Id: %d, Username: %s}", usr.Id, usr.Username)
}//-- end User.String

type UserData []User

func (data *UserData) Append (row model.Scannable) error {
	arr := []User(*data)
	arr = append(arr, User{})
	usr := &arr[len(arr) - 1]
	dummy := sql.NullInt64{}
	err := row.Scan(&usr.Id, &usr.Username, &usr.PassHash, &usr.PassSalt,
		&dummy)
	*data = UserData(arr)
	return err
}//-- end *[]Users.Append

func (data *UserData) Clear () error {
	*data = make(UserData, 0)
	return nil
}

func Tablename () string { return "users" }

func Fields () []model.Field {
	return []model.Field{
		model.Field{Name: "username", Type: model.Varchar, Length: 64,
			Unique: true},
		model.Field{Name: "pass_hash", Type: model.Char, Length: 128},
		model.Field{Name: "pass_salt", Type: model.Char, Length: 16},
		model.Field{Name: "profile_picture", Type: model.BigInt,
			Reference: "images", Null: true, OnDelete: model.SetNull}}
}//-- end Users.Fields

/*
model.Field{Name: "profile_picture", Reference: "images",
			Null: true, OnDelete: model.SetNull}
*/

type PostCount int

func (ct *PostCount) Append (row model.Scannable) error {
	return row.Scan(ct)
}//-- end PostCount.Append

var (
	queryUserExists model.SqlQuery
	createUser model.SqlCmd
	getRecentUsers model.SqlQuery
	getUserByName model.SqlQuery
	verifyCredentials model.SqlQuery
	fetchPosts model.SqlQuery
	fetchPrevPosts model.SqlQuery
	createPost model.SqlQuery
	deletePost model.SqlQuery
	getNumPosts model.SqlQuery
)//-- end vars

func Init (db model.Database) (err error) {
	if db == nil { log.Fatal("passed nil pointer") }
	deletePost, err = db.MakeQuery(`DELETE FROM %TABLE% WHERE user = ?
		AND id = ? LIMIT 1`, posts.Define())
	queryUserExists, err = db.MakeQuery(`SELECT %FIELDS% FROM %TABLE%
		WHERE username = ? LIMIT 1`, Define())
	if err != nil { return }
	createUser, err = db.MakeCmd(`INSERT INTO %TABLE%
		(username, pass_hash, pass_salt) VALUES
		( ? , MD5(CONCAT( ? , ? )), ? )`, Define())
	getRecentUsers, err = db.MakeQuery(`SELECT %FIELDS% FROM %TABLE%
		ORDER BY id DESC LIMIT 10`, Define())
	if err != nil { return }
	getUserByName, err = db.MakeQuery(`SELECT %FIELDS% FROM %TABLE%
		WHERE username = ? LIMIT 1`, Define())
	if err != nil { return }
	verifyCredentials, err = db.MakeQuery(`SELECT %FIELDS% FROM %TABLE%
		WHERE username = ? AND pass_hash = MD5(CONCAT(pass_salt, ? ))`,
		Define())
	if err != nil { return }
	fetchPosts, err = db.MakeQuery(`SELECT %FIELDS% FROM %TABLE%
		WHERE user = ?`, posts.Define())
	if err != nil { return }
	fetchPrevPosts, err = db.MakeQuery(`SELECT %FIELDS% FROM %TABLE%
		WHERE user = ? AND id < ? LIMIT ?`, posts.Define())
	if err != nil { return }
	createPost, err = db.MakeQuery(`INSERT INTO %TABLE% (user, title,
		body) VALUES ( ? , ? , ? )`, posts.Define())
	if err != nil { return }
	getNumPosts, err = db.MakeQuery(`SELECT COUNT(%TABLE%.id)
		FROM %TABLE% INNER JOIN users ON users.id = %TABLE%.user
		WHERE users.username = ?`, posts.Define())
	return
}//-- end func Init

func Define () *model.Definition {
	return &model.Definition{
		Tablename: "users",
		Fields: Fields(),
		Init: Init}
}//-- end func Define

func UserExists(username string) bool {
	data := UserData(make([]User, 0))
	err := queryUserExists(&data, username)
	if err != nil || len(data) == 0 { return false }
	return true
}//-- end Users.UserExists

func generateSalt(length int) string {
	output := make([]byte, length)
	rand.Read(output)
	for i := range output {
		output[i] = (output[i] % ('Z' - 'A')) + byte('A')
	}
	return string(output)
}//-- end func generateSalt

func Create(username, password string) error {
	// args for querier: username, pass_salt, password, pass_salt
	log.Printf("Creating new user %s", username)
	salt := generateSalt(16)
	_, err := createUser(username, salt, password, salt)
	return err
}//-- end Users.Create

func GetRecentUsers() ([]User, error) {
	data := make(UserData, 0)
	err := getRecentUsers(&data)
	return []User(data), err
}//-- end func Users.GetRecentUsers

func GetUserByName (username string) (*User, error) {
	data := make(UserData, 0)
	err := getUserByName(&data, username)
	if len(data) == 0 || data[0].Id == 0 {
		return nil, fmt.Errorf("User %s not found", username)
	}
	usr := &data[0]
	usr.GetNumPosts()
	return usr, err
}//-- end GetUserByName

func VerifyCredentials (username, password string) (bool, error) {
	data := make(UserData, 0)
	err := verifyCredentials(&data, username, password)
	if err != nil {
		log.Print(err.Error())
		return false, err
	}
	if len(data) > 0 { return true, nil }
	return false, nil
}//-- end func VerifyCredentials

func (usr *User) GetNumPosts () error {
	count := (*PostCount)(&usr.NumPosts)
	err := getNumPosts(count, usr.Username)
	return err
}//-- end func User.GetNumPosts

func (usr *User) FetchPosts () error {
	usr.Posts = make(posts.PostData, 0)
	return fetchPosts(&usr.Posts, usr.Id)
}//-- end func User.FetchPosts

func (usr *User) FetchPrevPosts (p *posts.Post, num int) error {
	usr.Posts = make(posts.PostData, 0)
	return fetchPrevPosts(&usr.Posts, usr.Id, p.Id, num)
}//-- end func User.FetchPrevPosts

func (usr *User) CreatePost (p *posts.Post) error {
	return createPost(nil, usr.Id, p.Title, p.Body)
}//-- end func User.CreatePost

func (usr *User) DeletePost (p *posts.Post) error {
	return deletePost(nil, usr.Id, p.Id)
}//-- end func User.DeletePost

