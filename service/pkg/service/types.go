package service

// Events
type SpeechEvent struct {
	Text            string
	IsBot           bool
	ParticipantName string
}

type JoinOrLeaveEvent struct {
	ParticipantName string
	Joined          bool
}

type RoomEvent struct {
	Join   *JoinOrLeaveEvent
	Speech *SpeechEvent
}
