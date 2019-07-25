package caret

func decode(r rune) rune {
	return r ^ 0x40
}
