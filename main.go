// Given a list of words like https://github.com/NodePrime/quiz/blob/master/word.list
// find the longest compound-word in the list, which is also a concatenation of other
// sub-words that exist in the list.
//
// The program should allow the user to input different data.
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"github.com/rapito/quiz/collections"
)

type CompoundPair struct {
	prefix string
	suffix string
//	prefixes []string
}

var DEFAULT_FILE = currentPath() + "./word.list"

func main() {
	// Assign file path
	file := DEFAULT_FILE
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	word := findLongestWord2(file)
	fmt.Printf("The longest compound-word is: %s\n", word)
}

func findLongestWord2(filename string) string {
	longest := "" // empty base string

	// Read and defer close file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scn := bufio.NewScanner(file)
	trie := collections.NewTrie()
	stack := collections.NewStack()

	for scn.Scan() {
		word := scn.Text()
		prefixes := trie.Insert(word)
		for _, prefix := range prefixes {
			suffix := word[len(prefix):]
			stack.Push(CompoundPair{word, suffix})
		}
	}

	for !stack.IsEmpty() {
		compound := stack.Pop().(CompoundPair)
		word := string(compound.prefix)
		suffix := string(compound.suffix)
//		fmt.Printf("word: %s, suffix: %s\n", word, suffix)
		if trie.HasWord(suffix) && len(word) > len(longest) {
//			fmt.Println("HasWord :D " + word)
			longest = word

		}
//		else {
//			prefixes := trie.PrefixesOfWord(word)
//			for i, prefix := range prefixes {
//				fmt.Printf("prefixes index:%d prefix:%s size:%d\n", i, prefix, len(prefixes))
//				suffix := word[len(prefix):]
//				fmt.Printf("suffix %s\n", suffix)
//				stack.Push(CompoundPair{word, suffix})
//			}
//		}
	}

	return longest
}

// Looks longest compound word on the filename
// passed as an argument.
func findLongestWord(filename string) string {

	word := "" // empty base string

	// Read and defer close file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// Bugger Read file with bufio's Scanner
	// and simply compare lengths
	scn := bufio.NewScanner(file)
	for scn.Scan() {
		t := scn.Text()
		if len(t) > len(word) {
			word = t
		}
	}

	return word
}

// Finds path of this file binary.
func currentPath() string {
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return path
}
