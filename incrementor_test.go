package incrementor

import (
	"github.com/stretchr/testify/assert"
	"testing"
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

	for _, test := range tests {
		test.incrementor.SetMaximumValue(test.maxValue)
		if test.incrementCount > 0 {
			var i uint
			for i = 0; i < test.incrementCount; i++ {
				test.incrementor.IncrementNumber()
			}
		}
		t.Run(test.name, func(t *testing.T) {
			actual := test.incrementor.GetNumber()
			assert.Equal(t, test.expectedValue, actual)
		})
	}
}
