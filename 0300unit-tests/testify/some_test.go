package testify

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDemo(t *testing.T) {
	assert.Equal(t, 5, 5, "shouldEqual")
}
