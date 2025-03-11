package main

type WordDictionary struct {
	// Каждый индекс отражает букву английского алфавита
	children [26]*WordDictionary
	isEnd    bool
}

func ConstructorW() WordDictionary {
	newTrie := WordDictionary{}
	return newTrie
}

func (this *WordDictionary) AddWord(word string) {
	curr := this
	for _, ch := range word {
		// Рассчитываем индекс
		idx := ch - 'a'
		// Если такого узла нет
		// Создаем новый узел и переходим в него
		if curr.children[idx] == nil {
			n := WordDictionary{}
			curr.children[idx] = &n
		}
		curr = curr.children[idx]
	}
	// Обозначаем конец слова
	curr.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	curr := this
	res := false
	for i, ch := range word {
		// Если ch = '.', то нужно проверить все возвможные случаи
		if ch == '.' {
			// Прохожусь по все дочерним узлам
			for j := range len(curr.children) {
				// для каждого запускаю поиск подстроки, следующей за точкой
				if curr.children[j] != nil {
					res = res || curr.children[j].Search(word[i+1:])
				}
				if res {
					return res
				}
			}
			return res
		}
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

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
