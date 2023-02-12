package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

const (
	reTemplateWord = `([,!'.":]+)|(-$)`
)

var reWord = regexp.MustCompile(reTemplateWord)

// Words ...
type Words struct {
	word      string
	frequency int
}

// Top10 ...
func Top10(text string) []string {
	// split text
	textSplit := strings.Fields(text)
	if len(textSplit) == 1 {
		return nil
	}
	// count words in text
	mapWords := make(map[string]int, len(textSplit))
	for _, v := range textSplit {
		templante := reWord.ReplaceAllString(v, "")
		if templante != "" {
			mapWords[strings.ToLower(templante)]++
		}
	}
	// sort map for frequency and lexicograph
	wSlices := make([]Words, 0, len(mapWords))
	for k, v := range mapWords {
		wSlices = append(wSlices, Words{k, v})
	}
	sort.Slice(wSlices, func(i, j int) bool {
		if wSlices[i].frequency < wSlices[j].frequency {
			return false
		}
		if wSlices[i].frequency > wSlices[j].frequency {
			return true
		}
		return strings.Compare(wSlices[i].word, wSlices[j].word) == -1
	})
	// get first 10 sorted words
	result := make([]string, 0, 10)
	for i := 0; i < len(wSlices) && i < 10; i++ {
		result = append(result, wSlices[i].word)
	}

	return result
}
