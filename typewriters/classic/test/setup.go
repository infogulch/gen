// Run this before testing: go setup.go

package main

import (
	"os"
	"strings"

	"github.com/clipperhouse/gen/typewriter"
	_ "github.com/clipperhouse/gen/typewriters/classic"
)

func main() {
	// don't let bad test or gen files get us stuck
	filter := func(f os.FileInfo) bool {
		return !strings.HasSuffix(f.Name(), "_test.go") && !strings.HasSuffix(f.Name(), "_gen.go")
	}

	a, err := typewriter.NewAppFiltered("+test", filter)
	if err != nil {
		panic(err)
	}
	a.WriteAll()
}
