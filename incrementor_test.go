package incrementor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncrementor(t *testing.T) {
	tests := []struct {
		name           string
		incrementor    Incrementor
		expectedValue  uint
		incrementCount uint
		maxValue       uint
	}{
		{
			name:           "default value is zero",
			incrementor:    New(),
			expectedValue:  0,
			incrementCount: 0,
			maxValue:       99,
		},
		{
			name:           "increment is work",
			incrementor:    New(),
			expectedValue:  1,
			incrementCount: 1,
			maxValue:       99,
		},
		{
			name:           "increment fill until number < maximumValue",
			incrementor:    New(),
			expectedValue:  2,
			incrementCount: 2,
			maxValue:       99,
		},
		{
			name:           "increment will refresh after exhausting of limit",
			incrementor:    New(),
			expectedValue:  0,
			incrementCount: 6,
			maxValue:       5,
		},
	}

	for _, tt := range tests {
		tt := tt
		tt.incrementor.SetMaximumValue(tt.maxValue)
		if tt.incrementCount > 0 {
			var i uint
			for i = 0; i < tt.incrementCount; i++ {
				tt.incrementor.IncrementNumber()
			}
		}
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.incrementor.GetNumber()
			assert.Equal(t, tt.expectedValue, actual)
		})
	}
}
