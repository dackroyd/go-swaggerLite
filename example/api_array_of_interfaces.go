package example

import (
	"github.com/gocraft/web"
)

// @Title GetArrayOfInterfaces
// @Description get array of interfaces
// @Accept  json
// @Success 200 {array} InterfaceType
// @Failure 400 {object} APIError "We need ID!!"
// @Failure 404 {object} APIError "Can not find ID"
// @Router /testapi/get-array-of-interfaces [get]
func (c *Context) GetArrayOfInterfaces(rw web.ResponseWriter, req *web.Request) {
	c.WriteResponse([]InterfaceType{"Some string", 123, "10"})
}
