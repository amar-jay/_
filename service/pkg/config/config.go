package config

import (
	"os"

	"gopkg.in/yaml.v3"

	e "github.com/amar-jay/comrade/pkg/error"
	"github.com/joho/godotenv"
	"github.com/livekit/protocol/logger"
)

// Config is the main configuration object for the server

type Config struct {
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
}

var (
	BotIdentity string = "comrade"
)

// Load loads the config from a file and environment variables
func New(content []byte) (*Config, error) {
	conf := Config{}
	err := yaml.Unmarshal(content, &conf)

	if err != nil {
		return nil, err
	}

	if conf.Production {
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
		return nil, e.ErrInvalidEnv
	}

	if livekitSecret == "" {
		return nil, e.ErrInvalidEnv
	}

	conf.Livekit.ApiKey = livekitApikey
	conf.Livekit.Secret = livekitSecret
	conf.Livekit.TestToken = os.Getenv("LIVEKIT_TEST_TOKEN")

	BotIdentity = conf.Livekit.BotIdentity

	return &conf, nil
}
