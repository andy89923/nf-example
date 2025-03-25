package sbi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) getGreetingRoute() []Route {
	return []Route{
		{
			Name:    "Greeting",
			Method:  "GET",
			Pattern: "/",
			APIFunc: s.GetGreeting,
			// Use
			// curl -X GET http://127.0.0.163:8000/greeting/ -w "\n"
		},
		{
			Name:    "Farewell",
			Method:  "POST",
			Pattern: "/",
			APIFunc: s.PostFarewell,
			// Use
			// curl -X POST http://127.0.0.163:8000/greeting/ -w "\n"
		},
		{
			Name:    "Farewell to someone",
			Method:  http.MethodPost,
			Pattern: "/to",
			APIFunc: s.Greetingto,
			// Use
			// curl -X POST http://127.0.0.163:8000/greeting/to -d '{"name": "Alisa"}' -w "\n"
		},
	}
}

func (s *Server) GetGreeting(c *gin.Context) {
	c.String(http.StatusOK, "Greetings!\n")
}

func (s *Server) PostFarewell(c *gin.Context) {
	c.String(http.StatusOK, "Farewells!\n")
}

func (s *Server) Greetingto(c *gin.Context) {
	var names struct {
		Name string `json:"name"`
	}

	err := c.ShouldBindJSON(&names)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	list.Characters = append(list.Characters, names.Name)
	c.String(http.StatusOK, "Farewell ~ "+names.Name+"!")
}