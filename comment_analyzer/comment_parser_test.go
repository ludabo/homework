package comment_analyzer

import (
	"testing"
)

// TestFindFirstUnescapedQuote tests various scenarios for FindFirstUnescapedQuote function.
func TestFindFirstUnescapedQuote(t *testing.T) {
	tests := []struct {
		content  string
		expected int
	}{
		{"", -1},
		{"  String out = this->hasComment(commentBefore) ? \"\\n\" : \"\";", 49},
		{"  String out = this->hasComment(commentBefore) ? \\\"\\n\" : \"\";", 53},
	}

	for _, test := range tests {
		t.Run(test.content, func(t *testing.T) {
			index := FindFirstUnescapedQuote(test.content)
			if index != test.expected {
				t.Errorf("For content %q, expected index %d, but got %d", test.content, test.expected, index)
			}
		})
	}
}

// TestFindQuoteIndices tests various scenarios for FindQuoteIndices function.
func TestFindQuoteIndices(t *testing.T) {
	tests := []struct {
		content        string
		expectedFirst  int
		expectedSecond int
	}{
		{"", -1, -1},
		{"  String out = this->hasComment(commentBefore) ? \"\\n\" : \"\";", 49, 52},
		{"  String out = this->hasComment(commentBefore) ? \\\"\\n\" : \"\";", 53, 57},
	}

	for _, test := range tests {
		t.Run(test.content, func(t *testing.T) {
			firstQuoteIndex, secondQuoteIndex := FindQuoteIndices(test.content)
			if firstQuoteIndex != test.expectedFirst {
				t.Errorf("For content %q, expected first quote index %d, but got %d", test.content, test.expectedFirst, firstQuoteIndex)
			}
			if secondQuoteIndex != test.expectedSecond {
				t.Errorf("For content %q, expected second quote index %d, but got %d", test.content, test.expectedSecond, secondQuoteIndex)
			}
		})
	}
}

// TestFindSingleQuoteIndices tests various scenarios for FindSingleQuoteIndices function.
func TestFindSingleQuoteIndices(t *testing.T) {
	tests := []struct {
		content        string
		expectedFirst  int
		expectedSecond int
	}{
		{"", -1, -1},
		{"  String out = thi's->hasComment(commentBef'ore) ? \"\\n\" : \"\";", 18, 43},
		{"  String out = this->has'Comment(commentBef'ore) ? \\\"\\n\" : \"\";", 24, 43},
	}

	for _, test := range tests {
		t.Run(test.content, func(t *testing.T) {
			firstQuoteIndex, secondQuoteIndex := FindSingleQuoteIndices(test.content)
			if firstQuoteIndex != test.expectedFirst {
				t.Errorf("For content %q, expected first single quote index %d, but got %d", test.content, test.expectedFirst, firstQuoteIndex)
			}
			if secondQuoteIndex != test.expectedSecond {
				t.Errorf("For content %q, expected second single quote index %d, but got %d", test.content, test.expectedSecond, secondQuoteIndex)
			}
		})
	}
}

// TestRemoveQuoteAndParen tests various scenarios for RemoveQuoteAndParen function.
func TestRemoveQuoteAndParen(t *testing.T) {
	tests := []struct {
		content           string
		firstQuoteIndex   int
		closingParenIndex int
		offset            int
		expectedResult    string
	}{
		{"", 3, 9, 0, ""},
		{"  String out = thi's->hasComment(commentBef'ore) ? \"\\n\" : \"\";", 3, 9, 0, "  Sout = thi's->hasComment(commentBef'ore) ? \"\\n\" : \"\";"},
	}

	for _, test := range tests {
		t.Run(test.content, func(t *testing.T) {
			result := RemoveQuoteAndParen(test.content, test.firstQuoteIndex, test.closingParenIndex, test.offset)
			if result != test.expectedResult {
				t.Errorf("For content %q, expected result %q, but got %q", test.content, test.expectedResult, result)
			}
		})
	}
}
