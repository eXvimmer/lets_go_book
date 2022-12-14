package main

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/exvimmer/lets_go/snippetbox/internal/assert"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()
	code, _, body := ts.get(t, "/ping")
	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, body, "OK")
}

func TestSnippetView(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{
			name:     "Valid ID",
			urlPath:  "/snippet/view/1",
			wantCode: http.StatusOK,
			wantBody: "An old silent pond...",
		},
		{
			name:     "Non-existent ID",
			urlPath:  "/snippet/view/2",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Negative ID",
			urlPath:  "/snippet/view/-1",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Decimal ID",
			urlPath:  "/snippet/view/1.23",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "String ID",
			urlPath:  "/snippet/view/foo",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Empty ID",
			urlPath:  "/snippet/view/",
			wantCode: http.StatusNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			code, _, body := ts.get(t, test.urlPath)
			assert.Equal(t, code, test.wantCode)
			if test.wantBody != "" {
				assert.StringConaints(t, body, test.wantBody)
			}
		})
	}
}

func TestUserSignUp(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()
	_, _, body := ts.get(t, "/user/signup")

	validCSRFToken := extractCSRFToken(t, body)

	const (
		validName     = "Mustafa"
		validPassword = "validPa$$word"
		validEmail    = "mustafa@gmail.com"
		formTag       = `<form action="/user/signup" method="POST" novalidate>`
	)

	tests := []struct {
		name           string
		userName       string
		userEmail      string
		userPassword   string
		csrfToken      string
		wantStatusCode int
		wantFormTag    string
	}{
		{
			name:           "Valid submission",
			userName:       validName,
			userEmail:      validEmail,
			userPassword:   validPassword,
			csrfToken:      validCSRFToken,
			wantStatusCode: http.StatusSeeOther,
		},
		{
			name:           "Invalid CSRF Token",
			userName:       validName,
			userEmail:      validEmail,
			userPassword:   validPassword,
			csrfToken:      "wrongToken",
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:           "Empty name",
			userName:       "",
			userEmail:      validEmail,
			userPassword:   validPassword,
			csrfToken:      validCSRFToken,
			wantStatusCode: http.StatusUnprocessableEntity,
			wantFormTag:    formTag,
		},
		{
			name:           "Empty email",
			userName:       validName,
			userEmail:      "",
			userPassword:   validPassword,
			csrfToken:      validCSRFToken,
			wantStatusCode: http.StatusUnprocessableEntity,
			wantFormTag:    formTag,
		},
		{
			name:           "Empty password",
			userName:       validName,
			userEmail:      validEmail,
			userPassword:   "",
			csrfToken:      validCSRFToken,
			wantStatusCode: http.StatusUnprocessableEntity,
			wantFormTag:    formTag,
		},
		{
			name:           "Invalid email",
			userName:       validName,
			userEmail:      "mustafa@gmail.",
			userPassword:   validPassword,
			csrfToken:      validCSRFToken,
			wantStatusCode: http.StatusUnprocessableEntity,
			wantFormTag:    formTag,
		},
		{
			name:           "Short password",
			userName:       validName,
			userEmail:      validEmail,
			userPassword:   "pa$$",
			csrfToken:      validCSRFToken,
			wantStatusCode: http.StatusUnprocessableEntity,
			wantFormTag:    formTag,
		},
		{
			name:           "Duplicate email",
			userName:       validName,
			userEmail:      "duplicated@gmail.com", // must be identical to mocks/users
			userPassword:   validPassword,
			csrfToken:      validCSRFToken,
			wantStatusCode: http.StatusUnprocessableEntity,
			wantFormTag:    formTag,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			form := url.Values{}
			form.Add("name", test.userName)
			form.Add("email", test.userEmail)
			form.Add("password", test.userPassword)
			form.Add("csrf_token", test.csrfToken)

			statusCode, _, body := ts.postForm(t, "/user/signup", form)
			assert.Equal(t, statusCode, test.wantStatusCode)

			if test.wantFormTag != "" {
				assert.StringConaints(t, body, test.wantFormTag)
			}
		})
	}
}
