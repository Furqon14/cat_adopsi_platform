package main

import (
	"cat_adoption_platform/config"
	"cat_adoption_platform/controller"
	"cat_adoption_platform/middleware"
	"cat_adoption_platform/repository"
	"cat_adoption_platform/service"
	"cat_adoption_platform/uploaders"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	uis    service.UploadImageService
	cr     service.CatResty
	cs     service.CatService
	rs     service.ReviewService
	mts    service.MidtransService
	js     service.JwtService
	am     middleware.AuthMiddleware
	engine *gin.Engine
}

func (s *Server) initiateRoute() {
	routerGroup := s.engine.Group("/api/v1")
	routerGroup.Static("/images", "./images")
	controller.NewCatController(&s.cs, routerGroup).Route()
	controller.NewCatControllerApi(&s.cr, routerGroup).Route() // Pastikan ini sesuai dengan definisi Route
	controller.NewReviewController(s.rs, routerGroup).Route()
	controller.NewUploadImageController(s.uis, routerGroup).Route()
	controller.NewMidtransController(s.mts, routerGroup).Route() // Pastikan ini sesuai dengan definisi Route
}

func (s *Server) Start() {
	s.initiateRoute()
	s.engine.Run(":2000")
}

func NewServer() *Server {
	// Panggil constructor NewConfig untuk memuat konfigurasi dari file .env
	cfg, _ := config.NewConfig()

	// Inisialisasi koneksi ke database
	db, err := sql.Open(cfg.DBDriver, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName))

	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	uiRepo := repository.NewUploadImageRepository(db)
	imageUploader := uploaders.NewImageUploader("./images")
	uiService := service.NewUploadImageService(uiRepo, imageUploader)

	csRepo := repository.NewCatRepository(db)
	csService := service.NewCatService(csRepo)

	crService := service.NewRestyService()

	rsRepo := repository.NewReviewRepository(db)
	rsService := service.NewReviewService(rsRepo)

	// Inisialisasi Midtrans client
	midtransRepo := repository.NewMidtransRepository(db)
	midtransService := service.NewMidtransService(midtransRepo)

	jwtService := service.NewJwtService(cfg.JwtConfig)
	authMiddleware := middleware.NewAuthMiddleware(jwtService)

	return &Server{
		uis:    uiService,
		cs:     csService,
		cr:     crService,
		rs:     rsService,
		mts:    midtransService,
		js:     jwtService,
		am:     authMiddleware,
		engine: gin.Default(),
	}
}
