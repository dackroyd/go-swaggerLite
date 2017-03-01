package example

import (
	"github.com/gocraft/web"
)

// @Title GetStructByInt
// @Description get struct by ID
// @Accept  json
// @Param   some_id     path    int     true        "Some ID"
// @Param   offset     query    int     true        "Offset"
// @Param   limit      query    int     true        "Offset"
// @Success 200 {object} StructureWithEmbededStructure
// @Failure 400 {object} APIError "We need ID!!"
// @Failure 404 {object} APIError "Can not find ID"
// @Router /testapi/get-struct-by-int/{some_id} [get]
func (c *Context) GetStructByInt(rw web.ResponseWriter, req *web.Request) {
	c.WriteResponse(StructureWithEmbededStructure{})
}
