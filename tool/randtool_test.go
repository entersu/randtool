package tool

import (
	"testing"
)

func TestRandNumber(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
		{
			"code1",
			1,
		}, {
			"code2",
			2,
		}, {
			"code3",
			3,
		}, {
			"code4",
			4,
		},{
			"code5",
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandNumber(); got != tt.want {
				t.Errorf("ReadNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
