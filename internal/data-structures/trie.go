package datastructures

type Trie struct {
	children map[string]*Trie
	end      bool
}

func NewTrie() *Trie {
	return &Trie{
		children: make(map[string]*Trie),
		end:      false,
	}
}

func (t *Trie) Insert(word string) {
	current := t
	for _, char := range word {
		charStr := string(char)
		if _, ok := current.children[charStr]; !ok {
			current.children[charStr] = NewTrie()
		}
		current = current.children[charStr]
	}
	current.end = true
}

func (t *Trie) InsertMany(words []string) {
	for _, word := range words {
		t.Insert(word)
	}
}

func (t *Trie) Search(word string) bool {
	current := t
	for _, char := range word {
		charStr := string(char)
		if _, ok := current.children[charStr]; !ok {
			return false
		}
		current = current.children[charStr]
	}
	return current.end
}

func (t *Trie) StartsWith(prefix string) bool {
	current := t
	for _, char := range prefix {
		charStr := string(char)
		if _, ok := current.children[charStr]; !ok {
			return false
		}
		current = current.children[charStr]
	}
	return true
}
