package main

type Trie struct {
	// Каждый индекс отражает цифру на телефоне
	children [28]*Trie
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

func getLetters(digit byte) string {
	switch digit {
	case '2':
		return "abc"
	case '3':
		return "def"
	case '4':
		return "ghi"
	case '5':
		return "jkl"
	case '6':
		return "mno"
	case '7':
		return "pqrs"
	case '8':
		return "tuv"
	case '9':
		return "wxyz"
	default:
		return ""
	}
}

func decode(T9word string, t *Trie, idx int, currentWord string, results *[]string) {
	if idx == len(T9word) {
		if t.Search(currentWord) {
			*results = append(*results, currentWord)
		}
		return
	}
	// Текущая цифра
	curr := T9word[idx]
	// Количестово её повторений
	count := 0
	for i := idx; i < len(T9word) && T9word[i] == curr; i++ {
		count++
	}
	// Возможные буквы для текущей цифры
	letters := getLetters(curr)
	if len(letters) == 0 {
		return
	}
	// определим, какая буква соответсвует такому количеству нажатий
	i := (count - 1) % len(letters)
	for j := 0; j <= i; j++ {
		letter := letters[j]
		decode(T9word, t, idx+j+1, currentWord+string(letter), results)
	}
}

// func main() {
// 	t := ConstructorTrie()
// 	t.Insert("hello")
// 	t.Insert("world")
// 	t.Insert("quantum")
// 	t.Insert("physics")
// 	res := []string{}
// 	// t9 := "7788266888674499977774442227777"
// 	t9 := "443355555566696667775553"
// 	var w strings.Builder
// 	for i := range len(t9) {
// 		w.WriteByte(t9[i])
// 		l := len(res)
// 		decode(w.String(), &t, 0, "", &res)
// 		if len(res) > l {
// 			w.Reset()
// 		}
// 	}
// 	for _, w := range res {
// 		fmt.Printf("%s ", strings.ToUpper(w))
// 	}
// 	fmt.Println()
// }
