package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	assert := assert.New(t)
	cases := []struct {
		Expected  string
		oriString string
	}{
		{"bab", "babad"},
		{"bb", "cbbd"},
	}
	for _, c := range cases {
		assert.Equal(c.Expected, longestPalindrome(c.oriString))
	}
}
