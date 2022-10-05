package http

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func PostProxy(target string, path string, body string, contentType string) (*http.Response, error) {
	// `{"medical_text": "back pain"}`
	jsonBody := []byte(body)
	bodyReader := bytes.NewReader(jsonBody)

	requestURL := fmt.Sprintf("http://%s%s", target, path)

	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)

	client := http.Client{
		Timeout: 15 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
