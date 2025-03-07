package main

import (
	"reflect"
	"testing"
)

func TestSortPages(t *testing.T) {
	tests := []struct {
		name      string
		inputUrls map[string]int
		expected  []string
	}{
		{
			name:      "test 1",
			inputUrls: map[string]int{},
			expected:  []string{},
		},
		{
			name: "test 2",
			inputUrls: map[string]int{
				"www.test1.com": 1,
				"www.test2.com": 2,
				"www.test3.com": 3,
				"www.test4.com": 4,
				"www.test5.com": 5,
			},
			expected: []string{
				"www.test5.com",
				"www.test4.com",
				"www.test3.com",
				"www.test2.com",
				"www.test1.com",
			},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := sortPages(tc.inputUrls)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected : %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
