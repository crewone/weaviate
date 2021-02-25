package clients

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWaitForStartup(t *testing.T) {
	t.Run("when the server is immediately ready", func(t *testing.T) {
		server := httptest.NewServer(&testReadyHandler{t: t})
		defer server.Close()
		c := New(server.URL)
		err := c.WaitForStartup(context.Background(), 50*time.Millisecond)

		assert.Nil(t, err)
	})

	t.Run("when the server is down", func(t *testing.T) {
		c := New("http://nothing-running-at-this-url")
		ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
		defer cancel()
		err := c.WaitForStartup(ctx, 50*time.Millisecond)

		require.NotNil(t, err)
		assert.Contains(t, err.Error(), "expired before remote was ready")
	})

	t.Run("when the server is alive, but not ready", func(t *testing.T) {
		server := httptest.NewServer(&testReadyHandler{
			t:         t,
			readyTime: time.Now().Add(1 * time.Minute),
		})
		c := New(server.URL)
		defer server.Close()
		ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
		defer cancel()
		err := c.WaitForStartup(ctx, 50*time.Millisecond)

		require.NotNil(t, err)
		assert.Contains(t, err.Error(), "expired before remote was ready")
	})

	t.Run("when the server is initially not ready, but then becomes ready",
		func(t *testing.T) {
			server := httptest.NewServer(&testReadyHandler{
				t:         t,
				readyTime: time.Now().Add(100 * time.Millisecond),
			})
			c := New(server.URL)
			defer server.Close()
			ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
			defer cancel()
			err := c.WaitForStartup(ctx, 50*time.Millisecond)

			require.Nil(t, err)
		})
}

type testReadyHandler struct {
	t *testing.T
	// the test handler will report as not ready before the time has passed
	readyTime time.Time
}

func (f *testReadyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	assert.Equal(f.t, "/.well-known/ready", r.URL.String())
	assert.Equal(f.t, http.MethodGet, r.Method)

	if time.Since(f.readyTime) < 0 {
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	w.WriteHeader(http.StatusNoContent)
}
