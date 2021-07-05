package kangaroo

import (
	"testing"

	u "github.com/geoirb/kangaroo/pkg/utils"

	"github.com/stretchr/testify/assert"
)

type testAlongWayCase struct {
	name           string
	x              int
	v              int
	coordinate     int
	expectedResult bool
}

var testAlongWaySet = []testAlongWayCase{
	{
		"Case 1",
		3802,
		3802,
		3802,
		true,
	},
	{
		"Case 2",
		3802,
		3802,
		48298,
		true,
	},

	{
		"Case 3",
		3802,
		-3802,
		-48298,
		true,
	},
	{
		"Case 4",
		3802,
		3802,
		-1,
		false,
	},
	{
		"Case 5",
		3802,
		-3802,
		48298,
		false,
	},
}

func TestAlongWay(t *testing.T) {
	for _, tCase := range testAlongWaySet {
		t.Run(tCase.name, func(t *testing.T) {
			k := &Kangaroo{
				x: tCase.x,
				v: tCase.v,
			}
			actualResult := k.alongWay(tCase.coordinate)
			assert.Equal(t, tCase.expectedResult, actualResult)
		})
	}
}

type testIntersectionPointCase struct {
	name               string
	x1                 int
	v1                 int
	x2                 int
	v2                 int
	expectedStep       *int
	expectedCoordinate *int
}

var (
	testIntersectionPointSet = []testIntersectionPointCase{
		{
			"Case 1",
			3802,
			3802,
			3802,
			3802,
			nil,
			allCoordinate,
		},
		{
			"Case 2",
			3802,
			3802,
			3,
			3802,
			nil,
			nil,
		},
		{
			"Case 3",
			0,
			3,
			0,
			2,
			&testCase3Step,
			&testCase3Coordinate,
		},
		{
			"Case 4",
			4,
			10,
			1,
			8,
			nil,
			nil,
		},
		{
			"Case 5",
			4,
			10,
			-11,
			5,
			nil,
			nil,
		},
		{
			"Case 6",
			4,
			10,
			-11,
			5,
			nil,
			nil,
		},
		{
			"Case 7",
			4,
			10,
			14,
			5,
			&testCase7Step,
			&testCase7Coordinate,
		},
	}

	testCase3Step       = 0
	testCase3Coordinate = 0

	testCase7Step       = 2
	testCase7Coordinate = 24
)

func TestIntersectionPoint(t *testing.T) {
	u := u.NewNumber(1e-9)
	fabric := NewKangarooFabric(u)

	for _, tCase := range testIntersectionPointSet {
		t.Run(tCase.name, func(t *testing.T) {
			first := fabric(tCase.x1, tCase.v1)
			second := fabric(tCase.x2, tCase.v2)

			actualStep, actualCoordinate := first.getIntersectionPoint(second)
			if actualStep == nil {
				assert.Equal(t, tCase.expectedStep, actualStep, "step")
				assert.Equal(t, tCase.expectedCoordinate, actualCoordinate, "coordinate")
			} else {
				assert.NotNil(t, actualCoordinate)
				assert.Equal(t, *tCase.expectedStep, *actualStep, "step")
				assert.Equal(t, *tCase.expectedCoordinate, *actualCoordinate, "coordinate")
			}
		})
	}
}

type testIsIntersectCase struct {
	name           string
	x1             int
	v1             int
	x2             int
	v2             int
	expectedResult bool
}

var testIsIntersectSet = []testIsIntersectCase{
	{
		"Case 1",
		4,
		10,
		-11,
		5,
		false,
	},
	{
		"Case2",
		0,
		3,
		4,
		2,
		true,
	},
	{
		"Case 3",
		4,
		10,
		14,
		5,
		true,
	},
	{
		"Case 4",
		0,
		1,
		0,
		2,
		true,
	},
}

func TestIntersect(t *testing.T) {
	u := u.NewNumber(1e-9)
	fabric := NewKangarooFabric(u)

	for _, tCase := range testIsIntersectSet {
		t.Run(tCase.name, func(t *testing.T) {
			first := fabric(tCase.x1, tCase.v1)
			second := fabric(tCase.x2, tCase.v2)

			actualResult := first.IsIntersect(second)
			assert.Equal(t, tCase.expectedResult, actualResult)
		})
	}
}
