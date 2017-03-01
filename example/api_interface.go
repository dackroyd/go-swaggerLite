package example

import (
	"github.com/gocraft/web"
)

// @Title GetInterface
// @Description get interface
// @Accept  json
// @Success 200 {object} InterfaceType
// @Failure 400 {object} APIError "We need ID!!"
// @Failure 404 {object} APIError "Can not find ID"
// @Router /testapi/get-interface [get]
func (c *Context) GetInterface(rw web.ResponseWriter, req *web.Request) {
	c.WriteResponse(InterfaceType("Some string"))
}