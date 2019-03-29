package middleware

/**
 * @author Oliver Kelton, oakelton@gmail.com
 * @date Mar 21, 2019
 * Middleware to parse last element in URL into two components: a "tag" and an
 * "ext" (extension).
 * Useful for cases such as determining type of data to serve in GET requests
 * ex: /path/to/file.pdf would be parsed to {"tag": "file","ext": "pdf"}
 */

import (
	"net/http"
	webapp "github.com/ollykel/webapp"
	"log"
)

func Tag (_ http.ResponseWriter, r *http.Request,
		data webapp.ReqData) bool {
	path := r.URL.Path
	getExt := true
	extIndex := len(path)
	for i := extIndex - 1; i != -1; i-- {
		if path[i] == '.' && getExt && i != extIndex - 1 {
			data["ext"] = path[i + 1:]
			getExt = false
			extIndex = i
		}
		if path[i] == '/' {
			if i != len(path) - 1 { data["tag"] = path[i + 1:extIndex] }
			log.Printf("Tag: %s, Ext: %s", data["tag"], data["ext"])
			return true
		}
	}//-- end for i
	return true
}//-- end func Tag

