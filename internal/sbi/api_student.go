package sbi

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// Estructura para el cuerpo del POST
type StudentRequest struct {
    Name   string `json:"name"`
    CodeID string `json:"code_id"`
}

// Rutas del grupo /student
func (s *Server) getStudentRoute() []Route {
    return []Route{
        {
            Name:    "GetStudent",
            Method:  "GET",
            Pattern: "/",
            APIFunc: GetStudent,
        },
        {
            Name:    "PostStudent",
            Method:  "POST",
            Pattern: "/",
            APIFunc: PostStudent,
        },
    }
}

// Handler para GET /student/
func GetStudent(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Hello from Student API!",
    })
}

// Handler para POST /student/
func PostStudent(c *gin.Context) {
    var req StudentRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Student registered successfully",
        "name":    req.Name,
        "code_id": req.CodeID,
    })
}
