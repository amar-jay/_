package config

import (
	"os"

	"gopkg.in/yaml.v3"

	e "github.com/amar-jay/comrade/pkg/errors"
	"github.com/joho/godotenv"
	"github.com/livekit/protocol/logger"
)

// Config is the main configuration object for the server

type Config struct {
	AppName        string        `yaml:"name"`
	Port           int           `yaml:"port"`
	Logger         logger.Config `yaml:"logger"`
	Production     bool          `yaml:"production"`
	Livekit        LiveKitConfig `yaml:"livekit"`
	GcpCredentials string        `yaml:"gcpCredentialPath"`
}

type LiveKitConfig struct {
	ApiKey      string // from env file
	Secret      string // from env file
	TestToken   string
	Url         string `yaml:"url"`
	BotIdentity string `yaml:"botIdentity"`
	BotLanguage string `yaml:"botLanguage"`
}

var (
	BotIdentity string = "comrade"
	BotLanguage string = "english"
	Production  bool   = false
	AppName     string = "comrade"
)

// Load loads the config from a file and environment variables
func New(content []byte) (*Config, error) {
	conf := Config{}
	err := yaml.Unmarshal(content, &conf)

	if err != nil {
		return nil, err
	}

	if !conf.Production {
		err = godotenv.Load(".env.local")
		if err != nil {
			return nil, err
		}
	} else {
		err = godotenv.Load(".env.production")
		if err != nil {
			return nil, err
		}
	}

	livekitApikey := os.Getenv("LIVEKIT_API_KEY")
	livekitSecret := os.Getenv("LIVEKIT_API_SECRET")

	if livekitApikey == "" {
		return nil, e.ErrInvalidEnv("LIVEKIT_API_KEY")
	}

	if livekitSecret == "" {
		return nil, e.ErrInvalidEnv("LIVEKIT_API_SECRET")
	}

	conf.Livekit.ApiKey = livekitApikey
	conf.Livekit.Secret = livekitSecret
	conf.Livekit.TestToken = os.Getenv("LIVEKIT_TEST_TOKEN")

	BotIdentity = conf.Livekit.BotIdentity
	BotLanguage = conf.Livekit.BotLanguage
	Production = conf.Production
	AppName = conf.AppName

	return &conf, nil
}
