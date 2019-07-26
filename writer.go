package caret

import (
	"io"
	"strings"
	"unicode/utf8"
)

// Writer lets you write text encoded in Caret Notation, and writes the decoded characters to the nested io.Writer.
//
// Example
//
// Here is an example using caret.Writer
//
//	var writer io.Writer
//	
//	// ...
//	
//	var caretWriter caret.Writer = caret.Writer{writer}
//	
//	var caretText = []byte("The "+ "\x1b" +"[34m" +"blue"+ "\x1b" +"[0m"+" text.")
//	
//	caretWriter.Write(caretText)
type Writer struct {
	Writer io.Writer
}

func (receiver *Writer) Write(p []byte) (int, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}

	writer := receiver.Writer
	if nil == writer {
		return 0, errNilWriter
	}

	if 0 >= len(p) {
		return 0, nil
	}

	var builder strings.Builder
	{
		var careted bool

		ptr := p
		for 0 < len(ptr) {
			r, size := utf8.DecodeRune(ptr)
			if utf8.RuneError == r && 0 == size {
				break
			}
			if utf8.RuneError == r {
				return 0, errRuneError
			}
			if 0 >= size {
				return 0, errInternalError
			}

			ptr = ptr[size:]

			switch {
			case careted:
				careted = false
				decoded := decode(r)
				builder.WriteRune(decoded)
			case '^' == r:
				careted = true
			default:
				builder.WriteRune(r)
			}

		}
	}

	n, err := io.WriteString(writer, builder.String())
	return encodedlen(n, p), err
}
