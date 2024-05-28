package event_test

import (
	"go-complaint/domain"
	"testing"
)

func TestDomainEventPublisherInstance(t *testing.T) {
	publisher := domain.DomainEventPublisherInstance()
	if publisher == nil {
		t.Error("Expected a non-nil instance of DomainEventPublisher")
	}
}
