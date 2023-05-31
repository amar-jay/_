package service

import (
	"sync"

	stt "cloud.google.com/go/speech/apiv1"
	tts "cloud.google.com/go/texttospeech/apiv1"
	"github.com/amar-jay/comrade/pkg/config"
	"github.com/livekit/protocol/auth"
	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
	lksdk "github.com/livekit/server-sdk-go"
)

type Participant struct {
	Connecting bool
}

// SpeechAI is a struct that contains the configuration for the speech AI
type SpeechAI struct {
	conf         *config.Config
	sttClient    *stt.Client
	ttsClient    *tts.Client
	participants map[string]*Participant
	// speechChan chan *SpeechRequest
	doneChan chan struct{}
	// closeChan chan struct{}
	auth.KeyProvider
	*lksdk.RoomServiceClient

	lock sync.Mutex
}

type AIConfig struct {
	Config    *config.Config
	SstClient *stt.Client
	TtsClient *tts.Client
}

func NewSpeechAI(c *AIConfig) (*SpeechAI, error) {
	roomService := lksdk.NewRoomServiceClient(c.Config.Livekit.Url, c.Config.Livekit.ApiKey, c.Config.Livekit.Secret)
	return &SpeechAI{
		conf:              c.Config,
		sttClient:         c.SstClient,
		ttsClient:         c.TtsClient,
		RoomServiceClient: roomService,
		KeyProvider:       auth.NewSimpleKeyProvider(c.Config.Livekit.ApiKey, c.Config.Livekit.Secret),
		// speechChan:        make(chan *SpeechRequest),
		doneChan: make(chan struct{}),
	}, nil
}

func (s *SpeechAI) joinRoom(room *livekit.Room) {
	s.lock.Lock()

	if _, ok := s.participants[room.Sid]; ok {
		logger.Infow("bot participant already connected",
			"room", room.Name,
			"participantCount", room.NumParticipants,
		)
		s.lock.Unlock()
		return
	}

	s.participants[room.Sid] = &Participant{
		Connecting: true,
	}
	s.lock.Unlock()

	// Participant := nil
	// s.lock.Lock()
	// s.participants[room.Sid] = &Participant{
	// 	Connecting: false,
	// 	Participant,
	// }
	// s.lock.Unlock()
}
