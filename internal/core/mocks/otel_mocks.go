package mocks

import (
	"context"

	"github.com/vopi-go-poc/internal/core/otel"
)

type SpanMock struct {}
func (s *SpanMock) End() {}
func (s *SpanMock) AddEvent(eventMessage string, attributes ...otel.OtelAttribute) {}
func (s *SpanMock) SetStatus(status otel.SpanStatus, description string) {}
func (s *SpanMock) Success(message string) {}
func (s *SpanMock) Error(err error, message string) {}


type MockTrace struct{}
func (m *MockTrace) Start(ctx context.Context, name string, attributes ...otel.OtelAttribute) (context.Context, otel.OtelSpan) {
	return context.Background(), &SpanMock{}
}