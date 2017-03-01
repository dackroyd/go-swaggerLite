package example

import (
	"fmt"

	"github.com/gocraft/web"
)

// @Title GetStringByInt
// @Description get string by ID
// @Accept  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {object} string
// @Failure 400 {object} APIError "We need ID!!"
// @Failure 404 {object} APIError "Can not find ID"
// @Router /testapi/get-string-by-int/{some_id} [get]
func (c *Context) GetStringByInt(rw web.ResponseWriter, req *web.Request) {
	c.WriteResponse(fmt.Sprint("Some data for %s ID", req.PathParams["some_id"]))
}