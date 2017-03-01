package example

import (
	"github.com/gocraft/web"
)

// @Title GetStruct2ByInt
// @Description get struct2 by ID
// @Accept  json
// @Param   some_id     path    int     true        "Some ID"
// @Param   offset     query    int     true        "Offset"
// @Param   limit      query    int     true        "Offset"
// @Success 200 {object} StructureWithEmbededPointer
// @Failure 400 {object} APIError "We need ID!!"
// @Failure 404 {object} APIError "Can not find ID"
// @Router /testapi/get-struct2-by-int/{some_id} [get]
func (c *Context) GetStruct2ByInt(rw web.ResponseWriter, req *web.Request) {
	c.WriteResponse(StructureWithEmbededPointer{})
}
