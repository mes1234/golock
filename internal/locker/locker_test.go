package locker_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/mes1234/golock/internal/keys"
	"github.com/mes1234/golock/internal/locker"
)

func createLocker() locker.Locker {
	return locker.GetMemoryLocker(uuid.New(), uuid.New())
}

func TestAddingItemToLocker(t *testing.T) {
	l := createLocker()
	content := []byte{0x01}

	resChan := make(chan error)

	go l.AddItem("dummy", keys.Value{}, content, resChan)

	err := <-resChan
	if err != nil {
		t.Fatalf("Error during adding item to locker: %v", err)
	}
}

func TestGetItemFromLocker(t *testing.T) {
	l := createLocker()
	content := []byte{0x01}

	resChan := make(chan error)

	go l.AddItem("dummy", keys.Value{}, content, resChan)

	err := <-resChan
	if err != nil {
		t.Fatalf("Error during adding item to locker: %v", err)
	}
	resChanGet := make(chan []byte)

	go l.GetItem(keys.Value{}, "dummy", resChanGet)
	res := <-resChanGet
	for i, s := range res {
		if s != content[i] {
			t.Fatalf("Value restored is not the same as saved in locker")
		}
	}
}

func TestGetItemFailedLocker(t *testing.T) {
	l := createLocker()
	content := []byte{0x01}

	resChan := make(chan error)

	go l.AddItem("dummy", keys.Value{}, content, resChan)

	err := <-resChan
	if err != nil {
		t.Fatalf("Error during adding item to locker: %v", err)
	}
	resChanGet := make(chan []byte)

	go l.GetItem(keys.Value{}, "dummy2", resChanGet)

	// Make sure that the function does close the channel
	_, ok := <-resChanGet

	// If we can recieve on the channel then it is NOT closed
	if ok {
		t.Error("Channel is not closed")
	}
}

func TestRemoveItemFailedGettingLocker(t *testing.T) {
	l := createLocker()
	content := []byte{0x01}

	resChan := make(chan error)

	go l.AddItem("dummy", keys.Value{}, content, resChan)

	err := <-resChan
	if err != nil {
		t.Fatalf("Error during removing item from locker: %v", err)
	}
	resChanGet := make(chan []byte)

	go l.RemoveItem("dummy", resChan)
	err = <-resChan
	if err != nil {
		t.Fatalf("Error during removing item from locker: %v", err)
	}
	go l.GetItem(keys.Value{}, "dummy", resChanGet)

	// Make sure that the function does close the channel
	_, ok := <-resChanGet

	// If we can recieve on the channel then it is NOT closed
	if ok {
		t.Error("Channel is not closed")
	}
}

func TestRemoveNonexistingItemFailedLocker(t *testing.T) {
	l := createLocker()
	resChan := make(chan error)

	go l.RemoveItem("dummy", resChan)
	err := <-resChan
	if err == nil {
		t.Fatalf("No error raised when trying to remove non existing item ")
	}
}
