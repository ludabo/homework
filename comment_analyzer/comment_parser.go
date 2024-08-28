package comment_analyzer

import (
	"compass.com/go-homework/model"
	"strings"
)

// findFirstUnescapedQuote finds the position of the first unescaped quote in the string
func FindFirstUnescapedQuote(s string) int {
	length := len(s)
	i := 0

	for i < length {
		pos := strings.Index(s[i:], `"`)
		if pos == -1 {
			return -1
		}
		actualPos := i + pos

		if actualPos == 0 || s[actualPos-1] != '\\' {
			return actualPos
		} else {
			count := 0
			for i := actualPos; i > 0; i-- {
				if s[i-1] == '\\' {
					count++
				} else {
					break
				}
			}
			if count%2 == 0 {
				return actualPos
			}
		}

		i = actualPos + 1
	}

	return -1
}

// parseLineForComments processes a line to count comments and handle string literals
func ParseLineForComments(line string, c *model.CommentStats) (string, bool) {
	commentIndex := strings.Index(line, `//`)
	firstQuoteIndex, secondQuoteIndex := FindQuoteIndices(line)

	singleQuoteIndex, secondSingleQuoteIndex := FindSingleQuoteIndices(line)

	if secondQuoteIndex != -1 {
		if (firstQuoteIndex < commentIndex || commentIndex == -1) && (firstQuoteIndex < singleQuoteIndex || secondSingleQuoteIndex == -1) {
			if line[firstQuoteIndex+1] == '(' {
				closingParenIndex := strings.Index(line, `)"`)
				if closingParenIndex != -1 {

					line = RemoveQuoteAndParen(line, firstQuoteIndex, closingParenIndex, 2)
					return ParseLineForComments(line, c)
				} else {
					line = RemoveQuoteAndParen(line, firstQuoteIndex, closingParenIndex, 1)
					return ParseLineForComments(line, c)
				}
			} else {
				line = RemoveQuoteAndParen(line, firstQuoteIndex, secondQuoteIndex, 1)
				return ParseLineForComments(line, c)
			}
		}
	}

	if secondSingleQuoteIndex != -1 {
		if (singleQuoteIndex < firstQuoteIndex || secondQuoteIndex == -1) && (singleQuoteIndex < commentIndex || commentIndex == -1) {
			line = RemoveQuoteAndParen(line, singleQuoteIndex, secondSingleQuoteIndex, 1)
			return ParseLineForComments(line, c)
		}
	}
	if (commentIndex < firstQuoteIndex || secondQuoteIndex == -1) && (commentIndex < singleQuoteIndex || secondSingleQuoteIndex == -1) && commentIndex != -1 {
		if firstQuoteIndex != -1 && (firstQuoteIndex < commentIndex && (secondQuoteIndex == -1 || secondQuoteIndex > commentIndex)) {
			InStringLiteral = false
		}
		if InStringLiteral {
			c.Inline++
			if line[len(line)-1] == '\\' {
				IsInLineComment = true
			}
		}

		line = strings.TrimSpace(line[:commentIndex+1])

		return line, InStringLiteral
	}
	if secondQuoteIndex == -1 && secondSingleQuoteIndex == -1 {
		return line, (firstQuoteIndex == -1 && InStringLiteral)
	}
	return "", (firstQuoteIndex == -1 && InStringLiteral)
}

// FindQuoteIndices finds the indices of the first and second unescaped double quotes in a given string.
// It returns the indices of these quotes, or -1 if the quote is not found.
func FindQuoteIndices(line string) (firstQuoteIndex, secondQuoteIndex int) {
	// Find the index of the first unescaped double quote
	firstQuoteIndex = FindFirstUnescapedQuote(line)
	secondQuoteIndex = -1

	if firstQuoteIndex != -1 {
		// Search for the second unescaped double quote after the first one
		val := line[firstQuoteIndex+1:]
		fval := FindFirstUnescapedQuote(val)
		if fval != -1 {
			// Adjust the index to be relative to the original string
			secondQuoteIndex = fval + firstQuoteIndex + 1
		}
	}

	// Assuming InStringLiteral is a global variable or passed as a parameter
	if !InStringLiteral && firstQuoteIndex != -1 {
		// If we're not currently inside a string literal, set the indices for the string literal
		secondQuoteIndex = firstQuoteIndex
		firstQuoteIndex = 0
		InStringLiteral = true
	}

	return firstQuoteIndex, secondQuoteIndex
}

// FindSingleQuoteIndices finds the indices of the first and second single quotes in a given string.
// It returns the indices of these quotes, or -1 if the quote is not found.
func FindSingleQuoteIndices(line string) (singleQuoteIndex, secondSingleQuoteIndex int) {
	// Find the index of the first single quote
	singleQuoteIndex = strings.Index(line, `'`)
	secondSingleQuoteIndex = -1

	if singleQuoteIndex != -1 {
		// Search for the second single quote after the first one
		val2 := line[singleQuoteIndex+1:]
		secondSingleQuoteIndex = strings.Index(val2, `'`)
		if secondSingleQuoteIndex != -1 {
			// Adjust the index to be relative to the original string
			secondSingleQuoteIndex = secondSingleQuoteIndex + singleQuoteIndex + 1
		}
	}

	return singleQuoteIndex, secondSingleQuoteIndex
}

// RemoveQuoteAndParen removes a portion of the string from a quote index to a parenthesis index plus an offset.
// This is useful for modifying strings by removing specific parts.
func RemoveQuoteAndParen(line string, firstQuoteIndex, closingParenIndex, offset int) string {
	if len(line) < closingParenIndex+offset {
		return line
	}
	// Extract the part of the string before the quote
	head := line[:firstQuoteIndex]
	// Extract the part of the string after the closing parenthesis and the given offset
	tail := line[closingParenIndex+offset:]
	// Combine the head and tail parts to form the new string
	return head + tail
}
