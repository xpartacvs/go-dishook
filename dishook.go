package dishook

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func Post(url string, payload Payload) (reso *http.Response, err error) {
	u := Url(url)
	if err := u.validate(); err != nil {
		return nil, err
	}

	reqBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return http.Post(url, "application/json", bytes.NewBuffer(reqBody))
}

func Send(url string, payload Payload) ([]byte, error) {
	resp, err := Post(url, payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bytesResp, nil
}
