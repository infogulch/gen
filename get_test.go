package main

import (
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	// just run it; not a great test but at least covers the code
	if err := get([]string{}); err != nil {
		t.Error(err)
	}

	if err := get([]string{"-d"}); err != nil {
		t.Error(err)
	}
}

func TestGetImports(t *testing.T) {
	// use custom name so test won't interfere with a real _gen.go
	setCustomName("_gen_test.go")
	defer revertCustomName()

	// remove existing files, start fresh
	os.Remove(customName)

	// standard
	imps, err := getTypewriterImports()

	if err != nil {
		t.Error(err)
	}

	if len(imps) != 2 {
		t.Errorf("should return 2 imports, got %v", len(imps))
	}

	// create a custom typewriter import file
	w, err := os.Create(customName)

	if err != nil {
		t.Error(err)
	}

	defer os.Remove(customName)

	p := pkg{
		Name: "main",
		Imports: []string{
			// non-standard typewriter
			`_ "github.com/clipperhouse/gen/typewriters/classic"`,
			`_ "github.com/clipperhouse/gen/typewriters/container"`,
			`_ "github.com/clipperhouse/gen/typewriters/foo"`,
		},
	}

	if err := tmpl.Execute(w, p); err != nil {
		t.Error(err)
	}

	// custom file now exists
	imps2, err := getTypewriterImports()

	if err != nil {
		t.Error(err)
	}

	if len(imps2) != 3 {
		t.Errorf("should return 3 custom imports, got %v", len(imps2))
	}

	// custom get
	if err := get([]string{}); err != nil {
		t.Error(err)
	}
}
