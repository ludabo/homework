package comment_analyzer

import (
	"compass.com/go-homework/model"
	"testing"
)

func TestCountComments(t *testing.T) {
	tests := []struct {
		path     string
		expected model.CommentStats
	}{
		{
			path:     `..\testing\cpp\lib_json\json_reader.cpp`,
			expected: model.CommentStats{Total: 1992, Inline: 134, Block: 0},
		}, {
			path:     `..\testing\cpp\lib_json\json_tool.h`,
			expected: model.CommentStats{Total: 138, Inline: 13, Block: 19},
		}, {
			path:     `..\testing\cpp\lib_json\json_value.cpp`,
			expected: model.CommentStats{Total: 1634, Inline: 111, Block: 18},
		}, {
			path:     `..\testing\cpp\lib_json\json_writer.cpp`,
			expected: model.CommentStats{Total: 1259, Inline: 89, Block: 0},
		}, {
			path:     `..\testing\cpp\special_cases.cpp`,
			expected: model.CommentStats{Total: 62, Inline: 6, Block: 34},
		}, {
			path:     `..\testing\cpp\test_lib_json\fuzz.cpp`,
			expected: model.CommentStats{Total: 54, Inline: 5, Block: 0},
		}, {
			path:     `..\testing\cpp\test_lib_json\fuzz.h`,
			expected: model.CommentStats{Total: 14, Inline: 5, Block: 0},
		}, {
			path:     `..\testing\cpp\test_lib_json\jsontest.cpp`,
			expected: model.CommentStats{Total: 430, Inline: 54, Block: 1},
		}, {
			path:     `..\testing\cpp\test_lib_json\jsontest.h`,
			expected: model.CommentStats{Total: 288, Inline: 52, Block: 8},
		}, {
			path:     `..\testing\cpp\test_lib_json\main.cpp`,
			expected: model.CommentStats{Total: 3971, Inline: 182, Block: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			stat, err := CountComments(tt.path)
			if err != nil {
				t.Errorf("error counting comments: %v", err)
			}
			if stat != tt.expected {
				t.Errorf("expected %+v, got %+v", tt.expected, stat)
			}
		})
	}
}
