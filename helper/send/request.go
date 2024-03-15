package send

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func EncodeRequest(body any) (*bytes.Buffer, error) {
	postBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	reqBody := bytes.NewBuffer(postBody)

	return reqBody, nil
}

func SendRequest(method string, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodDelete, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(req)

	return response, err
}

func DecodeResponse(responseBody io.ReadCloser) (string, error) {
	byteContect, err := io.ReadAll(responseBody)
	if err != nil {
		return "", err
	}

	return string(byteContect), nil
}
