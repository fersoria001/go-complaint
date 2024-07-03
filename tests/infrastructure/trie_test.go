package infrastructure_test

import (
	"go-complaint/infrastructure/trie"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	trie := trie.NewTrie()
	trie.InsertText("id1", "this is a piece of text 34 .,?", " ")
	// result := trie.Search("this is e") //this is split term and return sum
	result := trie.Search("th")
	assert.Equal(t, 1, result.Cardinality())
	result = trie.Search("this")
	assert.Equal(t, 1, result.Cardinality())
	result = trie.Search("this is a")
	assert.Equal(t, 1, result.Cardinality())
	result = trie.Search("this is a piece")
	assert.Equal(t, 1, result.Cardinality())
	result = trie.Search("this is a nice")
	assert.Equal(t, 0, result.Cardinality())

}
