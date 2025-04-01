package main

import (
	"testing"
)

func TestProcessECGData(t *testing.T) {
	tests := []ECGData{
		{HeartRate: 120, RRInterval: 0.75},
		{HeartRate: 50, RRInterval: 0.90},
		{HeartRate: 85, RRInterval: 0.40},
		{HeartRate: 70, RRInterval: 0.75},
	}

	for _, test := range tests {
		processECGData(test)
	}
}
