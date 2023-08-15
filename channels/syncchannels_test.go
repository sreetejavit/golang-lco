package main

import "testing"

func Test_routine(t *testing.T) {
	type args struct {
		wg *sync.WaitGroup
	}
	tests := []struct {
		name string
		args args
	}{
		&assert.True() // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
