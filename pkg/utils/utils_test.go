package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name           string
	epsilon        float64
	value          float64
	expectedResult bool
}

var testSet = []testCase{
	{
		"Case 1",
		1e-9,
		3802.0,
		true,
	},
	{
		"Case 2",
		1e-9,
		38.02,
		false,
	},
	{
		"Case 3",
		1e-9,
		3802e-9,
		false,
	},
	{
		"Case 4",
		1e-9,
		3802e-100,
		true,
	},
	{
		"Case 5",
		1e-9,
		0,
		true,
	},
	{
		"Case 6",
		1e-9,
		-3802e-9,
		false,
	},
	{
		"Case 7",
		1e-9,
		-38.02,
		false,
	},
	{
		"Case 8",
		1e-9,
		-3802.0,
		true,
	},
}

func TestIsFloatInt(t *testing.T) {
	for _, tCase := range testSet {
		t.Run(tCase.name, func(t *testing.T) {
			number := NewNumber(tCase.epsilon)
			actualResult := number.IsFloatInt(tCase.value)
			assert.Equal(t, tCase.expectedResult, actualResult)
		})
	}
}
