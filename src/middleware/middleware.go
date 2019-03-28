package middleware

import (
	webapp "github.com/ollykel/webapp"
)

func Middleware() []webapp.Middleware {
	return []webapp.Middleware{GetBodyData, CheckLogin}
}//-- end func Middleware
