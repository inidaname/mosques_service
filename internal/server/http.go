package server

import (
	"log"
	"net/http"

	handler "github.com/inidaname/mosque/mosques-service/internal/handler/mosque"
	"github.com/inidaname/mosque/mosques-service/internal/service"
	"github.com/inidaname/mosque/mosques-service/internal/types"
)

type httpServer struct {
	app types.Application
}

func NewHttpServer(app *types.Application) *httpServer {
	return &httpServer{app: *app}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	mosqueService := service.NewMosqueService(&s.app)
	authHandler := handler.NewHttpMosqueService(*mosqueService)
	authHandler.RegisterRouter(router)

	log.Println("Starting server on", s.app.Config.Server.HTTPPort)

	return http.ListenAndServe(s.app.Config.Server.HTTPPort, router)
}
