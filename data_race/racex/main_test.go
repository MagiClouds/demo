package main

import (
	"testing"
)

func Test_dataRace(t *testing.T) {
	tests := []struct {
		name string
		Want int
	}{
		{"A", 1000000}, {"B", 222}, {"C", 333}, {"D", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataRace()
		})
	}
}

func TestPrint(t *testing.T)  {


}