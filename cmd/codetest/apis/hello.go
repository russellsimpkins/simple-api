package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/russellsimpkins/simple-api/cmd/codetest/types"
	"net/http"
)

// GetHello godoc
// @Summary Executes a simple hello world for demonstration purposes.
// @Produce json
// @Success 200 {object} types.Hello
// @Param who path string true "Who to say hello to."
// @Router /hello/{who} [get]
// @Security ApiKeyAuth
func GetHello(c *gin.Context) {
	who := c.Param("who")
	w := types.Hello{Who: who}
	c.JSON(http.StatusOK, w)
}
