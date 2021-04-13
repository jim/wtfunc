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
			Output: "main\n",
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
			Output: "main\n",
		},
		{
			Name:   "method with value receiver",
			Args:   []string{"-line=12", "testdata/code.go"},
			Output: "valueReceiver\n",
		},
		{
			Name:   "method with pointer receiver",
			Args:   []string{"-line=16", "testdata/code.go"},
			Output: "pointerReceiver\n",
		},
		{
			Name:   "list all functions in file",
			Args:   []string{"-list", "testdata/code.go"},
			Output: "main\nvalueReceiver\npointerReceiver\n",
		},
		{
			Name:   "list all test functions in file",
			Args:   []string{"-test", "-list", "testdata/tests.go"},
			Output: "TestExample\n",
		},
		{
			Name:   "list all testfy methods in file",
			Args:   []string{"-test", "-list", "testdata/testify_tests.go"},
			Output: "TestExample\n",
		},
		{
			Name:   "find test on a line",
			Args:   []string{"-test", "-line=17", "testdata/tests.go"},
			Output: "TestExample\n",
		},
		{
			Name:   "find no test on a line",
			Args:   []string{"-test", "-line=10", "testdata/tests.go"},
			Output: "",
			Status: 2,
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
