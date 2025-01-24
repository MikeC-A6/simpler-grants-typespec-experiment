package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	api "custom-api/api"
)

// Server implements the api.ServerInterface
type Server struct{}

// Implement required methods from ServerInterface

func (s *Server) CustomOpportunitiesList(c *gin.Context) {
	c.JSON(http.StatusOK, []interface{}{}) // Return empty list for now
}

func (s *Server) OperationsList(c *gin.Context) {
	c.JSON(http.StatusOK, []interface{}{}) // Return empty list for now
}

func (s *Server) OperationsRead(c *gin.Context, id string) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":          id,
		"title":       "Example Opportunity",
		"description": "This is a test opportunity",
		"source":      "http://example.com",
		"metadata": map[string]interface{}{
			"lastModified":        "2024-01-23T00:00:00Z",
			"modifiedBy":         "test-user",
			"status":             "draft",
			"applicationDeadline": "2024-12-31",
			"reviewStatus":       "pending",
			"fundingCategories":  []string{"research", "technology", "education"},
		},
		"fundingDetails": map[string]interface{}{},
		"customFields": map[string]interface{}{
			"fundingOpportunityNumber": map[string]interface{}{
				"name":        "Funding Opportunity Number",
				"type":        "string",
				"value":       "TEST-2024-001",
				"description": "The unique identifier for a given opportunity within this API",
			},
		},
	})
}

func (s *Server) OperationsUpdate(c *gin.Context, id string) {
	var opportunity api.CustomModelsCustomOpportunity
	if err := c.ShouldBindJSON(&opportunity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, opportunity)
}

func (s *Server) CustomOpportunitiesRead(c *gin.Context, id string) {
	s.OperationsRead(c, id) // Reuse the implementation
}

func (s *Server) CustomOpportunitiesUpdate(c *gin.Context, id string) {
	s.OperationsUpdate(c, id) // Reuse the implementation
}

func main() {
	router := gin.Default()

	server := &Server{}
	wrapper := &api.ServerInterfaceWrapper{
		Handler: server,
	}

	// Register all routes
	router.GET("/", wrapper.CustomOpportunitiesList)
	router.GET("/api", wrapper.OperationsList)
	router.GET("/api/:id", wrapper.OperationsRead)
	router.PUT("/api/:id", wrapper.OperationsUpdate)
	router.GET("/:id", wrapper.CustomOpportunitiesRead)
	router.PUT("/:id", wrapper.CustomOpportunitiesUpdate)

	log.Println("Server starting on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
} 