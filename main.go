package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Please specify username and post content.")
		return
	}
	username := os.Args[1]
	postContent := os.Args[2]
	err := Post(username, postContent)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

type Request struct {
	User    string `json:"user"`
	Content string `json:"content"`
}

func Post(username, postContent string) error {
	body, err := json.Marshal(Request{username, postContent})
	if err != nil {
		return err
	}
	reader := bytes.NewReader(body)
	http.Post("https://sendme.josueart40.workers.dev", "application/json", reader)
	return nil
}
