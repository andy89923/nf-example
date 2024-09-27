package sbi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) getStu113Route() []Route {
	return []Route{
		{
			Name:    "Hello Student 113",
			Method:  http.MethodGet,
			Pattern: "/",
			APIFunc: func(c *gin.Context) {
				c.JSON(http.StatusOK, "Hello 2024!")
			},
			// Use
			// curl -X GET http://127.0.0.163:8000/stu113/ -w "\n"
		},
	}
}
