package comment_analyzer

import (
	"bufio"
	"compass.com/go-homework/model"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// countComments analyzes a file to count comment lines
func CountComments(filePath string) (model.CommentStats, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return model.CommentStats{}, err
	}
	defer file.Close()

	stats := model.CommentStats{}
	InStringLiteral = true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if IsInLineComment {
			stats.Inline++
			if line[len(line)-1] != '\\' {
				IsInLineComment = false
			}
		}

		if !IsInBlockComment {
			line, InStringLiteral = ParseLineForComments(line, &stats)
		} else {
			pos := strings.Index(line, "*/")
			if pos != -1 {
				line = line[pos:]
				line, InStringLiteral = ParseLineForComments(line, &stats)
			}
		}
		stats.Total++
		if !InStringLiteral {
			continue
		}
		if IsInBlockComment {
			stats.Block++
			if strings.Contains(line, "*/") {
				IsInBlockComment = false
			}
			if strings.Index(line, "*/") < strings.Index(line, "/*") {
				IsInBlockComment = true
			}
			continue
		}

		if strings.Contains(line, "/*") {
			IsInBlockComment = true
			stats.Block++
			if strings.Contains(line, "*/") {
				IsInBlockComment = false
			}
			continue
		}

		if strings.Contains(line, "//") {
			stats.Inline++
		}
	}

	if err := scanner.Err(); err != nil {
		return model.CommentStats{}, err
	}

	return stats, nil
}

// ProcessDirectory processes all files in the specified directory and returns a map of comment statistics
// for each source code file with extensions .c, .cpp, .h, or .hpp.
// It walks through the directory recursively and collects comment statistics using the CountComments function.
func ProcessDirectory(dir string) (map[string]model.CommentStats, error) {
	statsMap := make(map[string]model.CommentStats)

	err := filepath.WalkDir(dir, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !entry.IsDir() {
			ext := filepath.Ext(entry.Name())
			if ext == ".c" || ext == ".cpp" || ext == ".h" || ext == ".hpp" {
				stats, err := CountComments(path)
				if err != nil {
					return err
				}
				relPath, err := filepath.Rel(dir, path)
				if err != nil {
					return err
				}
				statsMap[relPath] = stats
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return statsMap, nil
}
