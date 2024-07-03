package trie

import (
	"regexp"
	"strings"
	"unicode"

	mapset "github.com/deckarep/golang-set/v2"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	t := new(Trie)
	t.root = NewNode()
	return t
}

func (t *Trie) Insert(key string, id string) {
	lvl := 0
	length := len(key)
	index := 0
	crawl := t.root
	alreadyExists := t.SearchWord(key)
	if alreadyExists != nil && alreadyExists.isEnd {
		alreadyExists.idMap.Add(id)
		return
	}
	for lvl < length {
		index = int(rune(key[lvl]) - 'a')
		if crawl.children[index] == nil {
			crawl.children[index] = NewNode()
		}
		crawl = crawl.children[index]
		lvl++
	}
	crawl.idMap.Add(id)
	crawl.isEnd = true

}

func (t *Trie) Search(key string) mapset.Set[string] {
	results := mapset.NewSet[string]()
	tokens := t.Tokenize(key, " ")
	for _, token := range tokens {
		node := t.SearchWithPrefix(token)
		node1 := t.SearchWord(token)
		if node != nil {
			results = results.Union(node.idMap)
		} else {
			results.Clear()
			break
		}
		if node1 != nil {
			results = results.Union(node1.idMap)
		}
	}

	return results
}

func (t *Trie) SearchWord(word string) *Node {
	word = t.Normalize(word)
	lvl := 0
	length := len(word)
	index := 0
	crawl := t.root
	for lvl < length {
		index = int(rune(word[lvl]) - 'a')
		if crawl.children[index] == nil {
			return nil
		}
		crawl = crawl.children[index]
		lvl++
	}
	return crawl
}

func (t *Trie) SearchWithPrefix(prefix string) *Node {
	prefix = t.Normalize(prefix)
	lvl := 0
	length := len(prefix)
	index := 0
	crawl := t.root
	for lvl < length {
		index = int(rune(prefix[lvl]) - 'a')
		if crawl.children[index] == nil {
			return nil
		}
		crawl = crawl.children[index]
		lvl++
	}
	return t.depthFirstSearch(crawl, prefix)
}

func (t *Trie) depthFirstSearch(node *Node, currentWord string) *Node {
	if node.isEnd {
		return node
	}
	for i := range ALPHABET_SIZE {
		if node.children[i] != nil {
			currentWord += string(rune('a' + i))
			return t.depthFirstSearch(node.children[i], currentWord)
		}
	}
	return nil
}

func (t *Trie) InsertText(id, text, separator string) {
	words := t.Tokenize(text, separator)
	for _, word := range words {
		t.Insert(word, id)
	}
}

func (t *Trie) Normalize(word string) string {
	word = strings.Trim(word, " ")
	word = strings.ToLower(word)
	ts := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	word, _, _ = transform.String(ts, word)
	n := norm.NFC.String(word)
	punctuation := regexp.MustCompile(`[[:punct:]]`)
	n = punctuation.ReplaceAllString(n, "")
	numbers := regexp.MustCompile(`[[:digit:]]`)
	n = numbers.ReplaceAllString(n, "")
	return n
}

func (t *Trie) Tokenize(text, separator string) []string {
	words := strings.Split(text, separator)
	for i, word := range words {
		words[i] = t.Normalize(word)
	}
	return words
}
