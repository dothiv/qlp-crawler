//
// Add the clickcounter code to the html file
//
package command

import (
	"io/ioutil"
	"os"
	"regexp"
)

type AddClickcounterCommand struct {
	source string
}

func NewAddClickcounterCommand(source string) (c *AddClickcounterCommand) {
	c = new(AddClickcounterCommand)
	c.source = source
	return
}

func (c *AddClickcounterCommand) Exec() (err error) {
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

	var hrefMatch = regexp.MustCompile(`</body>`)
	data = hrefMatch.ReplaceAll(data, []byte(`<script src="//dothiv-registry.appspot.com/static/clickcounter.min.js" type="text/javascript"></script>`+"\n"+`$0`))

	err = ioutil.WriteFile(c.source, data, 0644)
	if err != nil {
		return
	}

	return
}
