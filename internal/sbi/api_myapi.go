package sbi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// getMyAPIRoute define las rutas de nuestro nuevo servicio "myapi".
// Tendrá al menos un GET y un POST, como pide el lab.
func (s *Server) getMyAPIRoute() []Route {
	return []Route{
		{
			Name:    "MyAPI GET info",
			Method:  http.MethodGet,
			Pattern: "/",
			APIFunc: func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"service": "myapi",
					"status":  "ok",
				})
			},
		},
		{
			Name:    "MyAPI POST echo",
			Method:  http.MethodPost,
			Pattern: "/echo",
			APIFunc: func(c *gin.Context) {
				var req struct {
					Message string `json:"message" binding:"required"`
				}
				if err := c.ShouldBindJSON(&req); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
					return
				}
				c.JSON(http.StatusOK, gin.H{
					"echo": req.Message,
				})
			},
		},
	}
}
