//
// qlp-crawler bin
//
// Fetch a QLP page and convert it to a .hiv page
//
// Copyright 2014 TLD dotHIV Registry GmbH.
// @author Markus Tacker <m@click4life.hiv>
//
package main

import (
	"flag"
	"fmt"
	"github.com/dothiv/qlp-crawler/command"
	"os"
)

func main() {
	source := flag.String("source", "", "source url or file")
	target := flag.String("target", "", "target file")
	sitedomain := flag.String("sitedomain", "", "site domain")
	flag.Parse()

	if len(*source) == 0 {
		os.Stderr.WriteString("source is required\n")
		flag.Usage()
		os.Exit(1)
	}

	if len(*target) == 0 {
		os.Stderr.WriteString("target is required\n")
		flag.Usage()
		os.Exit(1)
	}

	if len(*sitedomain) == 0 {
		os.Stderr.WriteString("sitedomain is required\n")
		flag.Usage()
		os.Exit(1)
	}

	os.Stdout.WriteString(fmt.Sprintf("Converting %s ...\n", *source))

	c := command.NewFetchCommand(*source, *target)
	err := c.Exec()
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}

	c2 := command.NewReplaceLinkCommand(*target, *sitedomain)
	err = c2.Exec()
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}

	c3 := command.NewAddClickcounterCommand(*target)
	err = c3.Exec()
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}

	os.Stdout.WriteString(fmt.Sprintf("%s written.\n", *target))
	return
}
