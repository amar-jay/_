package service

import (
	"io"
	"sync"
	"time"

	"github.com/amar-jay/comrade/pkg/utils"
	"github.com/livekit/protocol/logger"
	"github.com/pion/webrtc/v3/pkg/media"
)

var OpusSilenceFrame = []byte{0xf8, 0xff, 0xfe}
var OpusSilenceFrameDuration = 20 * time.Millisecond

type provider struct {
	reader        *utils.OggReader
	lastGranule   uint64
	granuleOffset uint64
	queue         []*utils.OggReader
	onComplete    func(err error)
	lock          sync.Mutex
}

func newProvider() *provider {
	return &provider{}
}

func (p *provider) QueueReader(reader *utils.OggReader) {
	p.lock.Lock()
	p.queue = append(p.queue, reader)
	p.lock.Unlock()
}

func (p *provider) NextSample() (media.Sample, error) {
	p.lock.Lock()
	oncomplete := p.onComplete
	if p.reader == nil && len(p.queue) > 0 {
		p.lastGranule = 0
		p.reader = p.queue[0]
		p.queue = p.queue[1:]
	}
	p.lock.Unlock()
	if p.reader == nil {
		return media.Sample{
			Data:     OpusSilenceFrame,
			Duration: OpusSilenceFrameDuration,
		}, nil
	}

	data, err := p.reader.ReadPacket()
	if err != nil {
		if oncomplete != nil {
			oncomplete(err)
		}
		if err == io.EOF {
			p.lock.Lock()
			p.reader = nil
			p.NextSample()
			p.lock.Unlock()
		} else {
			logger.Errorw("failed to read packet", err)
			return media.Sample{}, err
		}
	}

	dur, err := utils.ParsePacketDuration(data)
	if err != nil {
		logger.Errorw("failed to parse packet duration", err)
		return media.Sample{}, err
	}

	return media.Sample{
		Data:     data,
		Duration: time.Duration(dur) * time.Millisecond,
	}, nil
}

// Called when the *one* oggReader finished reading
func (p *provider) OnComplete(f func(err error)) {
	p.lock.Lock()
	p.onComplete = f
	p.lock.Unlock()
}
