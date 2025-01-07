package processor

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (p *Processor) FindNote(c *gin.Context, targetName string) {
	if content, ok := p.Context().NoteData[targetName]; ok {
		c.String(http.StatusOK, fmt.Sprintf("Title: %s\nContent:\n\t%s\n", targetName, content))
		return
	}
	c.String(http.StatusNotFound, fmt.Sprintf("[%s] not found in Notebook\n", targetName))
}

func (p *Processor) UpdateNote(c *gin.Context, targetName string, newContent string) {
	p.Context().NoteData[targetName] = newContent

	if content, ok := p.Context().NoteData[targetName]; ok {
		c.String(http.StatusOK, fmt.Sprintf("Title: %s\nContent:\n\t%s\n", targetName, content))
		return
	}
	c.String(http.StatusNotFound, fmt.Sprintf("[%s] not found in Notebook\n", targetName))
}

func (p *Processor) CreateNote(c *gin.Context, targetName string, newContent string) {
	if _, ok := p.Context().NoteData[targetName]; !ok {
		p.Context().NoteData[targetName] = newContent

		content := p.Context().NoteData[targetName]
		c.String(http.StatusOK, fmt.Sprintf("Title: %s\nContent:\n\t%s\n", targetName, content))
		return
	} else if ok {
		c.String(http.StatusForbidden, fmt.Sprintf(
			"[%s] already exist. Please use POST to modify the content.\n",
			targetName),
		)
	}
}

func (p *Processor) NoteWhitespaceAppend(c *gin.Context, targetName string, newContent string) {
	if _, ok := p.Context().NoteData[targetName]; ok {
		var sb strings.Builder
		sb.WriteString(p.Context().NoteData[targetName])
		sb.WriteString(" ")
		sb.WriteString(newContent)

		p.Context().NoteData[targetName] = sb.String()

		content := p.Context().NoteData[targetName]
		c.String(http.StatusOK, fmt.Sprintf("Title: %s\nContent:\n\t%s\n", targetName, content))
		return
	}
	c.String(http.StatusNotFound, fmt.Sprintf("[%s] not found in Notebook\n", targetName))
}
