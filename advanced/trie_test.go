package advanced

import "testing"

func TestTrie_Search(t *testing.T) {
	testCases := []struct{
		keys []string
		target string
		result bool
	}{
		{
			[]string{"computer", "science", "programming", "language", "programmer"},
			"programmer",
			true,
		},
		{
			[]string{"morning", "taylor", "swift"},
			"wild",
			false,
		},
	}

	for _, c := range testCases {
		trie := NewTrie()
		for _, k := range c.keys {
			trie.Insert(k)
		}

		res := trie.Search(c.target)
		if res != c.result {
			t.Errorf("expect: %v, get: %v\n", c.result, res)
		}
	}
}

func TestTrie_Delete(t *testing.T) {
	testCases := []struct{
		keys []string
		target string
		result bool
	}{
		{
			[]string{"computer", "science", "programming", "language", "programmer"},
			"programmer",
			false,
		},
		{
			[]string{"morning", "taylor", "swift", "song"},
			"song",
			false,
		},
	}

	for _, c := range testCases {
		trie := NewTrie()
		for _, k := range c.keys {
			trie.Insert(k)
		}

		trie.Delete(c.target)

		res := trie.Search(c.target)
		if res != c.result {
			t.Errorf("expect: %v, get: %v\n", c.result, res)
		}
	}
}
