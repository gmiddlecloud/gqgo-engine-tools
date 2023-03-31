package handler

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/rest/internal/response"
)

func init() {
	log.SetOutput(io.Discard)
}

func TestTimeout(t *testing.T) {
	timeoutHandler := TimeoutHandler(time.Millisecond)
	handler := timeoutHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Minute)
	}))

	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusServiceUnavailable, resp.Code)
}

func TestWithinTimeout(t *testing.T) {
	timeoutHandler := TimeoutHandler(time.Second)
	handler := timeoutHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Millisecond)
	}))

	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestWithTimeoutTimedout(t *testing.T) {
	timeoutHandler := TimeoutHandler(time.Millisecond)
	handler := timeoutHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Millisecond * 10)
		w.Write([]byte(`foo`))
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusServiceUnavailable, resp.Code)
}

func TestWithoutTimeout(t *testing.T) {
	timeoutHandler := TimeoutHandler(0)
	handler := timeoutHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
	}))

	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestTimeoutPanic(t *testing.T) {
	timeoutHandler := TimeoutHandler(time.Minute)
	handler := timeoutHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("foo")
	}))

	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	resp := httptest.NewRecorder()
	assert.Panics(t, func() {
		handler.ServeHTTP(resp, req)
	})
}

func TestTimeoutWebsocket(t *testing.T) {
	timeoutHandler := TimeoutHandler(time.Millisecond)
	handler := timeoutHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Millisecond * 10)
	}))

	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	req.Header.Set(headerUpgrade, valueWebsocket)
	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestTimeoutWroteHeaderTwice(t *testing.T) {
	timeoutHandler := TimeoutHandler(time.Minute)
	handler := timeoutHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`hello`))
		w.Header().Set("foo", "bar")
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestTimeoutWriteBadCode(t *testing.T) {
	timeoutHandler := TimeoutHandler(time.Minute)
	handler := timeoutHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(1000)
	}))

	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	resp := httptest.NewRecorder()
	assert.Panics(t, func() {
		handler.ServeHTTP(resp, req)
	})
}

func TestTimeoutClientClosed(t *testing.T) {
	timeoutHandler := TimeoutHandler(time.Minute)
	handler := timeoutHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))

	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	ctx, cancel := context.WithCancel(context.Background())
	req = req.WithContext(ctx)
	cancel()
	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, statusClientClosedRequest, resp.Code)
}

func TestTimeoutHijack(t *testing.T) {
	resp := httptest.NewRecorder()

	writer := &timeoutWriter{
		w: &response.WithCodeResponseWriter{
			Writer: resp,
		},
	}

	assert.NotPanics(t, func() {
		writer.Hijack()
	})

	writer = &timeoutWriter{
		w: &response.WithCodeResponseWriter{
			Writer: mockedHijackable{resp},
		},
	}

	assert.NotPanics(t, func() {
		writer.Hijack()
	})
}

func TestTimeoutPusher(t *testing.T) {
	handler := &timeoutWriter{
		w: mockedPusher{},
	}

	assert.Panics(t, func() {
		handler.Push("any", nil)
	})

	handler = &timeoutWriter{
		w: httptest.NewRecorder(),
	}
	assert.Equal(t, http.ErrNotSupported, handler.Push("any", nil))
}

type mockedPusher struct{}

func (m mockedPusher) Header() http.Header {
	panic("implement me")
}

func (m mockedPusher) Write(bytes []byte) (int, error) {
	panic("implement me")
}

func (m mockedPusher) WriteHeader(statusCode int) {
	panic("implement me")
}

func (m mockedPusher) Push(target string, opts *http.PushOptions) error {
	panic("implement me")
}
