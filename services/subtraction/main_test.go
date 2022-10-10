package subtraction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubtraction(t *testing.T) {
	assert.Equal(t, -2, calculate(2, 4))
}
