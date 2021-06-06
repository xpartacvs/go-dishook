package dishook

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func Send(url string, payload Payload) ([]byte, error) {
	u := Url(url)
	if err := u.validate(); err != nil {
		return nil, err
	}

	reqBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
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
