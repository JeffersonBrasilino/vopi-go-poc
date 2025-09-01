package otel

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	traceTypes "go.opentelemetry.io/otel/trace"
)

type SpanStatus int

const (
	SpanStatusOK SpanStatus = iota
	SpanStatusError
)

type otelTrace struct {
	tracer traceTypes.Tracer
}

type otelSpan struct {
	span traceTypes.Span
	ctx  context.Context
}

func InitTrace(serviceName string) *otelTrace {
	tracer := otel.Tracer(serviceName)
	return &otelTrace{
		tracer: tracer,
	}
}

func (t *otelTrace) Start(
	ctx context.Context,
	name string,
	attributes ...OtelAttribute,
) (context.Context, OtelSpan) {
	ctx, span := t.tracer.Start(ctx, name, makeAttributes(attributes))
	return ctx, &otelSpan{
		span: span,
		ctx:  ctx,
	}
}

func (s *otelSpan) End() {
	s.span.End()
}

func (s *otelSpan) AddEvent(eventMessage string, attributes ...OtelAttribute) {
	s.span.AddEvent(eventMessage, makeAttributes(attributes))
}

func (o *otelSpan) SetStatus(status SpanStatus, description string) {
	o.span.SetStatus(status.otelStatus(), description)
}

func (s *otelSpan) Success(message string) {
	s.SetStatus(SpanStatusOK, message)
}

func (s *otelSpan) Error(err error, message string) {
	s.SetStatus(SpanStatusError, message)
}

func (s *SpanStatus) otelStatus() codes.Code {
	switch *s {
	case SpanStatusOK:
		return codes.Ok
	case SpanStatusError:
		return codes.Error
	default:
		return codes.Unset
	}
}