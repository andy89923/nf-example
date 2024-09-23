package sbi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
	"strings"
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
			APIFunc: func(c *gin.Context) {
				c.JSON(http.StatusOK, "Welcome to SAO!")
			},
			// Use
			// curl -X GET http://127.0.0.163:8000/sao/ -w "\n"
		},
		{
			Name:    "SAO Character List",
			Method:  http.MethodGet,
			Pattern: "/list",
			APIFunc: func(c *gin.Context) {
				c.JSON(http.StatusOK, "SAO Character List : " + strings.Join(list.Characters, ", "))
			},
			// Use
			// curl -X GET http://127.0.0.163:8000/sao/list -w "\n"
		},
	}
}

func (s *Server) postSAORoute() []Route {
	return []Route{
		{
			Name:    "Create SAO Character",
			Method:  http.MethodPost,
			Pattern: "/character",
			APIFunc: func(c *gin.Context) {
				var character struct {
					Name string `json:"name"`
				}
				
				err := c.ShouldBindJSON(&character)
				
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}
				
				list.Characters = append(list.Characters, character.Name)
				c.JSON(http.StatusOK, "Hello " + character.Name + "!")
			},
			// Use
			// curl -X POST http://127.0.0.163:8000/sao/character -d '{"name": "Kirito"}' -w "\n"
		},
	}
}



