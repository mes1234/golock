package dto

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/mes1234/golock/adapters"
	"github.com/mes1234/golock/internal/locker"
)

// Decode Http inbound message to domain accepted message
func DecodeHttpAddLockerRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var requestHttp AddLockerHttpInboundDto
	if err := json.NewDecoder(r.Body).Decode(&requestHttp); err != nil {
		return nil, err
	}

	clientId, _ := uuid.Parse("b48f5b98-e9e7-40ce-b8cf-cdc4d2c59061")

	request := adapters.AddLockerRequest{
		ClientId: clientId,
	}

	return request, nil
}

// Decode Http inbound message to domain accepted message
func DecodeHttpAddItemRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var requestHttp AddItemHttpInboundDto
	if err := json.NewDecoder(r.Body).Decode(&requestHttp); err != nil {
		return nil, err
	}
	clientId, err := uuid.Parse(requestHttp.ClientId)
	if err != nil {
		return nil, err
	}
	lockerId, err := uuid.Parse(requestHttp.LockerId)
	if err != nil {
		return nil, err
	}
	secretid := (requestHttp.SecretId)

	content, err := base64.StdEncoding.DecodeString(requestHttp.Content)
	if err != nil {
		return nil, err
	}

	request := adapters.AddItemRequest{
		ClientId: clientId,
		LockerId: lockerId,
		SecretId: locker.SecretId(secretid),
		Content:  locker.PlainContent{Value: content},
	}

	return request, nil
}

// Decode Http inbound message to domain accepted message
func DecodeHttpRemoveItemRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var requestHttp RemoveItemHttpInboundDto
	if err := json.NewDecoder(r.Body).Decode(&requestHttp); err != nil {
		return nil, err
	}
	clientId, err := uuid.Parse(requestHttp.ClientId)
	if err != nil {
		return nil, err
	}
	lockerId, err := uuid.Parse(requestHttp.LockerId)
	if err != nil {
		return nil, err
	}
	secretid := (requestHttp.SecretId)

	request := adapters.RemoveItemRequest{
		ClientId: clientId,
		LockerId: lockerId,
		SecretId: locker.SecretId(secretid),
	}

	return request, nil
}

// Decode Http inbound message to domain accepted message
func DecodeHttpGetItemRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var requestHttp GetItemHttpInboundDto
	if err := json.NewDecoder(r.Body).Decode(&requestHttp); err != nil {
		return nil, err
	}
	clientId, err := uuid.Parse(requestHttp.ClientId)
	if err != nil {
		return nil, err
	}
	lockerId, err := uuid.Parse(requestHttp.LockerId)
	if err != nil {
		return nil, err
	}
	secretid := (requestHttp.SecretId)

	request := adapters.GetItemRequest{
		ClientId: clientId,
		LockerId: lockerId,
		SecretId: locker.SecretId(secretid),
	}

	return request, nil
}

// Encode response to Http accepted representation
func EncodeHttpAddLockerResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {

	domainResponse := response.(adapters.AddLockerResponse)

	responseHttp := AddLockerHttpOutboundDto{
		LockerId: domainResponse.LockerId.String(),
	}

	json.NewEncoder(w).Encode(responseHttp)

	return nil
}

// Encode response to Http accepted representation
func EncodeHttpAddItemResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {

	domainResponse := response.(adapters.AddItemResponse)

	responseHttp := AddItemHttpOutboundDto{
		Status: domainResponse.Status,
	}

	json.NewEncoder(w).Encode(responseHttp)

	return nil
}

// Encode response to Http accepted representation
func EncodeHttpRemoveItemResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {

	domainResponse := response.(adapters.RemoveItemResponse)

	responseHttp := RemoveItemHttpOutboundDto{
		Status: domainResponse.Status,
	}

	json.NewEncoder(w).Encode(responseHttp)

	return nil
}

// Encode response to Http accepted representation
func EncodeHttpGetItemResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {

	domainResponse := response.(adapters.GetItemResponse)

	responseHttp := GetItemHttpOutboundDto{
		Content: base64.RawStdEncoding.EncodeToString(domainResponse.Content.Value),
	}

	json.NewEncoder(w).Encode(responseHttp)

	return nil
}
