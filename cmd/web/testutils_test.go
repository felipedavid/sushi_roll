package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/felipedavid/sushi_roll/internal/models/mocks"
)

func newTestApplication(t *testing.T) *app {
	templateCache, err := newTemplateCache()
	if err != nil {
		t.Fatal(err)
	}

	return &app{
		errLog:        log.New(io.Discard, "", 0),
		infoLog:       log.New(io.Discard, "", 0),
		templateCache: templateCache,
		games:         &mocks.GameModel{},
		comments:      &mocks.CommentModel{},
		categories:    &mocks.CategoryModel{},
	}
}

type testServer struct {
	*httptest.Server
}

// Create a newTestServer helper which initalizes and returns a new instance
// of our custom testServer type.
func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewServer(h)
	return &testServer{ts}
}

// Implement a get() method on our custom testServer type. This makes a GET
// request to a given url path using the test server client, and returns the
// response status code, headers and body.
func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, string) {
	rs, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}
	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	bytes.TrimSpace(body)
	return rs.StatusCode, rs.Header, string(body)
}

// Implement a get() method on our custom testServer type. This makes a GET
// request to a given url path using the test server client, and returns the
// response status code, headers and body.
func (ts *testServer) delete(t *testing.T, urlPath string) (int, http.Header, string) {
	r, _ := http.NewRequest("DELETE", ts.URL+urlPath, nil)
	rs, err := ts.Client().Do(r)
	if err != nil {
		t.Fatal(err)
	}
	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	bytes.TrimSpace(body)
	return rs.StatusCode, rs.Header, string(body)
}
