package service

import (
	"context"

	"github.com/mes1234/golock/adapters"
)

// AccessService Access interface describes a service that insert and retrieve data
type AccessService interface {

	// Add item to locker
	Add(
		ctx context.Context,
		request adapters.AddItemRequest,
	) (adapters.AddItemResponse, error) // status of operation

	// Get item from locker
	Get(
		ctx context.Context,
		request adapters.GetItemRequest,
	) (adapters.GetItemResponse, error) // uncrypted content

	// Remove item from locker
	Remove(
		ctx context.Context,
		request adapters.RemoveItemRequest,
	) (adapters.RemoveItemResponse, error) // status of operation

	// NewLocker Add new locker
	NewLocker(
		ctx context.Context,
		request adapters.AddLockerRequest, // Identification of client
	) (adapters.AddLockerResponse, error) // Identification of locker
}
