package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Msg struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type ErrorMsg struct {
	Error string `json:"error"`
}

type TimeMsg struct {
	Time time.Time `json:"time"`
}

type TextMsg struct {
	Text string `json:"text"`
}

func UnmarshalMsg(data []byte) error {
	var msg Msg
	if err := json.Unmarshal(data, &msg); err != nil {
		return err
	}
	fmt.Printf("Got message: %+v\n", msg)

	switch msg.Type {
	case "error":
		var e ErrorMsg
		if err := json.Unmarshal(msg.Data, &e); err != nil {
			return err
		}
		fmt.Println("ERROR:", e.Error)
	case "time":
		var t TimeMsg
		if err := json.Unmarshal(msg.Data, &t); err != nil {
			return err
		}
		fmt.Println("TIME: ", t.Time.Format(time.RFC3339Nano))
	case "text":
		var t TextMsg
		if err := json.Unmarshal(msg.Data, &t); err != nil {
			return err
		}
		fmt.Println("TEXT: ", t.Text)
	}
	fmt.Println("-----")
	return nil
}

func main() {
	msgs := []string{
		`{"type": "time", "data": { "time": "2014-12-25T15:00:00Z" } }`,
		`{"type": "text", "data": { "text": "hello world!"} }`,
		`{"type": "error", "data": { "error": "houston, we have a problem!" } }`,
	}

	for _, msg := range msgs {
		if err := UnmarshalMsg([]byte(msg)); err != nil {
			log.Fatal(err)
		}
	}
}
