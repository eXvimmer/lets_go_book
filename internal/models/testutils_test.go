package models

import (
	"database/sql"
	"io/ioutil"
	"testing"
)

func newTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open(
		"mysql",
		"test_web:pass@/test_snippetbox?parseTime=true&multiStatements=true",
	)
	if err != nil {
		t.Fatal(err)
	}
	script, err := ioutil.ReadFile("./testdata/setup.sql")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(string(script))
	if err != nil {
		t.Fatal()
	}
	// Use the t.Cleanup() to register a function *which will automatically be
	// called by Go when the current test (or sub-test) which calls newTestDB()
	// has finished*.
	t.Cleanup(func() {
		script, err := ioutil.ReadFile("./testdata/teardow.sql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = db.Exec(string(script))
		if err != nil {
			t.Fatal(err)
		}
		db.Close()
	})

	return db
}
