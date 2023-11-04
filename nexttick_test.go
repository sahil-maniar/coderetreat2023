package coderetreat2023

import (
	"reflect"
	"testing"
)

func TestNextTick(t *testing.T) {
	tests := []struct {
		name   string
		input  [][]bool
		output [][]bool
	}{
		{
			name: "Test single row single column alive",
			input: [][]bool{
				{true},
			},
			output: [][]bool{
				{false},
			},
		},
		{
			name: "Test single row single column dead",
			input: [][]bool{
				{false},
			},
			output: [][]bool{
				{false},
			},
		},
		{
			name: "Test two row two column single first dead",
			input: [][]bool{
				{false, true},
				{true, true},
			},
			output: [][]bool{
				{true, false},
				{false, true},
			},
		},
		{
			name: "Test two row two column first row alive",
			input: [][]bool{
				{true, true},
				{false, false},
			},
			output: [][]bool{
				{false, false},
				{false, false},
			},
		},
		{
			name: "Test two row two column first single alive",
			input: [][]bool{
				{true, false},
				{false, false},
			},
			output: [][]bool{
				{false, false},
				{false, false},
			},
		},
		{
			name: "Test two row two column diagonal alive",
			input: [][]bool{
				{true, false},
				{false, true},
			},
			output: [][]bool{
				{false, false},
				{false, false},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := nextTick(tt.input)
			if !reflect.DeepEqual(tt.output, got) {
				t.Errorf("Expected %+v, got %+v", tt.input, got)
			}
		})
	}

}
