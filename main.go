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
		fmt.Fprintln(os.Stderr, "sendme: please specify username and post content")
		return
	}
	user := os.Args[1]
	content := os.Args[2]
	err := Post(PostRequest{user, content})
	if err != nil {
		fmt.Fprintf(os.Stderr, "sendme: %v\n", err)
	}
}

type PostRequest struct {
	User    string `json:"user"`
	Content string `json:"content"`
}

func Post(postRequest PostRequest) error {
	body, err := json.Marshal(postRequest)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(body)
	response, err := http.Post("https://sendme.josueart40.workers.dev", "application/json", reader)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		return fmt.Errorf("server returned %s", response.Status)
	}
	return nil
}
