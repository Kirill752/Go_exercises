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
		// fmt.Printf("n = %d\n", n)
		// fmt.Printf("m = %d\n", m)
		// _, _ = in.ReadByte()
		// Создаем n объектов компьютеров
		computers := []computer{}
		// Множество использованных id для компьютеров
		usedIds := map[int]bool{}
		// Инициализируем рычаги
		switches := make([]switcher, m)
		// Считываем a_i b_i
		for i := 0; i < m; i++ {
			var ai int
			var bi int
			var ptr_a *computer
			var ptr_b *computer
			fmt.Fscan(in, &ai, &bi)
			_, _ = in.ReadByte()
			// Если такого компьютера еще не было,
			// создаем объект компьютера
			if _, ok := usedIds[ai]; !ok {
				usedIds[ai] = true
				computers = append(computers, computer{id: ai, switches: make([]*switcher, 0)})
				computers[len(computers)-1].switches = append(computers[len(computers)-1].switches, &switches[i])
				ptr_a = &computers[len(computers)-1]
			}
			if _, ok := usedIds[bi]; !ok {
				usedIds[bi] = true
				computers = append(computers, computer{id: bi, switches: make([]*switcher, 0)})
				computers[len(computers)-1].switches = append(computers[len(computers)-1].switches, &switches[i])
				ptr_b = &computers[len(computers)-1]
			}
			if ptr_a == nil {
				for j := 0; j < len(computers); j++ {
					if ai == computers[j].id {
						ptr_a = &computers[j]
						computers[j].switches = append(computers[j].switches, &switches[i])
						break
					}
				}
			}
			if ptr_b == nil {
				for j := 0; j < len(computers); j++ {
					if bi == computers[j].id {
						ptr_b = &computers[j]
						computers[j].switches = append(computers[j].switches, &switches[i])
						break
					}
				}
			}
			switches[i].id = i + 1
			switches[i].a = ptr_a
			switches[i].b = ptr_b
		}
		fmt.Println("Computers:")
		for i := 0; i < len(computers); i++ {
			computers[i].Print()
		}
		fmt.Println()
		fmt.Println("Switches:")
		for i := 0; i < len(switches); i++ {
			switches[i].Print()
		}
		fmt.Println()
		maxComps, countSwitches, usedSwitches := StrangeSwitch(switches)
		fmt.Fprintf(out, "%d\n", maxComps)
		fmt.Fprintf(out, "%d\n", countSwitches)
		for _, v := range usedSwitches {
			fmt.Fprintf(out, "%d ", v.id)
		}
		fmt.Fprintf(out, "\n")
	}
}
