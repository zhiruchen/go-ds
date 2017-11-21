package advanced

const ALPHABET_SIZE = 26

type TrieNode struct {
	children  []*TrieNode
	endOfWord bool
}

type Trie struct {
	root *TrieNode
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
