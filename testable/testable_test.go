package testable_test

import (
	"testable"
	"testing"
)

func TestStrInSlice(t *testing.T) {
	var tests = []struct {
		slice []string
		find  string
		want  bool
	}{
		{[]string{"a", "b"}, "c", false},
		{[]string{"a", "b"}, "b", true},
	}
	for _, tt := range tests {
		t.Run(tt.find, func(t *testing.T) {
			got := testable.StrInSlice(tt.slice, tt.find)
			if got != tt.want {
				t.Errorf("expecting %t, got %t", tt.want, got)
			}
		})
	}
}

func TestGetAverageStarsPerRepoUnit(t *testing.T) {
	var tests = []struct {
		username string
		want     float64
	}{
		{"muhammedsaidkaya", 3.5},
	}
	mockGithub := new(testable.MockGithub)
	for _, tt := range tests {
		t.Run(tt.username, func(t *testing.T) {
			got, err := testable.GetAverageStarsPerRepo(mockGithub, tt.username)
			if err != nil {
				t.Errorf("expecting nil err, got %v", err)
			}
			if got != tt.want {
				t.Errorf("expecting %f, got %f", tt.want, got)
			}
		})
	}
}

func TestGetAverageStarsPerRepoIntegration(t *testing.T) {
	var tests = []struct {
		username string
		want     float64
	}{
		{"muhammedsaidkaya", 0.18518518518518517},
	}
	github := new(testable.Github)
	for _, tt := range tests {
		t.Run(tt.username, func(t *testing.T) {
			got, err := testable.GetAverageStarsPerRepo(github, tt.username)
			if err != nil {
				t.Errorf("expecting nil err, got %v", err)
			}
			if got != tt.want {
				t.Errorf("expecting %f, got %f", tt.want, got)
			}
		})
	}
}
