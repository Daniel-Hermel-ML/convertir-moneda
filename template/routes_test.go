package template

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetupRoutes(t *testing.T) {
	r := gin.Default()
	SetupRoutes(r)

	testCases := []struct {
		path   string
		status int
	}{
		{"/V1/convert/real-to-dollar", http.StatusOK},
		{"/V1/convert/real-to-euro", http.StatusOK},
		{"/V1/ping", http.StatusOK},
	}

	for _, tc := range testCases {
		t.Run(tc.path, func(t *testing.T) {
			req, _ := http.NewRequest("GET", tc.path, nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			if w.Code != tc.status {
				t.Errorf("Expected status %d, got %d", tc.status, w.Code)
			}
		})
	}
}
