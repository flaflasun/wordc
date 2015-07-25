package cmd

import (
	"log"
	"os"
	"sort"

	"github.com/codegangsta/cli"
	"github.com/flaflasun/wordc/wordcount"
)

func Wordc(c *cli.Context) {
	input, err := getInput(c.Args().Get(0))
	if err != nil {
		log.Fatal(err)
	}

	edit := newEditor(input)
	edit.ignoreChar(c.String("ignorechar"))
	edit.addDelimiter(c.String("delimiter"))
	if c.Bool("ignorecase") {
		edit.ignoreCase()
	}

	words := make(wordcount.Words, 0, 0)
	words.AddWords(edit.text)

	if c.Bool("reverse") {
		sort.Sort(words)
	} else {
		sort.Sort(sort.Reverse(words))
	}
	words.String(os.Stdout)
}
