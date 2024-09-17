package main

import (
	"cat_adoption_platform/config"
	"cat_adoption_platform/controller"
	"cat_adoption_platform/service"
	"database/sql"
	"fmt"
	"log"

	"cat_adoption_platform/repository"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	cr     service.CatResty
	cs     service.CatService
	rs     service.ReviewService
	engine *gin.Engine
}

func (s *Server) initiateRoute() {
	routerGroup := s.engine.Group("/api/v1")
	controller.NewCatController(&s.cs, routerGroup).Route()
	controller.NewCatControllerApi(&s.cr, routerGroup).Route() // Pastikan ini sesuai dengan definisi Route
	controller.NewReviewController(s.rs, routerGroup).Route()
}

func (s *Server) Start() {
	s.initiateRoute()
	s.engine.Run(":2000")
}

func NewServer() *Server {
	// Panggil constructor NewConfig untuk memuat konfigurasi dari file .env
	cfg := config.NewConfig()

	// Inisialisasi koneksi ke database
	db, err := sql.Open(cfg.DBDriver, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName))

	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	csRepo := repository.NewCatRepository(db)
	csService := service.NewCatService(csRepo)

	crService := service.NewRestyService()

	rsRepo := repository.NewReviewRepository(db)
	rsService := service.NewReviewService(rsRepo)

	return &Server{
		cs:     csService,
		cr:     crService,
		rs:     rsService,
		engine: gin.Default(),
	}
}
