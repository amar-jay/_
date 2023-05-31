package service

import (
	"encoding/json"
	"errors"

	"github.com/amar-jay/comrade/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
	"github.com/livekit/protocol/webhook"
)

func (s *SpeechAI) Handle(c fiber.Router) error {
	c.Get("/join", s.joinHandler)
	c.Get("/leave", s.leaveHandler)
	c.Get("/health", s.healthHandler)
	c.Post("/webhook", s.webHookHandler)
	c.Post("/speech", s.speechHandler)
	return nil
}

func (s *SpeechAI) joinHandler(c *fiber.Ctx) error {
	// get room name from query params
	roomName := c.Query("room")
	if roomName == "" {
		return errors.New("room name not provided")
	}

	// get list of rooms by querying livekit server
	rooms, err := s.RoomServiceClient.ListRooms(c.Context(), &livekit.ListRoomsRequest{
		Names: []string{roomName},
	})

	if err != nil {
		return err
	}

	if len(rooms.Rooms) == 0 {
		return errors.New("room does not exist")
	}

	s.joinRoom(rooms.Rooms[0])
	res, err := json.Marshal(map[string]string{
		"status": "Success",
	})

	if err != nil {
		return err
	}
	c.Send(res)

	return nil
}

func (s *SpeechAI) leaveHandler(c *fiber.Ctx) error {

	return nil
}

func (s *SpeechAI) webHookHandler(c *fiber.Ctx) error {

	// convert fasthttp request using adapter
	req, err := adaptor.ConvertRequest(c, true)
	if err != nil {
		return err
	}

	// get http request from gin and pass it to webhook
	event, err := webhook.ReceiveWebhookEvent(req, s.KeyProvider)

	if err != nil {
		logger.Errorw("error receiving webhook event", err)
		return err
	}
	// when participant joins
	if event.Event == webhook.EventParticipantJoined {
		if event.Participant.Identity == config.BotIdentity {
			return nil
		}
		logger.Infow("participant connected: ", "participant", event.Participant.Identity)
		s.joinRoom(event.Room)
	}

	return nil
}

func (s *SpeechAI) speechHandler(c *fiber.Ctx) error {

	return nil
}

func (s *SpeechAI) healthHandler(c *fiber.Ctx) error {

	status, err := json.Marshal(map[string]string{
		"status": "ok",
	})

	if err != nil {
		return err
	}
	c.Send(status)
	return nil
}
