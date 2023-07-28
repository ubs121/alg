package concurrency

import (
	"testing"
	"time"
)

type Storage interface {
	Set(key string, value any)
	Get(key string) any // Get() should wait all active Set() calls and return most last updated value
}

type storage struct {
	data     map[string]any
	setCh    chan setRequest
	getCh    chan getRequest
	shutdown chan struct{}
}

type setRequest struct {
	key   string
	value any
}

type getRequest struct {
	key      string
	response chan any
}

func NewStorage() Storage {
	s := &storage{
		data:     make(map[string]any),
		setCh:    make(chan setRequest),
		getCh:    make(chan getRequest),
		shutdown: make(chan struct{}),
	}
	go s.processRequests()
	return s
}

func (s *storage) Set(key string, value any) {
	s.setCh <- setRequest{
		key:   key,
		value: value,
	}
}

func (s *storage) Get(key string) any {
	responseCh := make(chan any)
	s.getCh <- getRequest{
		key:      key,
		response: responseCh,
	}
	return <-responseCh
}

func (s *storage) processRequests() {
	for {
		select {
		// we prioritize the s.setCh by placing it first in the select statement
		case req := <-s.setCh:
			s.data[req.key] = req.value
		case req := <-s.getCh:
			req.response <- s.data[req.key]
		case <-s.shutdown:
			close(s.setCh)
			close(s.getCh)
			return
		}
	}
}

func TestStorage(t *testing.T) {
	storage := NewStorage()

	// Set a value for a key
	storage.Set("key", "value")

	// Get the value immediately
	value := storage.Get("key")
	if value != "value" {
		t.Errorf("Expected value 'value', but got '%v'", value)
	}

	// Set a value for the same key after a delay
	go func() {
		time.Sleep(100 * time.Nanosecond)
		storage.Set("key", "updated value")
	}()

	time.Sleep(120 * time.Nanosecond)

	// Get the value after the update
	value = storage.Get("key")
	if value != "updated value" {
		t.Errorf("Expected value 'updated value', but got '%v'", value)
	}
}
