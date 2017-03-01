package example

import (
	"github.com/gocraft/web"
)

// @Title GetSimpleArrayByString
// @Description get simple array by ID
// @Accept  json
// @Param   some_id     path    string     true        "Some ID"
// @Param   offset     query    int     true        "Offset"
// @Param   limit      query    int     true        "Offset"
// @Success 200 {array} string
// @Failure 400 {object} APIError "We need ID!!"
// @Failure 404 {object} APIError "Can not find ID"
// @Router /testapi/get-simple-array-by-string/{some_id} [get]
func (c *Context) GetSimpleArrayByString(rw web.ResponseWriter, req *web.Request) {
	c.WriteResponse([]string{"one", "two", "three"})
}
