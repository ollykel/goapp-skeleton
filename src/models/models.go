package models

import (
	"gopkg.in/ollykel/webapp.v0.1/model"
	"models/users"
	"models/sessions"
)

func Models () []*model.Definition {
	return []*model.Definition{users.Define(), sessions.Define()}
}//-- end func Definitions

