package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func readFile(t *testing.T, path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal("could not read file", err)
	}
	return string(content)
}

type ExampleTestSuite struct {
	suite.Suite
}

func (suite *ExampleTestSuite) TestExample() {
	assert.Equal(suite.T(), true, true)
}
