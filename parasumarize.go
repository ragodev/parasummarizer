package parasummarize

import "strings"

type Summarizer struct {
	corpus    string
	wordCount map[string]int
}

// ReturnCorpus returns the current corpus being used
func (summarizer *Summarizer) ReturnCorpus() string {
	return summarizer.corpus
}

// ReturnCorpus returns the current corpus being used
func (summarizer *Summarizer) ReturnWordCount() map[string]int {
	return summarizer.wordCount
}

func NewSummarizer(corpus string) *Summarizer {
	return &Summarizer{corpus: corpus, wordCount: WordCount(cleanString(corpus))}
}

func (summarizer *Summarizer) Summarize(paragraph string) []string {
	validStrings := []string{}
	substrs := strings.Fields(cleanString(paragraph))
	for _, word := range substrs {
		if val, ok := summarizer.wordCount[word]; ok {
			if val == 1 {
				validStrings = append(validStrings, word)
			}
		}
	}
	return validStrings
}

func main() {

}

func WordCount(s string) map[string]int {
	substrs := strings.Fields(s)
	counts := make(map[string]int)
	for _, word := range substrs {
		_, ok := counts[word]
		if ok == true {
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	}
	return counts
}

// https://www.socketloop.com/tutorials/golang-removes-punctuation-or-defined-delimiter-from-the-user-s-input
const delim = "?!.;,*"

func isDelim(c string) bool {
	if strings.Contains(delim, c) {
		return true
	}
	return false
}
func cleanString(input string) string {
	size := len(input)
	temp := ""
	var prevChar string
	for i := 0; i < size; i++ {
		str := string(input[i])
		if (str == " " && prevChar != " ") || !isDelim(str) {
			temp += str
			prevChar = str
		} else if prevChar != " " && isDelim(str) {
			temp += " "
		}
	}
	return temp
}
