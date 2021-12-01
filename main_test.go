package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	arg1, arg2, expected string
}

var testcases = []testCase{
	{"53QQ2", "Q53Q2", "Tie"},
	{"53888", "88385", "Tie"},
	{"AAAQQ", "QQAAA", "Tie"},
	{"QQAAA", "AAAQQ", "Tie"},
	{"Q53Q2", "53QQ2", "Tie"},
	{"88385", "53888", "Tie"},
	{"AAAQQ", "QQQAA", "Hand 1"},
	{"Q53Q4", "53QQ2", "Hand 1"},
	{"53888", "88375", "Hand 1"},
	{"33337", "QQAAA", "Hand 1"},
	{"22333", "AAA58", "Hand 1"},
	{"33389", "AAKK4", "Hand 1"},
	{"44223", "AA892", "Hand 1"},
	{"22456", "AKQJT", "Hand 1"},
	{"99977", "77799", "Hand 1"},
	{"99922", "88866", "Hand 1"},
	{"9922A", "9922K", "Hand 1"},
	{"99975", "99965", "Hand 1"},
	{"99975", "99974", "Hand 1"},
	{"99752", "99652", "Hand 1"},
	{"99752", "99742", "Hand 1"},
	{"99753", "99752", "Hand 1"},
	{"QQQAA", "AAAQQ", "Hand 2"},
	{"53QQ2", "Q53Q4", "Hand 2"},
	{"88375", "53888", "Hand 2"},
	{"QQAAA", "33337", "Hand 2"},
	{"AAA58", "22333", "Hand 2"},
	{"AAKK4", "33389", "Hand 2"},
	{"AA892", "44223", "Hand 2"},
	{"AKQJT", "22456", "Hand 2"},
	{"77799", "99977", "Hand 2"},
	{"88866", "99922", "Hand 2"},
	{"9922K", "9922A", "Hand 2"},
	{"99965", "99975", "Hand 2"},
	{"99974", "99975", "Hand 2"},
	{"99652", "99752", "Hand 2"},
	{"99742", "99752", "Hand 2"},
	{"99752", "99753", "Hand 2"},
}

func TestEvaluate(t *testing.T) {
	for _, test := range testcases {
		t.Logf("Test %q - %q", test.arg1, test.arg2)
		output, _ := evaluate(test.arg1, test.arg2)
		assert.Equal(t, test.expected, output)
	}
}
