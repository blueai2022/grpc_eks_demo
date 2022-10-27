package healthapi

import (
	"bytes"
	"net/http"
	"time"
)

const (
	defaultTimeoutSecs = 15
)

func post(url string, body []byte, contentType string) (*http.Response, error) {
	reader := bytes.NewReader(body)

	req, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)

	client := http.Client{
		Timeout: defaultTimeoutSecs * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
