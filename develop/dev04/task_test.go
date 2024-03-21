package main

import (
	"reflect"
	"testing"
)

func TestAnagramIdentification(t *testing.T) {
	testCases := []struct {
		words    []string
		expected map[string][]string
	}{
		{
			words: []string{"state", "taste", "listen", "silent", "below", "elbow"},
			expected: map[string][]string{
				"below":  {"below", "elbow"},
				"listen": {"listen", "silent"},
				"state":  {"state", "taste"},
			},
		},
	}

	for _, tc := range testCases {
		actual := findAnagrams(tc.words)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("Expected %v, but got %v", tc.expected, actual)
		}
	}
}
