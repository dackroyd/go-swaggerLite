// @SubApi Order management API [/orders]
package example

import (
	"encoding/json"

	"github.com/gocraft/web"
)

type Context struct {
	response interface{}
}

func (c *Context) WriteResponse(response interface{}) {
	c.response = response
}

func InitRouter() *web.Router {
	router := web.New(Context{}).
		Middleware(web.LoggerMiddleware).
		Middleware(web.ShowErrorsMiddleware).
		Middleware(func(c *Context, rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
		resultJSON, _ := json.Marshal(c.response)
		rw.Write(resultJSON)
	}).
		Get("/get-string-by-int/{some_id}", (*Context).GetStringByInt).
		Get("/testapi/get-struct-by-int/{some_id}", (*Context).GetStructByInt).
		Get("/testapi/get-simple-array-by-string/{some_id}", (*Context).GetSimpleArrayByString).
		Get("/testapi/get-struct-array-by-string/{some_id}", (*Context).GetStructArrayByString).
		Get("/testapi/get-interface", (*Context).GetInterface).
		Get("/testapi/get-simple-aliased", (*Context).GetSimpleAliased).
		Get("/testapi/get-array-of-interfaces", (*Context).GetArrayOfInterfaces).
		Get("/testapi/get-struct3", (*Context).GetStruct3).
		Get("/testapi/get-struct2-by-int/{some_id}", (*Context).GetStruct2ByInt)

	return router
}
