package error

import "errors"

var (
	ErrInvalidConfig     = errors.New("invalid config")
	ErrInvalidURL        = errors.New("invalid url")
	ErrInvalidPort       = errors.New("invalid port")
	ErrInvalidEnv        = errors.New("invalid environment variable")
	ErrCodecNotSupported = errors.New("this codec isn't supported")
	ErrBusy              = errors.New("the gpt participant is already used")
)
