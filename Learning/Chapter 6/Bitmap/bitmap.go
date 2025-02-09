package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// IntSet представляет собой множество небольших неотрицательных
// целых чисел. Нулевое значение представляет пустое множество.
type IntSet struct {
	words []uint
}

// Has указывает, содержит ли множество неотрицательное значение х.
// word - машинное слово
// Машинное слово есть фрагмент данных фиксированного размера,
// обрабатываемый как единое целое с помощью набора команд или аппаратного обеспечения процессора.
// Количество бит в машинном слове — размер слова (он же ширина или длина слова) —
// является важной характеристикой любой конкретной архитектуры процессора или компьютерной архитектуры.
func (s *IntSet) Has(x int) bool {
	wordSize := 32 << (uint(0) >> 63)
	word, bit := x/wordSize, uint(x%wordSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add добавляет неотрицательное значение x в множество.
// Пример: words хранит числа общим размером 64 бита. Поэтому при добавлении x. Число x  будет начинаться с x/64 индекс.
// мы вставляем этому числу x%64 бит.
func (s *IntSet) Add(x int) {
	wordSize := 32 << (uint(0) >> 63)
	word, bit := x/wordSize, uint(x%wordSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith делает множество s равным объединению множеств s и t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String возвращает множество как строку вида "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	wordSize := 32 << (uint(0) >> 63)
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < wordSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", wordSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Возвращает количество элементов (6.1)
func (s *IntSet) Len() int {
	return len(s.words)
}

// Удаляет x из множества (6.1)
func (s *IntSet) Remove(x int) {
	wordSize := 32 << (uint(0) >> 63)
	word, bit := x/wordSize, uint(x%wordSize)
	if word < len(s.words) {
		s.words[word] ^= 1 << bit
	}
}

// Удаляет все элементы множества (6.1)
func (s *IntSet) Clear() {
	s.words = nil
}

// Возвращает копию множества (6.1)
func (s *IntSet) Copy() *IntSet {
	newS := new(IntSet)
	newS.words = make([]uint, s.Len())
	copy(newS.words, s.words)
	return newS
}

// Добавляет нсколько элементов (6.2)
func (s *IntSet) AddAll(nums ...int) {
	for _, v := range nums {
		s.Add(v)
	}
}

// IntersectWith делает множество s равным пересечению множеств s и t. (6.3)
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}

// DifferenceWith делает множество s равным разности множеств s и t.(6.3)
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		}
	}
}

// SymetricDifferenceWith делает множество s равным симметричной разности множеств s и t.
// Симметричная разность двух множеств содержит эле­менты, имеющиеся в одном из множеств, но не в обоих одновременно. (6.3)
func (s *IntSet) SymetricDifferenceWith(t *IntSet) {
	intersect := s.Copy()
	intersect.IntersectWith(t)
	s.UnionWith(t)
	s.DifferenceWith(intersect)
}

// Elems возвращает срез, создержащий элементы множества (6.4)
func (s *IntSet) Elems() []int {
	str := s.String()[1 : len(s.String())-1]
	numsStr := strings.Split(str, " ")
	nums := make([]int, len(numsStr))
	for i, v := range numsStr {
		nums[i], _ = strconv.Atoi(v)
	}
	return nums
}
