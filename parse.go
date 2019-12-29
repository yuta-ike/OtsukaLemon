package main

import (
	"encoding/json"
)

func UnmarshalLineRequest(data []byte) (LineRequest, error) {
	var r LineRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

type LineRequest struct {
	Destination string
	Event       []Event
}

type Event struct {
	ReplyToken string
	Type       string
	Mode       string
	Timestamp  int64
	Source     Source
	Message    Message
}

type Source struct {
	Type   string
	UserID string
}

type Message struct {
	ID   string
	Type string
	Text string
}
