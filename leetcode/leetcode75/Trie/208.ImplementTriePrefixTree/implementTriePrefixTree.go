package main

import (
	"fmt"
)

// TrieNode represents a node in the Trie data structure
type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

// Trie represents the Trie data structure
type Trie struct {
	root *TrieNode
}

// Constructor initializes the Trie
func Constructor() Trie {
	return Trie{root: &TrieNode{}}
}

// Insert adds a word into the Trie
func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		index := ch - 'a'
		if node.children[index] == nil {
			node.children[index] = &TrieNode{}
		}
		node = node.children[index]
	}
	node.isEnd = true
}

// Search returns if the word is in the Trie
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		index := ch - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return node != nil && node.isEnd
}

// StartsWith returns if there is any word in the Trie that starts with the given prefix
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, ch := range prefix {
		index := ch - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return node != nil
}

func main() {
	trie := Constructor()

	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))   // Output: true
	fmt.Println(trie.Search("app"))     // Output: false
	fmt.Println(trie.StartsWith("app")) // Output: true

	trie.Insert("app")
	fmt.Println(trie.Search("app")) // Output: true
}
