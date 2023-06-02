package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	//logger middleware for gofiber
)

// Server is the main server struct and contains fiber and other dependencies
type Server struct {
	App *fiber.App
}

func NewServer() *Server {
	conf := fiber.Config{
		Marshal
	}
	App := fiber.New()
	s := &Server{App}
	s.middleware()

	s.App.Static("/", "./assets")
	s.App.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	return s
}

func (s *Server) Start() error {

	s.App.Listen(":3000")
	return nil
}

func (s *Server) Stop() {
	s.App.Shutdown()
}

// middleware
func (s *Server) middleware() {
	s.App.Use(logger.New())
	s.App.Use(recover.New())
	s.App.Use(func(c *fiber.Ctx) error {
		// print url requested
		fmt.Printf("request [from: %s, to: %s]\n", c.IP(), c.OriginalURL())
		return c.Next()
	})

}

// routes
