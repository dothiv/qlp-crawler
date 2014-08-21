//
// Helper to load files either from the filesystem or from URLs
//
// Copyright 2014 TLD dotHIV Registry GmbH.
// @author Markus Tacker <m@dotHIV.org>
//
package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Open a file, can be from an URL too.
func LoadUri(uri string) (l *os.File, err error) {
	if uri[0:4] == "http" {
		return loadHttp(uri)
	}
	return openFile(uri)
}

// opens a local file
func openFile(uri string) (l *os.File, err error) {
	l, err = os.Open(uri)
	return
}

// fetches an URL and saves it as a temp file
// then opens it
func loadHttp(uri string) (l *os.File, err error) {
	var response *http.Response
	response, err = http.Get(uri)
	if err != nil {
		print(1)
		return
	}
	var body []byte
	body, err = ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return
	}

	if response.StatusCode != 200 {
		err = fmt.Errorf("Failed to load '%s': %s", uri, body)
		return
	}

	var tmpFile *os.File
	tmpFile, err = ioutil.TempFile(os.TempDir(), "translations-updater")
	if err != nil {
		return
	}
	defer tmpFile.Close()

	tmpFile.Write(body)

	return openFile(tmpFile.Name())
}
