package middleware

/**
 * @author Oliver Kelton, oakelton@gmail.com
 * @date Mar 21, 2019
 * Determines content type to return for a given request. Currently supports
 * JSON and XML 
 */

import (
	"net/http"
	webapp "github.com/ollykel/webapp"
)

func fromExtension (r *http.Request, data webapp.ReqData) bool {
	switch (data["ext"]) {
		case "xml":
			data["Content-Type"] = "xml"
		default:
			data["Content-Type"] = "json"
	}//-- end switch
	return true
}//-- end func fromExtension

func Type (w http.ResponseWriter, r *http.Request,
		data webapp.ReqData) bool {
	if r.Method == "get" || r.Method == "GET" {
		return fromExtension(r, data)
	}
	switch (r.Header.Get("Content-Type")) {
		case "application/xml", "text/xml":
			data["Content-Type"] = "xml"
		default:
			data["Content-Type"] = "json"
	}//-- end switch
	return true
}//-- end func GetDataType

