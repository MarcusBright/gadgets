package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type slackWriter struct {
	client *http.Client
	url    string
	queue  chan []byte
}

func NewSlackWriter(url string) *slackWriter {
	slackWriter := &slackWriter{
		client: &http.Client{},
		url:    url,
		queue:  make(chan []byte, 128),
	}
	go slackWriter.send()
	return slackWriter
}

func (s *slackWriter) Write(p []byte) (n int, err error) {
	pc := bytes.Clone(p)
	s.queue <- pc
	return len(p), nil
}

func (s *slackWriter) Close() error {
	close(s.queue)
	return nil
}

func (s *slackWriter) send() {
	for msg := range s.queue {
		entry := map[string]any{}
		if err := json.Unmarshal(msg, &entry); err != nil {
			continue
		}
		if entry["level"] == "error" {
			text := fmt.Sprintf("[%s] %s", entry["server"].(string), entry["content"].(string))
			dataJson := fmt.Sprintf(`{"text": "%s"}`, text)
			req, _ := http.NewRequest("POST", s.url, strings.NewReader(dataJson))
			_, _ = s.client.Do(req)
		}
	}
}

func SendTo(url, message string) {
	dataJson := fmt.Sprintf(`{"text": "%s"}`, message)
	req, _ := http.NewRequest("POST", url, strings.NewReader(dataJson))
	_, _ = http.DefaultClient.Do(req)
}
