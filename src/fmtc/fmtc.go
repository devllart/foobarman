package fmtc

import (
	"bufio"
	"os"
	"strings"
)

var (
	Reader *bufio.Reader
	PreReplacer *strings.Replacer
	ColorReplacer *strings.Replacer
)

func init() {
	AddManyReplace(manyReplaceMap)

	Reader = bufio.NewReader(os.Stdin)
	PreReplacer = strings.NewReplacer(manyReplace...)
	ColorReplacer = strings.NewReplacer("%R", Red, "%G", Green, "%B", Blue, "%Y", Yellow, "%C", Reset, "%f", "%.3f")
}
