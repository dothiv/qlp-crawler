//
// Tests for LoadUri
//
// Copyright 2014 TLD dotHIV Registry GmbH.
// @author Markus Tacker <m@dotHIV.org>
//
package util

import (
	"fmt"
	assert "github.com/dothiv/translations-updater/testing"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOpenFile(t *testing.T) {
	_, err := LoadUri("./example/charity.csv")
	assert.NotNil(t, err, "error")
}

func TestOpenURL(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "some csv")
	}))
	defer ts.Close()

	f, err := LoadUri(ts.URL)
	assert.Nil(t, err, "error")
	b, _ := ioutil.ReadAll(f)
	assert.Equals(t, "some csv", string(b))
}
