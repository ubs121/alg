package strings

// https://en.wikipedia.org/wiki/Trie
// a  data structure that supports the insert, search, and startsWith operations efficiently
// can be used for matching operations like prefix search, exact match, longest common prefix, and more
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

func NewTrie() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
	}
}

func (root *TrieNode) Insert(word string) {
	current := root
	for _, char := range word {
		if _, ok := current.children[char]; !ok {
			current.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
				isEnd:    false,
			}
		}
		current = current.children[char]
	}
	current.isEnd = true
}

func (root *TrieNode) Search(word string) bool {
	current := root
	for _, char := range word {
		if _, ok := current.children[char]; !ok {
			return false
		}
		current = current.children[char]
	}
	return current.isEnd
}

func (root *TrieNode) StartsWith(prefix string) bool {
	current := root
	for _, char := range prefix {
		if _, ok := current.children[char]; !ok {
			return false
		}
		current = current.children[char]
	}
	return true
}
