package main

var stdImports = []string{
	`_ "github.com/clipperhouse/gen/typewriters/classic"`,
	`_ "github.com/clipperhouse/gen/typewriters/container"`,
}

type pkg struct {
	Name    string
	Imports []string
}
