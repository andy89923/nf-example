package processor

import (
	"net/http"

	"github.com/andy89923/nf-example/internal/context"
	"github.com/gin-gonic/gin"
)

func (p *Processor) CreateNewTask(c *gin.Context) {
	var newTask context.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	tasks := p.Context().Tasks
	newTask.ID = len(tasks) + 1
	p.Context().Tasks = append(tasks, newTask)

	c.JSON(http.StatusCreated, newTask)
}

func (p *Processor) GetAllTasks(c *gin.Context) {
	tasks := p.Context().Tasks
	c.JSON(http.StatusOK, tasks)
}
