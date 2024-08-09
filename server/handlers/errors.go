package handlers

import "errors"

var (
	ErrApi = errors.New("failed to retrieve data from external API")
	ErrReq = errors.New("failed to process request")

	ErrParam      = errors.New("got wrong parameters")
	ErrStatusCode = errors.New("got unsuccessful status code")
	ErrJsonDecode = errors.New("failed to unmarshall JSON to struct")
)
