package example

import (
	"github.com/gocraft/web"

	"github.com/RobotsAndPencils/go-swaggerLite/example/subpackage"
)

// @Title GetStructArrayByString
// @Description get struct array by ID
// @Accept  json
// @Param   some_id     path    string     true        "Some ID"
// @Param   offset     query    int     true        "Offset"
// @Param   limit      query    int     true        "Offset"
// @Success 200 {array} SimpleStructureWithAnnotations
// @Failure 400 {object} APIError "We need ID!!"
// @Failure 404 {object} APIError "Can not find ID"
// @Router /testapi/get-struct-array-by-string/{some_id} [get]
func (c *Context) GetStructArrayByString(rw web.ResponseWriter, req *web.Request) {
	c.WriteResponse([]subpackage.SimpleStructure{
		subpackage.SimpleStructure{},
		subpackage.SimpleStructure{},
		subpackage.SimpleStructure{},
	})
}
