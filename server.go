package main

import (
	"cat_adoption_platform/service"

	"cat_adoption_platform/repository"

	"github.com/gin-gonic/gin"
)

type Server struct {
	cs     service.CatService
	engine *gin.Engine
}

func (s *Server) initiateRoute() {
}
func (s *Server) Start() {
	s.initiateRoute()
	s.engine.Run(":2000")
}

func NewServer() *Server {

	csRepo := repository.NewCatRepository(db)
	csService := service.NewcatService(csRepo)
	return &Server{cs: csService, engine: gin.Default()}
}
