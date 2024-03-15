package send

import "net/http"

func Post(url string, body any) (stringContent string, err error) {
	// encode request
	reqBody, err := EncodeRequest(body)
	if err != nil {
		return
	}

	// send request
	response, err := SendRequest(http.MethodPost, url, reqBody)
	if err != nil {
		return
	}
	defer response.Body.Close()

	// extract response

	stringContent, err = DecodeResponse(response.Body)
	if err != nil {
		return
	}

	return
}
