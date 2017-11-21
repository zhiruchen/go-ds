package advanced

const ALPHABET_SIZE = 26

type TrieNode struct {
	children  []*TrieNode
	endOfWord bool
}

type Trie struct {
	root *TrieNode
}

func (t *TrieNode) isLeafNode() bool {
	return t.endOfWord
}

func (t *TrieNode) isFreeNode() bool {
	for _, n := range t.children {
		if n != nil {
			return false
		}
	}

	return true
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make([]*TrieNode, ALPHABET_SIZE),
		},
	}
}

func (t *Trie) Insert(key string) {
	keyLen := len(key)
	p := t.root

	for i := 0; i < keyLen; i++ {
		index := key[i] - 'a'

		if p.children[index] == nil {
			p.children[index] = &TrieNode{
				children: make([]*TrieNode, ALPHABET_SIZE),
			}
		}

		p = p.children[index]
	}

	p.endOfWord = true
}

func (t *Trie) Search(key string) bool {
	p := t.root

	for i := 0; i < len(key); i++ {
		index := key[i] - 'a'

		if p.children[index] == nil {
			return false
		}

		p = p.children[index]
	}

	return p != nil && p.endOfWord
}

func (t *Trie) Delete(key string) bool {
	return t.delete(key, t.root, 0, len(key))
}

func (t *Trie) delete(key string, currentNode *TrieNode, level, length int) bool {
	if currentNode != nil {
		if level == length {
			if currentNode.endOfWord {
				currentNode.endOfWord = false
			}

			return currentNode.isFreeNode()

		} else {
			index := key[level] - 'a'
			if t.delete(key, currentNode.children[index], level+1, length) {
				currentNode.children[index] = nil

				return !currentNode.isLeafNode() && currentNode.isFreeNode()
			}
		}
	}

	return false
}
