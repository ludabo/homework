package comment_analyzer

import (
	"compass.com/go-homework/model"
	"fmt"
	"sort"
)

// printStats prints the collected statistics in the required format
func PrintStats(statsMap map[string]model.CommentStats) {
	var filenames []string
	for filename := range statsMap {
		filenames = append(filenames, filename)
	}
	sort.Strings(filenames)

	for _, filename := range filenames {
		stats := statsMap[filename]
		fmt.Printf("%s      total: %4d    inline: %4d    block: %4d\n",
			filename, stats.Total, stats.Inline, stats.Block)
	}
}
