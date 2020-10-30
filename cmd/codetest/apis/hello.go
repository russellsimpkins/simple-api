package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/russellsimpkins/simple-api/cmd/codetest/types"
	"net/http"
)

// GetHello godoc
// @Summary Executes a simple hello world for demonstration purposes.
// @Produce JSON
// @Success 200 {object} types.World
// @Router /hello/{who} [get]
func GetHello(c *gin.Context) {
	who := c.Param("who")
	w := types.World{Hello: who}
	c.JSON(http.StatusOK, w)
}
