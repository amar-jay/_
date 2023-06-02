package service

import (
	"io"

	errors "github.com/amar-jay/comrade/pkg/errors"
	"github.com/amar-jay/comrade/pkg/utils"
	lksdk "github.com/livekit/server-sdk-go"
	"github.com/pion/webrtc/v3"
)

type aITrack struct {
	provider   *provider
	track      *lksdk.LocalSampleTrack
	doneChan   chan struct{}
	closedChan chan struct{}
}

type AITrack interface {
	OnComplete(func(error))                                                  // runs when the reader is done
	Reader(io.Reader) error                                                  // reads from the reader
	Publish(p *lksdk.LocalParticipant) (*lksdk.LocalTrackPublication, error) // publishes the track to the participant
}

func NewAITrack() (aitrack AITrack, err error) {
	capability := webrtc.RTPCodecCapability{
		MimeType:  webrtc.MimeTypeOpus,
		ClockRate: 48000,
		Channels:  1,
	}

	track, err := lksdk.NewLocalSampleTrack(capability)
	if err != nil {
		return nil, err
	}
	return &aITrack{
		track:      track,
		provider:   newProvider(),
		doneChan:   make(chan struct{}),
		closedChan: make(chan struct{}),
	}, nil
}

func (a *aITrack) Reader(r io.Reader) error {
	oggReader, oggHeader, err := utils.NewOggReader(r)
	if err != nil {
		return err
	}

	// oggHeader.SampleRate is _not_ the sample rate to use for playback.
	// see https://www.rfc-editor.org/rfc/rfc7845.html#section-3
	if oggHeader.Channels != 1 /*|| oggHeader.SampleRate != 48000*/ {
		return errors.ErrInvalidFormat
	}

	a.provider.QueueReader(oggReader)
	return nil
}

func (a *aITrack) Publish(p *lksdk.LocalParticipant) (*lksdk.LocalTrackPublication, error) {
	return p.PublishTrack(a.track, &lksdk.TrackPublicationOptions{})
}

func (a *aITrack) OnComplete(f func(error)) {
	a.provider.OnComplete(f)
	a.doneChan <- struct{}{}
}
