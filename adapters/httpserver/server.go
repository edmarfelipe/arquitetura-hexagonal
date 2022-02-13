package httpserver

import (
	"github.com/edmarfelipe/go-hexagonal/adapters/httpserver/handler"
	"github.com/edmarfelipe/go-hexagonal/application"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func MakeNewServcer() *WebServer {
	return &WebServer{}
}

var (
	ErrParseBodyError = "Could not parse body"
	ErrFailExecute    = "Could not complete the request"
)

func (w WebServer) Serve(addr string) {
	app := fiber.New()

	app.Use(logger.New())

	productHandlers := &handler.ProductHandler{
		Service: handler.CreateProductService(),
	}

	productHandlers.MakeProductHandlers(app)

	app.Listen(addr)
}
