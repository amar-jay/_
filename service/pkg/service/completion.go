package service

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/amar-jay/comrade/pkg/config"
	lksdk "github.com/livekit/server-sdk-go"
	"github.com/sashabaranov/go-openai"
)

type CompletionService interface {
	Complete(
		ctx context.Context,
		text string,
		participant *lksdk.RemoteParticipant,
		room *lksdk.Room,
		events []RoomEvent,
	) (*ChatStream, error)
}

type completionService struct {
	client *openai.Client
}

func NewCompletionService(client *openai.Client) CompletionService {
	return &completionService{client: client}
}

func (c *completionService) Complete(ctx context.Context, text string, participant *lksdk.RemoteParticipant, room *lksdk.Room, events []RoomEvent) (*ChatStream, error) {
	sb := strings.Builder{}
	// messages := make([]openai.ChatCompletionMessage, 0, len(events)+5)
	// var messages []openai.ChatCompletionMessage
	messages := []openai.ChatCompletionMessage{}
	participants := room.GetParticipants()
	for i, p := range participants {
		sb.WriteString(p.Name())
		if n := len(participants); i < n {
			sb.WriteString(", ")
		}
	}
	names := sb.String()
	sb.Reset()
	messages = append(messages, openai.ChatCompletionMessage{
		Role: openai.ChatMessageRoleSystem,
		Name: config.BotIdentity,
		Content: fmt.Sprintf(
			`Hey, you are %s, a voice assistant in a 
			live conference meeting by Comrade. 
			Keep your responses short and to the point. 
			Please do not use any profanity or offensive language. 
			If you are not sure what to say, just say %s. 
			If your response is a question end it with a question mark. 
			In this conference there are %d participants in the chat: %s. 
			Current language: %s language.`,
			config.BotIdentity,
			"\"I don't know\"",
			len(participants),
			names,
			config.BotLanguage,
		),
	})

	for _, e := range events {
		if e.Speech == nil {
			continue
		}
		if e.Speech.IsBot {
			messages = append(messages,
				openai.ChatCompletionMessage{
					Role:    openai.ChatMessageRoleAssistant,
					Content: e.Speech.Text,
					Name:    config.BotIdentity,
				},
			)
		} else {
			messages = append(messages,
				openai.ChatCompletionMessage{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("%s: %s\n", e.Speech.ParticipantName, e.Speech.Text),
					Name:    e.Speech.ParticipantName,
				})
		}

		if e.Join == nil {
			continue
		}

		if !e.Join.Joined {
			messages = append(messages, openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleSystem,
				Content: fmt.Sprint(e.Join.ParticipantName, " has left the meeting."),
			})
		} else {
			messages = append(messages, openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleSystem,
				Content: fmt.Sprint(e.Join.ParticipantName, " has joined the meeting."),
			})

		}
	}
	return &ChatStream{}, nil
}

// this is just a wrapper for openai stream it returns string instead of byte array
type ChatStream struct {
	stream *openai.ChatCompletionStream
}

func (cs *ChatStream) Close() {
	cs.stream.Close()
}

func (cs *ChatStream) String() (string, error) {
	sb := strings.Builder{}

	for {
		res, err := cs.stream.Recv()
		if err != nil {
			content := strings.Trim(sb.String(), " ")
			if err == io.EOF && len(content) > 0 {
				return content, nil
			}
			return "", err
		}

		content := res.Choices[0].Delta.Content
		sb.WriteString(content)

		content = strings.Trim(content, " ")
		if strings.HasSuffix(content, ".") {
			return sb.String(), nil
		}

	}
}
