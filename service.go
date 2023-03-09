package triesearch

func NewTrie() *Trie {
	return &Trie{
		newTrieNode("/"),
		newTrieNode("/"),
		10,
	}
}
