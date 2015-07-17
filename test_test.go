package foo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPass(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(1, 1)
}

func TestFail(t *testing.T) {
	assert := assert.New(t)
	assert.NotEqual(1, 1)
}
