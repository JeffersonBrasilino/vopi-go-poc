package http_test

import (
	"bytes"
	"errors"
	httpstd "net/http"
	"net/http/httptest"
	"testing"

	"context"

	"github.com/gin-gonic/gin"
	"github.com/vopi-go-poc/internal/chat/infra/http"
	"github.com/vopi-go-poc/internal/chat/usecase/createchat"
	"github.com/vopi-go-poc/internal/core/mocks"
	"github.com/vopi-go-poc/internal/core/otel"
)

type fakeUseCase struct {
	ExecuteFn func(ctx context.Context, input *createchat.CreateChatInput, tracer otel.OtelTracer) (any, error)
}

func (f *fakeUseCase) Execute(ctx context.Context, input *createchat.CreateChatInput, tracer otel.OtelTracer) (any, error) {
	return f.ExecuteFn(ctx, input, tracer)
}

func TestCreateChatHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("should success", func(t *testing.T) {
		r := gin.New()
		uc := &fakeUseCase{
			ExecuteFn: func(ctx context.Context, input *createchat.CreateChatInput, tracer otel.OtelTracer) (any, error) {
				return gin.H{"ok": true}, nil
			},
		}
		tracer := &mocks.MockTrace{}
		r.POST("/chat", func(c *gin.Context) {
			http.CreateChatHandler(c, uc, tracer)
		})
	
		body := `{"channelId":"ch1","botName":"bot","participants":[{"name":"n","document":"d","contact":"c"}],"messages":[{"content":"msg","status":"sent","sender":"n"}]}`
		req := httptest.NewRequest(httpstd.MethodPost, "/chat", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != httpstd.StatusCreated {
			t.Errorf("esperava status %d, obteve %d", httpstd.StatusCreated, w.Code)
		}
	})
	t.Run("should fail on empty body", func(t *testing.T) {
		r := gin.New()
		uc := &fakeUseCase{
			ExecuteFn: func(ctx context.Context, input *createchat.CreateChatInput, tracer otel.OtelTracer) (any, error) {
				return nil, nil
			},
		}
		tracer := &mocks.MockTrace{}
		r.POST("/chat", func(c *gin.Context) {
			http.CreateChatHandler(c, uc, tracer)
		})

		req := httptest.NewRequest(httpstd.MethodPost, "/chat", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != httpstd.StatusBadRequest {
			t.Errorf("esperava status %d, obteve %d", httpstd.StatusBadRequest, w.Code)
		}
	})

	t.Run("should fail on use case error", func(t *testing.T) {
		r := gin.New()
		uc := &fakeUseCase{
			ExecuteFn: func(ctx context.Context, input *createchat.CreateChatInput, tracer otel.OtelTracer) (any, error) {
				return nil, errors.New("fail")
			},
		}
		tracer := &mocks.MockTrace{}
		r.POST("/chat", func(c *gin.Context) {
			http.CreateChatHandler(c, uc, tracer)
		})

		body := `{"channelId":"ch1","botName":"bot","participants":[{"name":"n","document":"d","contact":"c"}],"messages":[{"content":"msg","status":"sent","sender":"n"}]}`
		req := httptest.NewRequest(httpstd.MethodPost, "/chat", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != httpstd.StatusUnprocessableEntity {
			t.Errorf("esperava status %d, obteve %d", httpstd.StatusUnprocessableEntity, w.Code)
		}
	})
}