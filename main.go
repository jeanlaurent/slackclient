package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type slackMessage struct {
	Text     string `json:"text"`
	Channel  string `json:"channel"`
	Username string `json:"username"`
	Icon     string `json:"icon_emoji"`
}

func main() {
	webHook := os.Getenv("SLACK_WEBHOOK")
	if webHook == "" {
		fmt.Println("SLACK_WEBHOOK missing")
		os.Exit(3)
	}

	if len(os.Args) < 1 {
		fmt.Println("text missing")
		os.Exit(4)
	}

	text := strings.Join(os.Args[1:], " ")

	payload, err := json.Marshal(slackMessage{Text: text})
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	req, err := http.NewRequest("POST", webHook, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println(err)
		os.Exit(5)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
	fmt.Println(resp.Header)
}
