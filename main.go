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
}

var DEFAULT_FILE = currentPath() + "./word.list"

func main() {
	// Assign file path
	file := DEFAULT_FILE
	if len(os.Args) > 1 {
		file = os.Args[1]
		if (file == "-h") {
			printUsage()
			return
		}

	}

	word := findLongestWord(file)
	fmt.Printf("The longest compound-word is: %s\n", word)
}

// Looks longest compound word on the filename
// passed as an argument.
//
// Basically whats happening is: We store all words on a Trie Tree,
// then, for each inserted word, get all words that have at least
// one previously scanned word as a prefix; We store this word as a
// potential candidate alongside the separated-remaining suffix.
// Then we loop until our Stack is empty by  popping out candidates. On
// each candidate, we check if the suffix is already a word on the tree
// and if the full word length is greater than our longest one (which is nil at the beggining).
// If not, then we check if the suffix, may be another compound word itself
// by breaking it down to prefixes and adding candidates to the stack as well.
//
func findLongestWord(filename string) string {
	longest := "" // empty base string

	// Read and defer close file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Make sure you have word.list alongside the executable.")
		fmt.Println(err)
		printUsage();
		os.Exit(1)
	}
	defer file.Close()

	// Start our scanner and initialize our
	// needed collections
	scn := bufio.NewScanner(file)
	trie := collections.NewTrie() // stores words and lookup prefixes
	stack := collections.NewStack() // saves compound candidates

	for scn.Scan() {
		word := scn.Text()

		prefixes := trie.Insert(word)
		// loop through the prefixes found while inserting the word,
		// and push to the stack potential candidates.
		for _, prefix := range prefixes {
			suffix := word[len(prefix):]
			stack.Push(CompoundPair{word, suffix})
		}
	}

	for !stack.IsEmpty() {
		compound := stack.Pop().(CompoundPair)
		word := string(compound.prefix)
		suffix := string(compound.suffix)
		// if the suffix is a word itself, then just compare lenghts
		if trie.HasWord(suffix) && len(word) > len(longest) {
			longest = word
		} else {
			// break own suffix to see if it is a compound word itself.
			prefixes := trie.PrefixesOfWord(suffix)
			for _, prefix := range prefixes {
				suffix := suffix[len(prefix):] // get new suffix
				stack.Push(CompoundPair{word, suffix})
			}
		}
	}

	return longest
}

// Prints how-to-use
func printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" main.exe [command/file]")
	fmt.Println(" command -h: displays this message")
	fmt.Println(" file: absolute/relative path of file to load")
	fmt.Println("Examples:")
	fmt.Println(" main.exe \"C:\\user\\word.list\"")
	fmt.Println(" main.exe word2.list")
	fmt.Println(" main.exe -h (Displays this message)")
}

// Finds path of this file binary.
func currentPath() string {
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return path
}
