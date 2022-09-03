package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddition(t *testing.T) {
	assert.Equal(t, 6, calculate(2, 4, Addition))

}
