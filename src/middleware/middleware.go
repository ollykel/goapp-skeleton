package middleware

import (
	webapp "github.com/ollykel/webapp"
)

func Middleware() []webapp.Middleware {
	return []webapp.Middleware{Tag, Type, CheckLogin}
}//-- end func Middleware

