package caret_test

import (
	"github.com/reiver/go-caret"

	"bytes"
	"io"

	"testing"
)

func TestDecoderWrite(t *testing.T) {

	tests := []struct{
		CaretText string
		Expected  string
	}{
		{
			CaretText: "",
			Expected:  "",
		},



		{
			CaretText: "apple",
			Expected:  "apple",
		},
		{
			CaretText: "BANANA",
			Expected:  "BANANA",
		},
		{
			CaretText: "Cherry",
			Expected:  "Cherry",
		},
		{
			CaretText: "dATE",
			Expected:  "dATE",
		},



		{
			CaretText: "Hello world!",
			Expected:  "Hello world!",
		},



		{
			CaretText: "first_name"+ "^_"   +"middle_name"+ "^_"   +"family_name"+ "^^"   +"James"+ "^_"   +"Tiberius"+ "^_"   +"Kirk",
			Expected:  "first_name"+ "\x1F" +"middle_name"+ "\x1F" +"family_name"+ "\x1E" +"James"+ "\x1F" +"Tiberius"+ "\x1F" +"Kirk",
		},



		{
			CaretText: "The "+ "^["   +"[1m" +"bold"+ "^["   +"[0m"+" text.",
			Expected:  "The "+ "\x1b" +"[1m" +"bold"+ "\x1b" +"[0m"+" text.",
		},



		{
			CaretText: "Hi! 🙂 "+ "^["   +"[31m" +"red ۵"+ "^["   +"[0m"+" 👻 - 👽 / 👾",
			Expected:  "Hi! 🙂 "+ "\x1b" +"[31m" +"red ۵"+ "\x1b" +"[0m"+" 👻 - 👽 / 👾",
		},
	}

	for testNumber, test := range tests {

		var buffer bytes.Buffer

		var caretDecoder caret.Decoder = caret.Decoder{&buffer}

		n, err := io.WriteString(&caretDecoder, test.CaretText)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			continue
		}
		if expected, actual := len(test.CaretText), n; expected != actual {
			t.Errorf("For test #%d, the number of bytes actually written, was not what was expected.", testNumber)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			t.Log("EXPECTED text:...",)
			t.Log(test.Expected)
			t.Log("ACTUAL text:...",)
			t.Log(buffer.String())
			continue
		}

		if expected, actual := test.Expected, buffer.String(); expected != actual {
			t.Errorf("For test #%d, what was written was not what was expected", testNumber)
			t.Log("EXPECTED:...")
			t.Log(expected)
			t.Log("ACTUAL:...")
			t.Log(actual)
			continue
		}
	}
}
