package sbi

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// GET /myservice/hello
func GetHello(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Hello from MyService!",
    })
}

// POST /myservice/data
func PostData(c *gin.Context) {
    var jsonData map[string]interface{}
    if err := c.BindJSON(&jsonData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    // You could add logic here (or call a processor function)
    c.JSON(http.StatusOK, gin.H{
        "received": jsonData,
        "status":   "Data processed successfully",
    })
}
