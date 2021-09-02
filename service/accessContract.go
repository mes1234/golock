package service

import (
	"context"

	"github.com/mes1234/golock/dto"
)

//Access interface describes a service that insert and retrieve data
type AccessService interface {

	// Add item to locker
	Add(
		ctx context.Context,
		request dto.AddItemRequest,
	) (dto.AddItemResponse, error) // status of operation

	// Get item from locker
	Get(
		ctx context.Context,
		request dto.GetItemRequest,
	) (dto.GetItemResponse, error) // uncrypted content

	// Remove item from locker
	Remove(
		ctx context.Context,
		request dto.RemoveItemRequest,
	) (dto.RemoveItemResponse, error) // status of operation

	// Add new locker
	NewLocker(
		ctx context.Context,
		request dto.AddLockerRequest, // Identification of client
	) (dto.AddLockerResponse, error) // Identification of locker
}
