package models

import (
	"github.com/ollykel/webapp/model"
	"models/users"
	"models/sessions"
)

func Definitions () []*model.Definition {
	return []*model.Definition{users.Define(), sessions.Define()}
}//-- end func Definitions

