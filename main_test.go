package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddition(t *testing.T) {
	assert.Equal(t, 6, calculate(2, 4, Addition))
	assert.Equal(t, 10.0, calculate(4.0, 6.0, Addition))
}
