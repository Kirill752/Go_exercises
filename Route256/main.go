package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscanf(in, "%d", &t)
	_, _ = in.ReadByte()
	for ; t > 0; t-- {
		var n int
		var m int
		fmt.Fscanf(in, "%d %d", &n, &m)
		// Создаем n объектов компьютеров
		computers := make([]*computer, 0)
		// Множество использованных id для компьютеров
		usedIds := map[int]bool{}
		// Инициализируем рычаги
		switches := make([]switcher, m)
		// Считываем a_i b_i
		for i := 0; i < m; i++ {
			var ai int
			var bi int
			fmt.Fscan(in, &ai, &bi)
			_, _ = in.ReadByte()
			// Новый переключатель
			switches[i].id = i + 1
			// Создаем компьютер a, если компьютера с таким id еще нет
			if _, ok := usedIds[ai]; !ok {
				a := new(computer)
				a.id = ai
				a.switches = append(a.switches, &switches[i])
				switches[i].a = a
				usedIds[ai] = true
				computers = append(computers, a)
			} else {
				// Иначе ищем этот компьютер
				for j := 0; j < len(computers); j++ {
					if computers[j].id == ai {
						computers[j].switches = append(computers[j].switches, &switches[i])
						switches[i].a = computers[j]
					}
				}
			}
			// Создаем компьютер b, если компьютера с таким id еще нет
			if _, ok := usedIds[bi]; !ok {
				b := new(computer)
				b.id = bi
				b.switches = append(b.switches, &switches[i])
				switches[i].b = b
				usedIds[bi] = true
				computers = append(computers, b)
			} else {
				// Иначе ищем этот компьютер
				for j := 0; j < len(computers); j++ {
					if computers[j].id == bi {
						computers[j].switches = append(computers[j].switches, &switches[i])
						switches[i].b = computers[j]
					}
				}
			}
		}
		maxComps, countSwitches, usedSwitches := StrangeSwitch(switches)
		fmt.Fprintf(out, "%d\n", maxComps)
		fmt.Fprintf(out, "%d\n", countSwitches)
		for _, v := range usedSwitches {
			fmt.Fprintf(out, "%d ", v.id)
		}
		fmt.Fprintf(out, "\n")
	}
}
