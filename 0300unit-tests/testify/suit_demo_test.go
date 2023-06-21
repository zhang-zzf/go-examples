package testify

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// 面向对象的思想 writing test code
type SuitTestDemo struct {
	suite.Suite
	// some obj that will use in the test code
	aVar int
}

// set up the obj before each test case in this suit
func (ts *SuitTestDemo) SetupTest() {
	ts.aVar = 5
}

// reset the obj after each test case in this suit
func (ts *SuitTestDemo) TearDownTest() {
	ts.aVar = 0
}

// func start with Test is a test case
func (ts *SuitTestDemo) TestDemo1() {
	assert.Equal(ts.T(), ts.aVar, 5)
}

func TestThisSuit(t *testing.T) {
	suite.Run(t, new(SuitTestDemo))
}
