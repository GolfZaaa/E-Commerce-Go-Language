package servers

import (
	"github.com/GolfZaaa/Golang-ProjectE-commerce/modules/middlewares/middlewaresHandlers"
	"github.com/GolfZaaa/Golang-ProjectE-commerce/modules/middlewares/middlewaresRepositories"
	"github.com/GolfZaaa/Golang-ProjectE-commerce/modules/middlewares/middlewaresUsecases"
	"github.com/GolfZaaa/Golang-ProjectE-commerce/modules/monitor/monitorHandlers"
	"github.com/gofiber/fiber/v2"
)

type IModuleFactory interface {
	MonitorModule()
}

type moduleFactory struct {
	r   fiber.Router
	s   *server
	mid middlewaresHandlers.IMiddlewaresHandler
}

func InitModule(r fiber.Router, s *server, mid middlewaresHandlers.IMiddlewaresHandler) IModuleFactory {
	return &moduleFactory{
		r:   r,
		s:   s,
		mid: mid,
	}
}

func InitMiddlewares(s *server) middlewaresHandlers.IMiddlewaresHandler {
	repository := middlewaresRepositories.MiddlewaresRepository(s.db)
	usecase := middlewaresUsecases.MiddlewareUsecase(repository)
	return middlewaresHandlers.MiddlewaresHandler(s.cfg, usecase)
}

func (m *moduleFactory) MonitorModule() {
	handler := monitorHandlers.MonitorHandler(m.s.cfg)

	m.r.Get("/", handler.HealthCheck)
}
