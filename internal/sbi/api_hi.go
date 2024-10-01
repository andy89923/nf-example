package sbi

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func (s *Server) getHiRoute() []Route {
	return []Route {
		{
			Name: "Hi",
			Method: http.MethodGet,
			Pattern: "/",
			APIFunc: func(c *gin.Context) {
				c.JSON(http.StatusOK, "Today is a good day.")
			},
			// Use
			// curl -X GET http://127.0.0.163:8000/hi/ -w "\n"
		},
	}
}
