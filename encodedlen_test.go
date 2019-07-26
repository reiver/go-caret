package caret

import (
	"testing"
)

func TestEncodedlen(t *testing.T) {

	tests := []struct{
		Encoded   string
		DecodeLen int
		ExpectedEncodeLen int
	}{
		{
			Encoded: "",
			DecodeLen:         0,
			ExpectedEncodeLen: 0,
		},



		{
			Encoded: "a",
			DecodeLen:         1,
			ExpectedEncodeLen: 1,
		},
		{
			Encoded: "b",
			DecodeLen:         1,
			ExpectedEncodeLen: 1,
		},
		{
			Encoded: "c",
			DecodeLen:         1,
			ExpectedEncodeLen: 1,
		},
		{
			Encoded: "d",
			DecodeLen:         1,
			ExpectedEncodeLen: 1,
		},



		{
			Encoded: "Hello world!",
			DecodeLen:         0,
			ExpectedEncodeLen: 0,
		},
		{
			Encoded: "Hello world!",
			DecodeLen:         1,
			ExpectedEncodeLen: 1,
		},
		{
			Encoded: "Hello world!",
			DecodeLen:         2,
			ExpectedEncodeLen: 2,
		},
		{
			Encoded: "Hello world!",
			DecodeLen:         3,
			ExpectedEncodeLen: 3,
		},
		{
			Encoded: "Hello world!",
			DecodeLen:         4,
			ExpectedEncodeLen: 4,
		},
		{
			Encoded: "Hello world!",
			DecodeLen:         5,
			ExpectedEncodeLen: 5,
		},
		{
			Encoded: "Hello world!",
			DecodeLen:         6,
			ExpectedEncodeLen: 6,
		},
		{
			Encoded: "Hello world!",
			DecodeLen:         7,
			ExpectedEncodeLen: 7,
		},
		{
			Encoded: "Hello world!",
			DecodeLen:         8,
			ExpectedEncodeLen: 8,
		},
		{
			Encoded: "Hello world!",
			DecodeLen:         9,
			ExpectedEncodeLen: 9,
		},
		{
			Encoded: "Hello world!",
			DecodeLen:         10,
			ExpectedEncodeLen: 10,
		},
		{
			Encoded: "Hello world!",
			DecodeLen:         11,
			ExpectedEncodeLen: 11,
		},
		{
			Encoded: "Hello world!",
			DecodeLen:         12,
			ExpectedEncodeLen: 12,
		},



		{
			Encoded: "The "+ "^["   +"[1m" +"bold"+ "^[" +"[0m"+" text.",
			//       ""
			DecodeLen:         0,
			ExpectedEncodeLen: 0,
		},
		{
			Encoded: "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0m"+" text.",
			//       "T"
			DecodeLen:         1,
			ExpectedEncodeLen: 1,
		},
		{
			Encoded: "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0m"+" text.",
			//       "Th"
			DecodeLen:         2,
			ExpectedEncodeLen: 2,
		},
		{
			Encoded: "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0m"+" text.",
			//       "The"
			DecodeLen:         3,
			ExpectedEncodeLen: 3,
		},
		{
			Encoded: "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0m"+" text.",
			//       "The "
			DecodeLen:         4,
			ExpectedEncodeLen: 4,
		},
		{
			Encoded: "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0m"+" text.",
			//       "The "+ "^["
			DecodeLen:         5,
			ExpectedEncodeLen: 6,
		},
		{
			Encoded: "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0m"+" text.",
			//       "The "+ "^[" +"["
			DecodeLen:         6,
			ExpectedEncodeLen: 7,
		},
		{
			Encoded: "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0m"+" text.",
			//       "The "+ "^[" +"[1"
			DecodeLen:         7,
			ExpectedEncodeLen: 8,
		},
		{
			Encoded: "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0m"+" text.",
			//       "The "+ "^[" +"[1m"
			DecodeLen:         8,
			ExpectedEncodeLen: 9,
		},
		{
			Encoded: "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0m"+" text.",
			//       "The "+ "^[" +"[1m" +"b"
			DecodeLen:         9,
			ExpectedEncodeLen: 10,
		},
		{
			Encoded: "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0m"+" text.",
			//       "The "+ "^[" +"[1m" +"bo"
			DecodeLen:         10,
			ExpectedEncodeLen: 11,
		},
		{
			Encoded: "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0m"+" text.",
			//       "The "+ "^[" +"[1m" +"bol"
			DecodeLen:         11,
			ExpectedEncodeLen: 12,
		},
		{
			Encoded: "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0m"+" text.",
			//       "The "+ "^[" +"[1m" +"bold"
			DecodeLen:         12,
			ExpectedEncodeLen: 13,
		},
		{
			Encoded: "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0m"+" text.",
			//       "The "+ "^[" +"[1m" +"bold"+ "^["
			DecodeLen:         13,
			ExpectedEncodeLen: 15,
		},
		{
			Encoded: "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0m"+" text.",
			//       "The "+ "^[" +"[1m" +"bold"+ "^[" +"["
			DecodeLen:         14,
			ExpectedEncodeLen: 16,
		},
		{
			Encoded: "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0m"+" text.",
			//       "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0"
			DecodeLen:         15,
			ExpectedEncodeLen: 17,
		},
		{
			Encoded: "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0m"+" text.",
			//       "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0m"
			DecodeLen:         16,
			ExpectedEncodeLen: 18,
		},
		{
			Encoded: "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0m"+" text.",
			//       "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0m"+" "
			DecodeLen:         17,
			ExpectedEncodeLen: 19,
		},
		{
			Encoded: "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0m"+" text.",
			//       "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0m"+" t"
			DecodeLen:         18,
			ExpectedEncodeLen: 20,
		},
		{
			Encoded: "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0m"+" text.",
			//       "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0m"+" te"
			DecodeLen:         19,
			ExpectedEncodeLen: 21,
		},
		{
			Encoded: "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0m"+" text.",
			//       "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0m"+" tex"
			DecodeLen:         20,
			ExpectedEncodeLen: 22,
		},
		{
			Encoded: "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0m"+" text.",
			//       "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0m"+" text"
			DecodeLen:         21,
			ExpectedEncodeLen: 23,
		},
		{
			Encoded: "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0m"+" text.",
			//       "The "+ "^[" +"[1m" +"bold"+ "^[" +"[0m"+" text."
			DecodeLen:         22,
			ExpectedEncodeLen: 24,
		},
	}

	for testNumber, test := range tests {

		actual := encodedlen(test.DecodeLen, []byte(test.Encoded))

		if expected := test.ExpectedEncodeLen; expected != actual {
			t.Errorf("For test #%d, the actual calculated encoded length was not what was expected", testNumber)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			t.Logf("DECODED LENGTH: %d", test.DecodeLen)
			t.Log("ENCODED:...")
			t.Logf("%q", test.Encoded)
			continue
		}
	}
}
