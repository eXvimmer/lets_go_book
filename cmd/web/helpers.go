package main

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/go-playground/form/v4"
	"github.com/justinas/nosurf"
)

func (app *application) render(
	w http.ResponseWriter,
	status int,
	page string,
	data *templateData,
) {
	ts, ok := app.templateCache[page] // get template for a page from cache
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, err)
		return
	}
	buf := new(bytes.Buffer)
	err := ts.ExecuteTemplate(buf, "base", data) // catch runtime errors by writing to buf
	if err != nil {
		app.serverError(w, err)
		return
	}
	w.WriteHeader(status)
	buf.WriteTo(w)
}

func (app *application) newTemplateData(r *http.Request) *templateData {
	return &templateData{
		CurrentYear:     time.Now().Year(),
		Flash:           app.sessionManager.PopString(r.Context(), "flash"),
		IsAuthenticated: app.isAuthenticated(r),
		CSRFToken:       nosurf.Token(r),
	}
}

// The serverError helper writes an error message and stack trace to the
// errorLog, then sends a generic 500 Internal Server Error response to the
// user.
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(
		w,
		http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError,
	)
}

// The clientError helper sends a specific status code and corresponding
// description to the user.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// This is simply a convenience wrapper around clientError which sends a 404
// Not Found response to the user.
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

// Decode form data into the destination
func (app *application) decodePostForm(r *http.Request, dst any) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	err = app.formDecoder.Decode(dst, r.PostForm)
	if err != nil {
		// if the destination is invalid (e.g. it's a nil pointer), then panic
		var invalidDecodeError *form.InvalidDecoderError
		if errors.As(err, &invalidDecodeError) {
			panic(err)
		}
		return err
	}
	return nil
}

func (app *application) isAuthenticated(r *http.Request) bool {
	return app.sessionManager.Exists(r.Context(), "authenticatedUserID")
}
