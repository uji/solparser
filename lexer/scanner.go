package lexer

import (
	"unicode/utf8"

	"github.com/uji/solparser/token"
)

// isSpace reports whether the character is a Unicode white space character.
// We avoid dependency on the unicode package, but check validity of the implementation
// in the tests.
func isSpace(r rune) bool {
	if r <= '\u00FF' {
		// Obvious ASCII ones: \t through \r plus space. Plus two Latin-1 oddballs.
		switch r {
		case ' ', '\t', '\v', '\f', '\r':
			return true
		case '\u0085', '\u00A0':
			return true
		}
		return false
	}
	// High-valued ones.
	if '\u2000' <= r && r <= '\u200a' {
		return true
	}
	switch r {
	case '\u1680', '\u2028', '\u2029', '\u202f', '\u205f', '\u3000':
		return true
	}
	return false
}

func isMiscToken(r rune) bool {
	return token.AsMiscToken(r) != token.Unknown
}

func ScanTokens(data []byte, atEOF bool) (advance int, token []byte, err error) {
	start := 0

	// Return newline code or misk token.
	r, width := utf8.DecodeRune(data[start:])
	if r == '\n' || isMiscToken(r) {
		return start + width, data[start : start+width], nil
	}

	tokenIsSpace := isSpace(r)
	// Scan until isSpace result changed, marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if r == '\n' || isMiscToken(r) || isSpace(r) != tokenIsSpace {
			return i, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}
