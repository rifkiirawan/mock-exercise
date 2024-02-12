package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	result := Sum(10, 10)

	// if result != 20 {
	// 	panic("Result should be 20")
	// }

	//testify
	assert.Equal(t, 20, result, "result has to be 20")

	fmt.Println("testSum Eksekusi Terhenti")
}

func TestTableSum(t *testing.T) {
	tests := []struct {
		request  int
		expected int
		errMsg   string
	}{
		{
			request:  Sum(10, 10),
			expected: 20,
			errMsg:   "result has to be 20",
		},

		{
			request:  Sum(20, 20),
			expected: 40,
			errMsg:   "result has to be 40",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			require.Equal(t, test.expected, test.request, test.errMsg)
		})
	}
}
