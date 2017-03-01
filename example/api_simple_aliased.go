package example

import (
	"github.com/gocraft/web"
)

// @Title GetSimpleAliased
// @Description get simple aliases
// @Accept  json
// @Success 200 {object} SimpleAlias
// @Failure 400 {object} APIError "We need ID!!"
// @Failure 404 {object} APIError "Can not find ID"
// @Router /testapi/get-simple-aliased [get]
func (c *Context) GetSimpleAliased(rw web.ResponseWriter, req *web.Request) {
	c.WriteResponse("Some string")
}