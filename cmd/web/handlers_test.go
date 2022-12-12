package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/exvimmer/lets_go/snippetbox/internal/assert"
)

func TestPing(t *testing.T) {
	rr := httptest.NewRecorder() // the responseRecorder
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		// NOTE: Typically you should call t.Fatal() in situations where it doesn't
		// make sense to continue the current test — such as an error during a
		// setup step, or where an unexpected error from a Go standard library
		// function means you can't proceed with the test.
		t.Fatal(err)
	}

	ping(rr, r)
	rs := rr.Result()
	defer rs.Body.Close()
	assert.Equal(t, rs.StatusCode, http.StatusOK)

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	bytes.TrimSpace(body)
	assert.Equal(t, string(body), "OK")
}
