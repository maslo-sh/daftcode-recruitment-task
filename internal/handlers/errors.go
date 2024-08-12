package handlers

import "errors"

var (
	ErrApi              = errors.New("failed to retrieve data from external API")
	ErrWrongFloatFormat = errors.New("wrong float format")
	ErrNoSuchCrypto     = errors.New("not found given cryptocurrency")
	ErrParam            = errors.New("got wrong parameters")
	ErrStatusCode       = errors.New("got unsuccessful status code")
	ErrJsonDecode       = errors.New("failed to unmarshall JSON to struct")
)
