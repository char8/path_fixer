package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRankForPath(t *testing.T) {
	var pathRankTests = []struct {
		inputPath    string
		expectedRank int
	}{
		{"/home/char8/bin", -1},
		{"/home/char8/foo/bar/baz", -1},
		{"/homer/foo/baz", 0},
		{"/Users/charith/bin", -2},
		{"/Users/someone/bin/", -2},
		{"/usr/bin/", 2},
		{"/bin/", 3},
	}

	for _, test := range pathRankTests {
		rank := getRankForPath(test.inputPath)
		assert.Equal(t, test.expectedRank, rank)
	}
}

func TestGetOrderedPathsEmpty(t *testing.T) {
	result, err := GetOrderedPaths([]string{})
	assert.Nil(t, err)
	assert.Len(t, result, 0)
}

func TestGetOrderedPathsSimpleDeduplicates(t *testing.T) {
	paths := []string{"/usr/bin/test", "/home/char8/bin/test", "/usr/bin/test"}
	expected := []string{"/home/char8/bin/test", "/usr/bin/test"}
	ranked, err := GetOrderedPaths(paths)

	assert.Nil(t, err)
	assert.ElementsMatch(t, expected, ranked)
}

func TestGetOrderedPaths(t *testing.T) {
	paths := []string{
		"/usr/bin/site_perl",
		"/home/char8/bin/",
		"/usr/local/sbin",
		"/home/char8/.local/bin",
		"/usr/bin/core_perl",
		"/home/char8/bin/",
		"/usr/local/bin",
		"/home/char8/.local/bin",
		"/usr/bin/vendor_perl",
		"/usr/bin",
	}

	ranked, err := GetOrderedPaths(paths)
	assert.Nil(t, err)

	expected := []string{
		"/home/char8/.local/bin",
		"/home/char8/bin",
		"/usr/local/bin",
		"/usr/local/sbin",
		"/usr/bin",
		"/usr/bin/core_perl",
		"/usr/bin/site_perl",
		"/usr/bin/vendor_perl",
	}
	assert.ElementsMatch(t, expected, ranked)
}
