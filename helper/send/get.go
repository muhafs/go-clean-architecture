package send

import "net/http"

func Get(url string) (stringContent string, err error) {
	// send request
	response, err := SendRequest(http.MethodPut, url, nil)
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
