package sbi

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// Rutas del nuevo servicio /lab6
func (s *Server) getLab6Route() []Route {
    return []Route{
        {
            Name:   "Lab6 Info",
            Method: http.MethodGet,
            Pattern: "/", // -> /lab6/
            APIFunc: func(c *gin.Context) {
                c.JSON(http.StatusOK, gin.H{
                    "message": "Lab 6 OK",
                    "author":  "ValentinoV",
                })
            },
        },
        {
            Name:   "Lab6 Echo",
            Method: http.MethodPost,
            Pattern: "/echo", // -> /lab6/echo
            APIFunc: func(c *gin.Context) {
                var body map[string]interface{}

                if err := c.BindJSON(&body); err != nil {
                    c.JSON(http.StatusBadRequest, gin.H{
                        "error": "invalid JSON body",
                    })
                    return
                }

                c.JSON(http.StatusOK, gin.H{
                    "message":  "Lab 6 echo",
                    "received": body,
                })
            },
        },
    }
}
