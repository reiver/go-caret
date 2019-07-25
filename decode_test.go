package caret

import (
	"github.com/reiver/go-utf8s"

	"testing"
)

func TestDecode(t *testing.T) {

	tests := []struct{
		Code     rune
		Expected rune
	}{
		{
			Code:     '@',
			Expected: 0x00,
		},
		{
			Code:     'A',
			Expected: 0x01,
		},
		{
			Code:     'B',
			Expected: 0x02,
		},
		{
			Code:     'C',
			Expected: 0x03,
		},
		{
			Code:     'D',
			Expected: 0x04,
		},
		{
			Code:     'E',
			Expected: 0x05,
		},
		{
			Code:     'F',
			Expected: 0x06,
		},
		{
			Code:     'G',
			Expected: '\a',
		},
		{
			Code:     'H',
			Expected: '\b',
		},
		{
			Code:     'I',
			Expected: '\t',
		},
		{
			Code:     'J',
			Expected: '\n',
		},
		{
			Code:     'K',
			Expected: '\v',
		},
		{
			Code:     'L',
			Expected: '\f',
		},
		{
			Code:     'M',
			Expected: '\r',
		},
		{
			Code:     'N',
			Expected: 0x0e,
		},
		{
			Code:     'O',
			Expected: 0x0f,
		},
		{
			Code:     'P',
			Expected: 0x10,
		},
		{
			Code:     'Q',
			Expected: 0x11,
		},
		{
			Code:     'R',
			Expected: 0x12,
		},
		{
			Code:     'S',
			Expected: 0x13,
		},
		{
			Code:     'T',
			Expected: 0x14,
		},
		{
			Code:     'U',
			Expected: 0x15,
		},
		{
			Code:     'V',
			Expected: 0x16,
		},
		{
			Code:     'W',
			Expected: 0x17,
		},
		{
			Code:     'X',
			Expected: 0x18,
		},
		{
			Code:     'Y',
			Expected: 0x19,
		},
		{
			Code:     'Z',
			Expected: 0x1A,
		},
		{
			Code:     '[',
			Expected: 0x1B,
		},
		{
			Code:     '\\',
			Expected: 0x1C,
		},
		{
			Code:     ']',
			Expected: 0x1D,
		},
		{
			Code:     '^',
			Expected: 0x1E,
		},
		{
			Code:     '_',
			Expected: 0x1F,
		},
		{
			Code:     '?',
			Expected: 0x7F,
		},
	}

	for testNumber, test := range tests {

		decoded := decode(test.Code)

		if expected, actual := test.Expected, decoded; expected != actual {
			t.Errorf("For test #%d, what was decoded was not what was expected", testNumber)
			t.Logf("ORIGINAL: %s (%d) %q", utf8s.FormatBinary(test.Code),   test.Code, test.Code)
			t.Logf("EXPECTED: %s (%d) %q", utf8s.FormatBinary(expected), expected,  expected)
			t.Logf("ACTUAL:   %s (%d) %q", utf8s.FormatBinary(actual),   actual,    actual)
			continue
		}
	}
}
