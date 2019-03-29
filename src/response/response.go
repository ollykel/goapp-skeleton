package response

/**
 * @author Oliver Kelton, oakelton@gmail.com
 * @date Mar 21, 2019
 * @dependencies webapp/resp
 * Standardized response format for data endpoints.
 */

import (
	"net/http"
	// framework dependencies
	"github.com/ollykel/webapp"
	"github.com/ollykel/webapp/resp"
)

type Protocol struct {
	Authorized bool
	Success bool
	Error string
	Data interface{}
}//-- end Protocol struct

type Data resp.Data

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

type Status struct {
	Success bool
	Error string
}//-- end StatusData struct

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

func Error (w http.ResponseWriter, data *response.Data,
		code int, msg string) {
	data.Code, data.Data = code, &Status{Success: false, Error: msg}
	data.Write(w)
}//-- end func errorResponse

func Success(w http.ResponseWriter, data *response.Data) {
	data.Data = &Status{Success: true}
	data.Write(w)
}//-- end func successResponse

func Fmt (output *response.Data, data webapp.ReqData) {
	output.Type = data["Content-Type"]
}//-- end func fmtResponse


