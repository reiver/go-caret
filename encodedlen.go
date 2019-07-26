package caret

import (
	"unicode/utf8"
)

// encodedlen returns the length of bytes in ‘p’ that were written.
//
// So, for example, if the Caret Notation -- the encoded text -- is:
//
//	encoded := "The "+ "^["   +"[1m" +"bold"+ "^["   +"[0m"+" text."
//
// Then the decoded text is:
//
//	decoded := "The "+ "\x1b" +"[1m" +"bold"+ "\x1b" +"[0m"+" text."
//
// And if we were to write 10 bytes of the decoded text to the nested io.Writer in caret.Writer;
//  i.e., the following decoded text was written to the nested io.Writer in caret.Writer:
//
//	           "The "+ "\x1b" +"[1m" +"bo"
//
// Then this is 10 bytes.
//
// But the byte count caret.Writer.Write() returns needs to be the number of the input []byte
// (not how many were actually writte).
//
// So we need to return the number of bytes that represents in the Caret Notation -- in the encoded text:
//
//	           "The "+ "^[" +"[1m" +"bo"
//
// And that is 11 bytes.
//
// This is what this function does.
func encodedlen(decodedlen int, p []byte) int {
	if 0 == decodedlen {
		return 0
	}

	var encodedLength int
	{
		x := decodedlen

		ptr := p

		for 0 < x {
			r, size := utf8.DecodeRune(ptr)
			if utf8.RuneError == r && 0 == size {
				//ERROR
				return encodedLength
			}
			if utf8.RuneError == r {
				//ERROR
				return encodedLength
			}
			if 0 >= size {
				//ERROR
				return encodedLength
			}

			ptr = ptr[size:]

			if '^' != r {
				x--
			}
			encodedLength += size
		}
	}

	return encodedLength
}
