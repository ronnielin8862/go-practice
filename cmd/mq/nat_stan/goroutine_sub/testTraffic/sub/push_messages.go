package sub

import (
	"encoding/json"
)

const (
	RoomTextLiveMessageType   = "textLive"
	RoomScoreLiveMessageType  = "scoreLive"
	RoomStatsLiveMessageType  = "statsLive"
	RoomLineupMessageType     = "lineupLive"
	RoomRecordMessageType     = "recordLive"
	RoomIndexMessageType      = "indexLive"
	RoomUserLoveMessageType   = "attentAnchor"
	RoomQuizMessageType       = "quiz"
	RoomQuizPromptMessageType = "quizPrompt"
	RoomClearChatType         = "clearChat"
	RoomMessageAliveTime      = 24
)

type (
	// RoomUserLoveMessage 用户关注了主播后的广播消息
	RoomUserLoveMessage struct {
		SentMessageStruct
		ClientName string `json:"from_client_name"`
		Content    string `json:"content"`
	}

	SentMessageStruct struct {
		Type    string `json:"type"`
		Message string `default:"" json:"message,omitempty"`
	}

	RoomQuizMessage struct {
		SentMessageStruct
		Content string `json:"content"`
	}

	RoomBasketballRecordMessage struct {
		SentMessageStruct
		Record BasketballTeamRecord `json:"team_record"`
	}

	RoomBasketballStatsMessage struct {
		SentMessageStruct
		StatusLive []BasketballStatus `json:"stats_live"`
	}

	RoomBasketballScoreMessage struct {
		SentMessageStruct
		Scores BasketballScore `json:"basketball_scores"`
	}

	RoomTextLiveMessage struct {
		SentMessageStruct
		TextLive []TextLiveStruct `json:"text_live"`
	}

	RoomLineupMessage struct {
		SentMessageStruct
		Lineup Lineup `json:"line_up"`
	}

	RoomFootballStatsMessage struct {
		SentMessageStruct
		StatusLive []FootballStatus `json:"stats_live"`
	}

	RoomFootballScoreMessage struct {
		SentMessageStruct
		ScoreLive FootballScores `json:"football_scores"`
	}

	RoomIndexMessage struct {
		SentMessageStruct
		Index []Index `json:"message"`
	}

	RoomMessageToNats struct {
		RoomId  int    `json:"room_id"`
		Type    string `json:"type"`
		Content []byte `json:"content"`
	}

	ClearChatMessage struct {
		SentMessageStruct
		ChatId int `json:"chat_id"`
		UserId int `json:"user_id"`
		RoomId int `json:"room_id"`
	}

	ForwardMsg struct {
		Type    string `json:"type"`
		Content string `json:"content"`
	}

	ForwardMsgToRoom struct {
		ForwardMsg
		RoomID int `json:"room_id"`
	}

	ForwardMsgToAuthedUser struct {
		ForwardMsg
		UserID int `json:"user_d"`
	}

	ForwardMsgToAlUsers struct {
		ForwardMsg
	}
)

func (m *SentMessageStruct) GetType() string {
	return m.Type
}

func (m *SentMessageStruct) ToJSON() []byte {
	b, _ := json.Marshal(m)
	return b
}
