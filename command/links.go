//
// Convert relative links in a file to absolute links pointing to the given domain name
//
package command

import (
	"io/ioutil"
	"os"
	"regexp"
)

type ReplaceLinkCommand struct {
	source     string
	sitedomain string
}

func NewReplaceLinkCommand(source string, sitedomain string) (c *ReplaceLinkCommand) {
	c = new(ReplaceLinkCommand)
	c.source = source
	c.sitedomain = sitedomain
	return
}

func (c *ReplaceLinkCommand) Exec() (err error) {
	var sourcefile *os.File
	sourcefile, err = os.Open(c.source)
	if err != nil {
		return
	}
	defer sourcefile.Close()

	var data []byte
	data, err = ioutil.ReadAll(sourcefile)
	if err != nil {
		return
	}
	sourcefile.Close()

	// href, src
	var hrefSrcMatch = regexp.MustCompile(`((?i)(href|SRC)=["']?\.?)(/[^/])`)
	data = hrefSrcMatch.ReplaceAll(data, []byte("$1//"+c.sitedomain+"$3"))
	// css url
	var urlMatch = regexp.MustCompile(`url\(["']?(/[^/][^"'\)]+)["']?\)`)
	data = urlMatch.ReplaceAll(data, []byte(`url("//`+c.sitedomain+`$1")`))

	err = ioutil.WriteFile(c.source, data, 0644)
	if err != nil {
		return
	}

	return
}
