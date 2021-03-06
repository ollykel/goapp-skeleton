# /src
By convention, src should contain the following packages:
## serve
Contains the main executable for the app. Should consist of a single
file, serve.go, which contains two functions: Methods and main.
Methods returns a map of paths to Method structs, defining which endpoints
should be used for each path. Methods should be implemented as follows:
```go
func Methods () map[string]*goapp.Method {
	return map[string]*goapp.Method{
		"/comments": &goapp.Method{
			GET: views.Comments,
			POST: controllers.CreateComment,
			PUT: controllers.UpdateComment,
			DELETE: controllers.DeleteComment},
		// ...other *goapp.Method declarations
	}
}
```
The main function sets up the app, registering the models and the Methods
provided by func Methods().

As long as you follow the default file structure for /src, you shouldn't
need to edit the default import paths or func main() in serve.go.
 
## models
Packages that interface with the app's chosen database.
Models should consist of several subpackages, each one responsible for a
different model/database table.
Each model subpackage should implement a function "Define" that returns a
goapp/model.Definition struct:
```go
func Define () *model.Definition {
	return &model.Definition{
		Tablename: "model",
		Fields: []model.Field{
			// ...model.Field structs
		},
		Init: Init,
	}
}
```
The main models package should implement a Models func as follows:
```go
func Models () []*model.Definition {
	return []*model.Definition{ foo.Define(), bar.Define() }
}
```
For more information about the implementation of models packages, see
/src/models/README.md.

## controllers
Functions that handle POST, PUT, PATCH, and DELETE requests.
Controllers should be sorted into files based on the path they handle
(i.e. "users.go", "accounts.go"), while all being part of the same
"controllers" package.
All controllers must satisfy the goapp.Controller type definition:
```go
func (w http.ResponseWriter, r *http.Request, data goapp.ReqData)
```
## views
Functions that handle GET requests.
Like controllers, they should be sorted into files based on the path they
handle and satisfy the goapp.View type definition (identical to Controller).

## middleware
Global middleware functions that are called before the execution of each
endpoint (views and controllers).
Middleware functions must satisfy the goapp.Middleware type definition:
```go
func (w http.ResponseWriter, r *http.Request, data goapp.ReqData) bool
```
Middleware functions can pass state to endpoints via data, which is a map of
type map[string]string. They can also abort the execution of endpoint by
writing a response and returning false. A middleware function must return
true to allow an endpoint to execute.

You may choose to create sub-packages for your middleware depending on their
size, or leave them in the main "middleware" package; what's important is that
the middleware package exports a function "Middleware" that returns a slice of middleware
functions:
```go
func Middleware () []goapp.Middleware
```
The order of the middleware functions matters: middleware are executed
sequentially; if one middleware function depends on state provided by the
data parameter, it must come after the middleware that provides that
parameter.

## response
Response should contain wrapper functions for goapp/resp, defining
protocols for serving data and handling errors for use in views and
controllers.

-----
If you choose to implement your app as a SPA (Goapp is optimized to serve a
SPA and a RESTful API from the same server), you should put your client-side
source code in a subdirectory called client. The example client directory
provided here implements a SPA created using React and compiled to a bundle
using Webpack. The tools you use to build your SPA are up to you; all that
matters is that the files necessary to render the SPA client-side are
linked to in index.html and accessible to Goapp's static server from
StaticDir (defined in /config/config.json).

The files to be included in your static directory should be moved to the
app's main directory during creation of the Docker image, as the final image
should not contain the /src directory.

