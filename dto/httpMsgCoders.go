package dto

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

// Decode Http inbound message to domain accepted message
func DecodeHttpAddLockerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var requestHttp AddLockerHttpInboundDto
	if err := json.NewDecoder(r.Body).Decode(&requestHttp); err != nil {
		return nil, err
	}
	clientId, err := uuid.Parse(requestHttp.Client)
	if err != nil {
		return nil, err
	}

	request := AddLockerRequest{
		Client: clientId,
	}

	return request, nil
}

// Decode response to Http accepted representation
func EncodeHttpAddLockerResponse(
	_ context.Context,
	w http.ResponseWriter,
	response interface{}) error {

	domainResponse := response.(AddLockerResponse)

	responseHttp := AddLockerHttpOutboundDto{
		LockerId: domainResponse.LockerId.String(),
	}

	json.NewEncoder(w).Encode(responseHttp)

	return nil
}
