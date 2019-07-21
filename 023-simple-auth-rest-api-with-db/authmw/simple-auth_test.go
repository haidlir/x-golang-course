package authmw_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	nauthmw "github.com/haidlir/x-golang-course/022-simple-auth-rest-api/authmw"

	"github.com/gorilla/mux"
)

func handlerTest(w http.ResponseWriter, r *http.Request) {
	resp := []byte(`{"status": "ok}`)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func TestSiswaHandler(t *testing.T) {
	testURL := "/api/test"
	// Setup Routing
	r := mux.NewRouter()
	r.HandleFunc(testURL, handlerTest).Methods(http.MethodGet)
	// Auth Middleware
	amw := nauthmw.NewAuthMiddleware()
	r.Use(amw.Middleware)
	// Create httptest Server
	httpServer := httptest.NewServer(r)
	defer httpServer.Close()
	serverURL, _ := url.Parse(httpServer.URL)
	// Hit API Endpoint
	targetPath := fmt.Sprintf("%v%v", serverURL, testURL)
	// Auth OK
	t.Run("Auth OK", func(t *testing.T) {
		// Hit API Endpoint
		req, _ := http.NewRequest(http.MethodGet, targetPath, nil)
		req.Header.Add("X-Session-Token", "00000000")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Unable to hit API: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("status code %v instead of 200", resp.StatusCode)
		}
	})
	// Auth NOK
	t.Run("Auth NOK", func(t *testing.T) {
		// Hit API Endpoint
		req, _ := http.NewRequest(http.MethodGet, targetPath, nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Unable to hit API: %v", err)
		}
		if resp.StatusCode != http.StatusForbidden {
			t.Fatalf("status code %v instead of 403", resp.StatusCode)
		}
	})
}
