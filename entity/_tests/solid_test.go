package _tests

import (
	"github.com/scantrust/dxf-golang/entity"
	"github.com/scantrust/dxf-golang/format"
	"testing"
)

func TestFloatFormatting(t *testing.T) {
	testCases := map[float64]string{
		12:       "12",
		12.1:     "12.1",
		12.12:    "12.12",
		12.123:   "12.123",
		12.1234:  "12.123",
		12.1236:  "12.124",
		8.309999: "8.31",
	}

	for input, expected := range testCases {
		actual := format.FormatFloat(input, 3)
		if actual != expected {
			t.Errorf("Input: %.10f, expected %s, got %s", input, expected, actual)
		}
	}

	// this is not equal to 8.31 due to float addition, but 8.309999999999
	value := 8.30999999999
	actual := format.FormatFloat(value, 6)
	if "8.31" != actual {
		t.Errorf("Expected 8.31, got %s", actual)
	}
}

func TestSolidRectFormat(t *testing.T) {
	rect := entity.NewRect2D(5.31, 7.12, 3, 5)
	actual := rect.String()
	expected := `0
SOLID
5
0
100
AcDbEntity
8
0
100
AcDbTrace
10
5.31
20
7.12
30
0
11
8.31
21
7.12
31
0
12
5.31
22
12.12
32
0
13
8.31
23
12.12
33
0`
	if expected != actual {
		t.Errorf("Invalid output: %s", actual)
	}
}
