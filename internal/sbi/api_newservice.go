package sbi

import (
	"net/http"

	"github.com/andy89923/nf-example/internal/logger"
	"github.com/gin-gonic/gin"
)

func (s *Server) getNewServiceRoute() []Route {
	return []Route{
		{
			Name:    "Get New Service Info",
			Method:  http.MethodGet,
			Pattern: "/info",
			APIFunc: s.HTTPGetNewServiceInfo,
			// Use
			// curl -X GET http://127.0.0.163:8000/newservice/info -w "\n"
		},
		{
			Name:    "Post New Service Data",
			Method:  http.MethodPost,
			Pattern: "/data",
			APIFunc: s.HTTPPostNewServiceData,
			// Use
			// curl -X POST -H "Content-Type: application/json" -d '{"name":"Test User"}' http://127.0.0.163:8000/newservice/data -w "\n"
		},
	}
}

func (s *Server) HTTPGetNewServiceInfo(c *gin.Context) {
	logger.SBILog.Infof("In HTTPGetNewServiceInfo")
	c.JSON(http.StatusOK, gin.H{
		"message": "This is the GET method of the new service",
		"service": "NewService",
	})
}

func (s *Server) HTTPPostNewServiceData(c *gin.Context) {
	logger.SBILog.Infof("In HTTPPostNewServiceData")
	var data struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Data received",
		"name":    data.Name,
	})
}
