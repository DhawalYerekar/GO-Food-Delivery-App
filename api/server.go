package api

import (
	db "github.com/DhawalYerekar/Users-App/DB/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) (*Server, error) {
	server := &Server{store: store}

	server.setupRouter()
	return server, nil

}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)

	server.router = router

}
