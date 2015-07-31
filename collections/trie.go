package collections

type Trie struct  {
	root *Node
}

type Node struct {
	char   byte
	isWord bool
	children map[byte]*Node
}

func NewTrie() *Trie {
	return &Trie{
		root: NewNode(byte(1)),
	}
}

// Inserts new word to trie and returns
// prefixes found along the way
func (t *Trie) Insert(word string) ([]string) {
	//	fmt.Println("Insert: " + word);
	previousWords := make([]string, 0)

	currentNode := t.root
	for i, char := range word {
		//		fmt.Printf("i: %d, char: %s \n", i, string(char));

		// append to prefix
		if currentNode.isWord {
			previousWords = append(previousWords, string(word[:i]))
		}

		child := currentNode.children[byte(char)]
		if child == nil {
			child = NewNode(byte(char))
			currentNode.children[byte(char)] = child
		}
		currentNode = child
	}
	currentNode.isWord = true

	return previousWords
}

// Return prefixes found beofre reaching the end of the word
func (t *Trie) PrefixesOfWord(word string) []string {
	previousWords := make([]string, 0)

	currentNode := t.root

	for i := 0; i < len(word); i++ {
		char := word[i]
		// append to prefix
		if currentNode.isWord {
			previousWords = append(previousWords, word[:i])
		}

		child := currentNode.children[char]
		if child == nil {
			break
		}
		currentNode = child
	}

	return previousWords
}

// Return whether the Trie has or not the word
func (t *Trie) HasWord(word string) bool {

//	fmt.Println("HasWord: " + word);
	currentNode := t.root

	for i,char := range word {
//		char := word[i]
		// append to prefix
		if currentNode.isWord && i == len(word)-1 {
			return true;
		}

		child := currentNode.children[byte(char)]
		if child == nil {
			return false
		}
		currentNode = child
	}

	return true
}

func NewNode(char byte) *Node {
	return &Node{
		char: char,
		isWord: false,
		children: make(map[byte]*Node),
	}
}
