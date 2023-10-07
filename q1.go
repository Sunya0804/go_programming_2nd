package ds_hw_0

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// Find the top K most common words in a text document.
//
//	path: location of the document
//	numWords: number of words to return (i.e. k)
//	charThreshold: character threshold for whether a token qualifies as a word,
//		e.g. charThreshold = 5 means "apple" is a word but "pear" is not.
//
// Matching is case-insensitive, e.g. "Orange" and "orange" is considered the same word.
// A word comprises alphanumeric characters only. All punctuations and other characters
// are removed, e.g. "don't" becomes "dont".
// You should use `checkError` to handle potential errors.
func TopWords(path string, numWords int, charThreshold int) []WordCount {
	// TODO: implement me
	// HINT: You may find the `strings.Fields` and `strings.ToLower` functions helpful
	// HINT: To keep only alphanumeric characters, use the regex "[^0-9a-zA-Z]+"

	// Open the file
	file, err := os.Open(path)
	checkError(err)
	defer file.Close()

	// Create a map to store the word counts
	wordCounts := make(map[string]int)

	// Create a slice to store the word counts
	var wordCountSlice []WordCount

	for {
		// Read the file line by line
		line, err := readLine(file)
		if err != nil {
			break
		}

		// Split the line into words
		words := strings.Fields(strings.ToLower(line))

		// Remove all punctuations and other characters
		for i, word := range words {
			words[i] = removePunctuations(word)
		}

		// Count the words
		for _, word := range words {
			if len(word) >= charThreshold {
				wordCounts[word]++
			}
		}
	}

	// Convert the map to a slice
	for word, count := range wordCounts {
		wordCountSlice = append(wordCountSlice, WordCount{word, count})
	}

	// Sort the slice
	sortWordCounts(wordCountSlice)

	// Return the top K words
	return wordCountSlice[:numWords]
}

// Helper function to read a line from a file
func readLine(file *os.File) (string, error) {
	var line string
	var err error
	var buf [1]byte

	for {
		_, err = file.Read(buf[:])
		if err != nil {
			return "", err
		}
		if buf[0] == '\n' {
			break
		}
		line += string(buf[0])
	}

	return line, nil
}

// Helper function to remove all punctuations and other characters
func removePunctuations(str string) string {
	var newLine string
	for _, char := range str {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			newLine += string(char)
		}
	}
	return newLine
}

// WordCount A struct that represents how many times a word is observed in a document
type WordCount struct {
	Word  string
	Count int
}

func (wc WordCount) String() string {
	return fmt.Sprintf("%v: %v", wc.Word, wc.Count)
}

// Helper function to sort a list of word counts in place.
// This sorts by the count in decreasing order, breaking ties using the word.
// DO NOT MODIFY THIS FUNCTION!
func sortWordCounts(wordCounts []WordCount) {
	sort.Slice(wordCounts, func(i, j int) bool {
		wc1 := wordCounts[i]
		wc2 := wordCounts[j]
		if wc1.Count == wc2.Count {
			return wc1.Word < wc2.Word
		}
		return wc1.Count > wc2.Count
	})
}
