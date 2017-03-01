package example

import (
	"github.com/gocraft/web"
)

// @Title GetStruct3
// @Description get struct3
// @Accept  json
// @Success 200 {object} StructureWithSlice
// @Failure 400 {object} APIError "We need ID!!"
// @Failure 404 {object} APIError "Can not find ID"
// @Router /testapi/get-struct3 [get]
func (c *Context) GetStruct3(rw web.ResponseWriter, req *web.Request) {
	c.WriteResponse(StructureWithSlice{})
}
