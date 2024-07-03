package trie

import mapset "github.com/deckarep/golang-set/v2"

const ALPHABET_SIZE = 26

type Node struct {
	children []*Node
	isEnd    bool
	idMap    mapset.Set[string]
}

func NewNode() *Node {
	n := new(Node)
	n.children = make([]*Node, ALPHABET_SIZE)
	n.isEnd = false
	n.idMap = mapset.NewSet[string]()
	return n
}

func (n *Node) IdMap() mapset.Set[string] {
	return n.idMap
}
