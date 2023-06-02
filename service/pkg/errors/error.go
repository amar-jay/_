package errors

import "errors"

var (
	ErrInvalidConfig     = errors.New("invalid config")
	ErrInvalidURL        = errors.New("invalid url")
	ErrInvalidPort       = errors.New("invalid port")
	ErrCodecNotSupported = errors.New("this codec isn't supported")
	ErrBusy              = errors.New("the gpt participant is already used")
	ErrInvalidFormat     = errors.New("invalid format")
	ErrMuted             = errors.New("the gpt participant is muted")
)

func ErrInvalidEnv(env string) error {
	return errors.New("invalid env variable: " + env)
}
