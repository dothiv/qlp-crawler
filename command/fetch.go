//
// Fetch a file from url, filename and save it to a local file
//
package command

import (
	"github.com/dothiv/qlp-crawler/util"
	"io/ioutil"
	"os"
)

type FetchCommand struct {
	source string
	target string
}

func NewFetchCommand(source string, target string) (c *FetchCommand) {
	c = new(FetchCommand)
	c.source = source
	c.target = target
	return
}

func (c *FetchCommand) Exec() (err error) {
	var sourcefile *os.File
	sourcefile, err = util.LoadUri(c.source)
	if err != nil {
		return
	}
	defer sourcefile.Close()

	var data []byte
	data, err = ioutil.ReadAll(sourcefile)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(c.target, data, 0644)
	if err != nil {
		return
	}

	return
}
