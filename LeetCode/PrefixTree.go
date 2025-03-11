package main

type Trie struct {
	// Каждый индекс отражает букву английского алфавита
	children [26]*Trie
	isEnd    bool
}

func ConstructorTrie() Trie {
	newTrie := Trie{}
	return newTrie
}

func (this *Trie) Insert(word string) {
	curr := this
	for _, ch := range word {
		// Рассчитываем индекс
		idx := ch - 'a'
		// Если такого узла нет
		// Создаем новый узел и переходим в него
		if curr.children[idx] == nil {
			n := Trie{}
			curr.children[idx] = &n
		}
		curr = curr.children[idx]
	}
	// Обозначаем конец слова
	curr.isEnd = true
}

func (this *Trie) Search(word string) bool {
	curr := this
	for _, ch := range word {
		// Рассчитываем индекс
		idx := ch - 'a'
		// Если такого узла нет возвращаем false
		if curr.children[idx] == nil {
			return false
		}
		curr = curr.children[idx]
	}
	return curr.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for _, ch := range prefix {
		// Рассчитываем индекс
		idx := ch - 'a'
		// Если такого узла нет возвращаем false
		if curr.children[idx] == nil {
			return false
		}
		curr = curr.children[idx]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
