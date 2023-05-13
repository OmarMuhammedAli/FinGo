package api

import (
	db "github.com/OmarMuhammedAli/FinGo/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Create Server struct with a store and a router
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// Create a function to return a new Server instance
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// Declare routes here
	router.POST("/accounts", server.createAccount)

	server.router = router
	return server
}

// Create a function to start a server

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

// Create a generic error response function to serialize errors
func errorResponse(e error) gin.H {
	return gin.H{"error": e.Error()}
}
