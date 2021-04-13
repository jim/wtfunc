package main

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

func readFile(t *testing.T, path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal("could not read file", err)
	}
	return string(content)
}

func TestCLI(t *testing.T) {
	for _, test := range []struct {
		Name   string
		Args   []string
		In     string
		Output string
		Err    string
		Status int
	}{
		{
			Name:   "happy path",
			Args:   []string{"-line=6", "testdata/code.go"},
			Output: "main",
			Status: 0,
		},
		{
			Name:   "no matching func",
			Args:   []string{"-line=4", "testdata/code.go"},
			Status: 2,
		},
		{
			Name:   "reading from standard in",
			Args:   []string{"-line=6"},
			In:     readFile(t, "testdata/code.go"),
			Output: "main",
			Status: 0,
		},
		{
			Name:   "method with value receiver",
			Args:   []string{"-line=12", "testdata/code.go"},
			Output: "valueReceiver",
			Status: 0,
		},
		{
			Name:   "method with pointer receiver",
			Args:   []string{"-line=16", "testdata/code.go"},
			Output: "pointerReceiver",
			Status: 0,
		},
		{
			Name:   "list all functions in file",
			Args:   []string{"-list", "testdata/code.go"},
			Output: "main\nvalueReceiver\npointerReceiver\n",
			Status: 0,
		},
	} {
		t.Run(strings.Join(test.Args, ""), func(t *testing.T) {
			in := bytes.NewBufferString(test.In)
			out := new(bytes.Buffer)
			err := new(bytes.Buffer)

			returnValue := run(in, out, err, test.Args)
			if returnValue != test.Status {
				t.Errorf("expected status %d, but got %d", test.Status, returnValue)
			}

			if output := out.String(); output != test.Output {
				t.Errorf("expected output %s, but got %s", test.Output, output)
			}
		})
	}
}
