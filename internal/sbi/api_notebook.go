package sbi

import (
	"net/http"

	"github.com/andy89923/nf-example/internal/logger"
	"github.com/gin-gonic/gin"
)

func (s *Server) getNotebookRoute() []Route {
	return []Route{
		{
			Name:    "Show Note",
			Method:  http.MethodGet,
			Pattern: "/:Title",
			APIFunc: s.HTTPShowNote,
			// Use
			// curl -X GET http://127.0.0.163:8000/spyfamily/ -w "\n"
		},
		{
			Name:    "Create Note",
			Method:  http.MethodPut,
			Pattern: "/:Title/:Content",
			APIFunc: s.HTTPCreateNote,
			// Use
			// curl -X GET http://127.0.0.163:8000/spyfamily/Anya -w "\n"
			// "Character: Anya Forger"
		},
		{
			Name:    "Update Note",
			Method:  http.MethodPost,
			Pattern: "/:Title/:Content",
			APIFunc: s.HTTPUpdateNote,
			// Use
			// curl -X GET http://127.0.0.163:8000/spyfamily/Anya -w "\n"
			// "Character: Anya Forger"
		},
		{
			Name:    "Note Append with whitespace at front.",
			Method:  http.MethodPost,
			Pattern: "/:Title/append/:Content_append",
			APIFunc: s.HTTPNoteWhitespaceAppend,
			// Use
			// curl -X GET http://127.0.0.163:8000/spyfamily/Anya -w "\n"
			// "Character: Anya Forger"
		},
	}
}

func (s *Server) HTTPShowNote(c *gin.Context) {
	logger.SBILog.Infof("In HTTPShowNote")

	targetName := c.Param("Title")
	if targetName == "" {
		c.String(http.StatusBadRequest, "No name provided")
		return
	}

	s.Processor().FindNote(c, targetName)
}

func (s *Server) HTTPUpdateNote(c *gin.Context) {
	logger.SBILog.Infof("In HTTPUpdateNote")

	targetName := c.Param("Title")
	if targetName == "" {
		c.String(http.StatusBadRequest, "No name provided")
		return
	}

	newContent := c.Param("Content")

	s.Processor().UpdateNote(c, targetName, newContent)
}

func (s *Server) HTTPCreateNote(c *gin.Context) {
	logger.SBILog.Infof("In HTTPCreateNote")

	targetName := c.Param("Title")
	if targetName == "" {
		c.String(http.StatusBadRequest, "No name provided")
		return
	}

	newContent := c.Param("Content")

	s.Processor().CreateNote(c, targetName, newContent)
}

func (s *Server) HTTPNoteWhitespaceAppend(c *gin.Context) {
	logger.SBILog.Infof("In HTTPNoteAppend")

	targetName := c.Param("Title")
	if targetName == "" {
		c.String(http.StatusBadRequest, "No name provided")
		return
	}

	newContent := c.Param("Content_append")

	s.Processor().NoteWhitespaceAppend(c, targetName, newContent)
}
