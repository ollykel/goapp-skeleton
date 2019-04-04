package middleware

import (
	webapp "gopkg.in/ollykel/webapp.v0.1"
)

func Middleware() []webapp.Middleware {
	return []webapp.Middleware{Tag, Type, CheckLogin}
}//-- end func Middleware

