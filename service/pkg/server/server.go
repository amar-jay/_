package server

import (
	"fmt"
	"strconv"

	"github.com/amar-jay/comrade/pkg/config"
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	//logger middleware for gofiber
)

// Server is the main server struct and contains fiber and other dependencies
type Server struct {
	App  *fiber.App
	conf *config.Config
}

func NewServer(conf *config.Config) *Server {
	c := fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	}
	App := fiber.New(c)
	s := &Server{App, conf}
	s.middleware()

	s.App.Static("/", "./assets")
	s.App.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	return s
}

func (s *Server) Start() error {

	s.App.Listen(":" + strconv.Itoa(s.conf.Port))
	return nil
}

func (s *Server) Stop() {
	s.App.Shutdown()
}

// middleware
func (s *Server) middleware() {
	if !s.conf.Production {
		s.App.Use(logger.New())
	}
	s.App.Use(cors.New())
	s.App.Use(recover.New())

	// prometheus
	prometheus := fiberprometheus.New(s.conf.AppName)
	prometheus.RegisterAt(s.App, "/metrics")
	s.App.Use(prometheus.Middleware)

}

// routes
