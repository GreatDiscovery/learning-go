package alogrithm

import (
	"testing"
)

func TestTrieIsPrefixMatch(t *testing.T) {
	trie := NewTrie()

	prefixes := []string{"apple", "app", "ball", "c", "ccc"}
	for _, p := range prefixes {
		trie.Insert(p)
	}

	checkWords := map[string]bool{
		"a":     false,
		"ap":    false,
		"app":   true,
		"appl":  true,
		"appla": true,
		"apx":   false,
		"ban":   false,
		"bal":   false,
		"ball":  true,
		"c":     true,
		"cc":    true,
	}
	for word, expected := range checkWords {
		if trie.IsPrefixMatch(word) != expected {
			t.Fatalf("%s : expected(%v), actual(%v)", word, expected, trie.IsPrefixMatch(word))
		}
	}
}

func TestTrieSearch(t *testing.T) {
	trie := NewTrie()

	keyWords := []string{"apple", "app", "ball", "c", "ccc"}
	for _, p := range keyWords {
		trie.Insert(p)
	}

	checkWords := map[string]bool{
		"a":     false,
		"ap":    false,
		"app":   true,
		"appl":  false,
		"appla": false,
		"apx":   false,
		"ban":   false,
		"bal":   false,
		"ball":  true,
		"c":     true,
		"cc":    false,
	}
	for word, expected := range checkWords {
		if trie.Search(word) != expected {
			t.Fatalf("%s : expected(%v), actual(%v)", word, expected, trie.Search(word))
		}
	}
}

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// Trie, not concurrency safe
type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
			isEnd:    false,
		},
	}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if node.children[ch] == nil {
			node.children[ch] = &TrieNode{
				children: make(map[rune]*TrieNode),
				isEnd:    false,
			}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) IsPrefixMatch(word string) bool {
	node := t.root
	for _, ch := range word {
		node = node.children[ch]
		if node == nil {
			return false
		}
		if node.isEnd {
			return true
		}
	}
	return false
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		node = node.children[ch]
		if node == nil {
			return false
		}
	}
	return node.isEnd
}
