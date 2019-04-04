package response

/**
 * @author Oliver Kelton, oakelton@gmail.com
 * @date Mar 21, 2019
 * @dependencies webapp/resp
 * Standardized response formats for data-based (json, xml) endpoints.
 * Extends functions provided by the framework in the "resp" package.
 * Rewrite this package at will to match the design of your
 * RESTful web service.
 */

import (
	"net/http"
	"encoding/json"
	"encoding/xml"
	// framework dependencies
	"gopkg.in/ollykel/webapp.v0.1"
	"gopkg.in/ollykel/webapp.v0.1/resp"
)

/**
 * Standard format for all json and xml responses
 */
type Protocol struct {
	Authorized bool//-- true if user authorized, false otherwise
	Success bool//-- true if request valid and responded to successfully
	Error string//-- if not successful, an explanation of the problem
	Data interface{}//-- if successful, the body of the data
}//-- end Protocol struct

/**
 * Extends the abilities of the resp package's Data struct.
 * The Write method parses the Data field to the type (json or xml)
 * specified by the Type field
 */
type Data resp.Data

/**
 * Utility function for serving data wrapped in a Protocol struct.
 */
func (d *Data) Write (w http.ResponseWriter) {
	proto := Protocol{
		Authorized: true,
		Success: d.Code == http.StatusOK || d.Code == 0,
		Error: d.Msg,
		Data: d.Data}
	d.Data = &proto
	(*resp.Data)(d).Write(w)
}//-- end Data.Write

func HasAll(data webapp.ReqData, keys ...string) bool {
	var exists bool
	for _, k := range keys {
		_, exists = data[k]
		if !exists { return false }
	}
	return true
}//-- end func hasAll

type decoder interface {
	Decode (interface{}) error
}

func ParseBody (dest interface{}, r *http.Request, dataType string) error {
	var dec decoder
	switch (dataType) {
		case "xml", "XML":
			dec = xml.NewDecoder(r.Body)
		default:
			dec = json.NewDecoder(r.Body)
	}//-- end switch
	return dec.Decode(dest)
}//-- end func parseBody

func Error (w http.ResponseWriter, data *Data, code int, msg string) {
	data.Code, data.Msg = code, msg
	data.Write(w)
}//-- end func errorResponse

func Success(w http.ResponseWriter, data *Data) {
	data.Code = http.StatusOK
	data.Write(w)
}//-- end func successResponse

/**
 * Depends on the "Type" global middleware being installed.
 */
func Fmt (output *Data, data webapp.ReqData) {
	output.Type = data["Content-Type"]
}//-- end func fmtResponse


