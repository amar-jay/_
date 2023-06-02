package main

import (
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/amar-jay/comrade/pkg/config"
	"github.com/amar-jay/comrade/pkg/server"
	"github.com/amar-jay/comrade/pkg/service"
	"github.com/livekit/protocol/logger"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "server",
		Usage: "LiveKit server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config",
				Usage:    "config file",
				EnvVars:  []string{"CONFIG"},
				Required: false,
			},
			&cli.StringFlag{
				Name:     "gcp-credentials",
				Usage:    "Path to Google Cloud Credentials File",
				Required: false,
			},

			&cli.StringFlag{
				Name:     "port",
				Usage:    "port to run server on",
				Required: false,
			},
		},
		Action: runServer,
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

// func run(c *cli.Context) error {
// 	logger.Infow("starting server")
// 	configFile := c.String("config")

// 	server, err := service.NewSpeechAI(config, sttClient, ttsClient)
// 	if err != nil {
// 		return err
// 	}

// 	return server.Start()
// }

func runServer(c *cli.Context) error {

	println("starting server")
	// load config
	configFile := c.String("config")

	if configFile == "" {
		logger.Debugw("config file not found, using default config")
		configFile = "config.yml"
	}

	file, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}
	Config, err := config.New([]byte(file))

	if err != nil {
		return err
	}

	// load gcp speech client and tts client
	gcpFile := c.String("gcp-credentials")
	if gcpFile == "" {
		if Config.GcpCredentials == "" {
			return errors.New("gcp credentials file path not found")
		}
		gcpFile = Config.GcpCredentials
	}

	// gcpCred := option.WithCredentialsFile(gcpFile)

	// ctx := context.Background()
	// SttClient, err := stt.NewClient(ctx, gcpCred)
	// if err != nil {
	// 	return err
	// }
	// TtsClient, err := tts.NewClient(ctx, gcpCred)
	// if err != nil {
	// 	return nil
	// }

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	s := server.NewServer()
	speechService, err := service.NewSpeechAI(&service.AIConfig{Config: Config, SstClient: nil, TtsClient: nil})
	speechService.Handle(s.App.Group("/speech"))

	if err != nil {
		return err
	}

	go func() {
		sig := <-sigChan
		logger.Infow("exit requested, shutting down", "signal", sig)
		s.Stop()
	}()

	if err := s.Start(); err != nil {
		return err
	}

	return nil
}
