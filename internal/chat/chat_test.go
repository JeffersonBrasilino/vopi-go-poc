package chat_test

import (
	"database/sql"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/vopi-go-poc/internal/chat"
)

type fakeDb struct{}

func (f *fakeDb) Instance() *sql.DB { return nil }
func (f *fakeDb) Disconnect() error { return nil }

func TestNewChatModule(t *testing.T) {
	fake := &fakeDb{}
	module := chat.NewChatModule(fake)
	t.Run("should create success", func(t *testing.T) {
		if module == nil {
			t.Fatalf("esperava módulo criado, mas obteve nil")
		}
	})
}

func TestWithHttpFunction(t *testing.T){
	gin.SetMode(gin.TestMode)
	r := gin.New()
	fake := &fakeDb{}
	module := chat.NewChatModule(fake)
	result := module.WithHttp(r)
	t.Run("should with http", func(t *testing.T) {
		if result != module {
			t.Errorf("esperava retorno do próprio módulo")
		}
	})
	t.Run("should call CreateChatHandler via HTTP", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/chat",nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
	})
}