package sbi

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type SAOCharacterList struct {
	Characters []string
}

var list = &SAOCharacterList{
	Characters: []string{},
}

func (s *Server) getSAORoute() []Route {
	return []Route{
		{
			Name:    "Welcome to SAO!",
			Method:  http.MethodGet,
			Pattern: "/",
			APIFunc: s.WelcomeSAO,
			// Use
			// curl -X GET http://127.0.0.163:8000/sao/ -w "\n"
		},
		{
			Name:    "SAO Character List",
			Method:  http.MethodGet,
			Pattern: "/list",
			APIFunc: s.ListSAOCharacter,
			// Use
			// curl -X GET http://127.0.0.163:8000/sao/list -w "\n"
		},
		{
			Name:    "Create SAO Character",
			Method:  http.MethodPost,
			Pattern: "/character",
			APIFunc: s.CreateSAOCharacter,
			// Use
			// curl -X POST http://127.0.0.163:8000/sao/character -d '{"name": "Kirito"}' -w "\n"
		},
	}
}

func (s *Server) WelcomeSAO(c *gin.Context){
	c.String(http.StatusOK, "Welcome to SAO!")
}

func (s *Server) ListSAOCharacter(c *gin.Context){
	c.String(http.StatusOK, "SAO Character List : "+strings.Join(list.Characters, ", "))
}

func (s *Server) CreateSAOCharacter(c *gin.Context){
	var character struct {
		Name string `json:"name"`
	}

	err := c.ShouldBindJSON(&character)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	list.Characters = append(list.Characters, character.Name)
	c.String(http.StatusOK, "Hello "+character.Name+"!")
}
