package tiresearch

func NewTire() *Tire {
	return &Tire{
		newTrieNode("/"),
		10,
	}
}
