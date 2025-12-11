package main

import (
	"testing"
)

func TestPt1(t *testing.T) {
	cases := []struct {
		caseName    string
		fileName    string
		expectedAns int
	}{
		{
			caseName:    "sample",
			fileName:    "01_sample.txt",
			expectedAns: 3,
		},
	}
	for _, tt := range cases {
		t.Run(tt.caseName, func(t *testing.T) {
			output, err := p1(tt.fileName)
			if err != nil {
				t.Errorf("Encountered error: %v", err)
			}
			if output != tt.expectedAns {
				t.Errorf("Expected %v, got %v", tt.expectedAns, output)
			}
		},
		)
	}
}

func TestPt2(t *testing.T) {
	cases := []struct {
		caseName    string
		fileName    string
		expectedAns int
	}{
		{
			caseName:    "sample",
			fileName:    "01_sample.txt",
			expectedAns: 6,
		},
	}
	for _, tt := range cases {
		t.Run(tt.caseName, func(t *testing.T) {
			output, err := p2(tt.fileName)
			if err != nil {
				t.Errorf("Encountered error: %v", err)
			}
			if output != tt.expectedAns {
				t.Errorf("Expected %v, got %v", tt.expectedAns, output)
			}
		},
		)
	}
}
