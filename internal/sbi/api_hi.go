package sbi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) getHiRoute() []Route {
	return []Route{
		{
			Name:    "sayHello",
			Method:  http.MethodGet,
			Pattern: "/",
			APIFunc: s.SayHello,
			// Use
			// curl -X GET http://127.0.0.163:8000/hi/ -w "\n"
		},
		{
			Name:    "sayGoodBye",
			Method:  http.MethodPost,
			Pattern: "/",
			APIFunc: s.SayGoodBye,
			// Use
			// curl -X POST http://127.0.0.163:8000/hi/ -w "\n"
		},
	}
}

func (s *Server) SayHello(c *gin.Context) {
	c.String(http.StatusOK, "Hello!\n")
}

func (s *Server) SayGoodBye(c *gin.Context) {
	c.String(http.StatusOK, "Good-bye!\n")
}
