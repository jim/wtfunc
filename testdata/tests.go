package main

import (
	"io/ioutil"
	"testing"
)

func readFile(t *testing.T, path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal("could not read file", err)
	}
	return string(content)
}

func TestExample(t *testing.T) {
	t.Parallel()
}
