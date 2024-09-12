package object

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type listOptions struct {
	IDs string `form:"ids" json:"id,omitempty"`
}

func list(c *gin.Context) {
	opts := &listOptions{}

	if err := c.BindQuery(opts); err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{"ids": strings.Split(opts.IDs, ",")})
}

func RegisterAPI(r gin.IRouter) {
	r.GET("objects", list)
}
