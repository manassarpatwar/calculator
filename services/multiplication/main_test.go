package multiplication

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	assert.Equal(t, 24.0, calculate(4.0, 6.0))
}
