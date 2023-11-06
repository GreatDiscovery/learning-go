package trie

import (
	"fmt"
	patricia "github.com/tchap/go-patricia/v2/patricia"
	"testing"
)

// 单词匹配树：https://bbs.huaweicloud.com/blogs/366935
// 优点： 查找快、可以做前缀匹配
// 缺点：空间大，典型的用空间换时间

func TestWordSearchTree(t *testing.T) {
	printItem := func(prefix patricia.Prefix, item patricia.Item) error {
		fmt.Printf("%q: %v\n", prefix, item)
		return nil
	}

	// Create a new default trie (using the default parameter values).
	//trie := patricia.NewTrie()

	// Create a new custom trie.
	trie := patricia.NewTrie(patricia.MaxPrefixPerNode(16), patricia.MaxChildrenPerSparseNode(10))

	// Insert some items.
	trie.Insert(patricia.Prefix("Pepa Novak"), 1)
	trie.Insert(patricia.Prefix("Pepa Sindelar"), 2)
	trie.Insert(patricia.Prefix("Karel Macha"), 3)
	trie.Insert(patricia.Prefix("Karel Hynek Macha"), 4)

	// Just check if some things are present in the tree.
	key := patricia.Prefix("Pepa Novak")
	fmt.Printf("%q present? %v\n", key, trie.Match(key))
	// "Pepa Novak" present? true
	key = patricia.Prefix("Karel")
	fmt.Printf("Anybody called %q here? %v\n", key, trie.MatchSubtree(key))
	// Anybody called "Karel" here? true

	// Walk the tree in alphabetical order.
	trie.Visit(printItem)
	// "Karel Hynek Macha": 4
	// "Karel Macha": 3
	// "Pepa Novak": 1
	// "Pepa Sindelar": 2

	// Walk a subtree.
	trie.VisitSubtree(patricia.Prefix("Pepa"), printItem)
	// "Pepa Novak": 1
	// "Pepa Sindelar": 2

	// Modify an item, then fetch it from the tree.
	trie.Set(patricia.Prefix("Karel Hynek Macha"), 10)
	key = patricia.Prefix("Karel Hynek Macha")
	fmt.Printf("%q: %v\n", key, trie.Get(key))
	// "Karel Hynek Macha": 10

	// Walk prefixes.
	prefix := patricia.Prefix("Karel Hynek Macha je kouzelnik")
	trie.VisitPrefixes(prefix, printItem)
	// "Karel Hynek Macha": 10

	// Delete some items.
	trie.Delete(patricia.Prefix("Pepa Novak"))
	trie.Delete(patricia.Prefix("Karel Macha"))

	// Walk again.
	trie.Visit(printItem)
	// "Karel Hynek Macha": 10
	// "Pepa Sindelar": 2

	// Delete a subtree.
	trie.DeleteSubtree(patricia.Prefix("Pepa"))

	// Print what is left.
	trie.Visit(printItem)
	// "Karel Hynek Macha": 10
}
