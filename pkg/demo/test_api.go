package demo

import "github.com/gin-gonic/gin"

func TestAPI(c *gin.Context) (interface{}, error) {
	q := c.Param("q")

	return Test(c, q)
}
