package sbi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) getHelloRoute() []Route {
	return []Route{
		{
			Name:    "GetHello",
			Method:  "GET",
			Pattern: "/",
			APIFunc: s.GetHello,
		},
		{
			Name:    "PostHello",
			Method:  "POST",
			Pattern: "/:message",
			APIFunc: s.PostHello,
		},
	}
}

func (s *Server) GetHello(c *gin.Context) {
	c.String(http.StatusOK, "Hello from feature/hello!")
}

func (s *Server) PostHello(c *gin.Context) {
	message := c.Param("message")
	c.String(http.StatusOK, "Received message: %s", message)
}
