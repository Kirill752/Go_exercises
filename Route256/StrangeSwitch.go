package main

import "fmt"

type computer struct {
	id int
	// false - выключен
	// true - включен
	state bool
	// Список переключателей привязанных к данному компьютеру
	switches []*switcher
}

func (comp computer) Print() {
	fmt.Printf("{%d, %v, [", comp.id, comp.state)
	for i := 0; i < len(comp.switches)-1; i++ {
		fmt.Printf("%d ", comp.switches[i].id)
	}
	fmt.Printf("%d]} ", comp.switches[len(comp.switches)-1].id)
}

type switcher struct {
	id   int
	used bool
	a    *computer
	b    *computer
}

func (sw switcher) Print() {
	fmt.Printf("{%d, %v, [", sw.id, sw.used)
	fmt.Printf("[a: %d, b: %d]", sw.a.id, sw.b.id)
}

func StrangeSwitch(switches []switcher) (int, int, []switcher) {
	maxComps := 0
	countSwitches := 0
	usedSwitches := []switcher{}
	for i := 0; i < len(switches); i++ {
		// Если рычаг вкючит 2 компютера то дергаем
		if !switches[i].a.state && !switches[i].b.state {
			switches[i].used = true
			switches[i].a.state = true
			switches[i].b.state = true
			maxComps += 2
			countSwitches++
			usedSwitches = append(usedSwitches, switches[i])
			// continue
		}
		// // Если рычаг вкючит компьютер a, но выключит b,
		// // То количество включенных компьютеров не изменится
		// // Можем дернуть, если существует еще не использованный рычаг,
		// // который включит компьютер b обратно
		// // Соответсвенно дергаем тот рычаг тоже
		// if !switches[i].a.state && switches[i].b.state {
		// 	// Дергаем исходный рычаг
		// 	switches[i].a.state = true
		// 	switches[i].b.state = false
		// 	// Ищем еще не использованный рычаг, который включит
		// 	// b обратно
		// 	for i, sw := range switches[i].b.switches {
		// 		if !sw.used {
		// 			// Дергаем его
		// 			sw.used = true
		// 			sw.a.state = !sw.a.state
		// 			sw.b.state = !sw.b.state
		// 			countSwitches += 2
		// 			usedSwitches = append(usedSwitches, switches[i])
		// 			usedSwitches = append(usedSwitches, *sw)
		// 			// Определим какой из двух компьютеров, принадлежащих этому переключателю
		// 			// мы включили
		// 			maxComps++
		// 			// Если это компьютер а
		// 			if sw.a == switches[i].b {
		// 				if sw.b.state {
		// 					maxComps++
		// 				}
		// 			} else {
		// 				if sw.a.state {
		// 					maxComps++
		// 				}
		// 			}
		// 			break
		// 		}
		// 		// Если мы не нашли такого переключателя, то
		// 		// не дергаем исходный рычаг
		// 		if i == len(switches[i].b.switches)-1 {
		// 			// возвращаем все как было
		// 			switches[i].a.state = false
		// 			switches[i].b.state = true
		// 		}
		// 	}
		// }
		// // Аналогично если переключатель выключит а, но
		// // включит b
		// if switches[i].a.state && !switches[i].b.state {
		// 	// Дергаем исходный рычаг
		// 	switches[i].b.state = true
		// 	switches[i].a.state = false
		// 	// Ищем еще не использованный рычаг, который включит
		// 	// a обратно
		// 	for i, sw := range switches[i].a.switches {
		// 		if !sw.used {
		// 			// Дергаем его
		// 			sw.used = true
		// 			sw.b.state = !sw.b.state
		// 			sw.a.state = !sw.a.state
		// 			countSwitches += 2
		// 			usedSwitches = append(usedSwitches, switches[i])
		// 			usedSwitches = append(usedSwitches, *sw)
		// 			// Определим какой из двух компьютеров, принадлежащих этому переключателю
		// 			// мы включили
		// 			maxComps++
		// 			// Если это компьютер b
		// 			if sw.b == switches[i].a {
		// 				if sw.b.state {
		// 					maxComps++
		// 				}
		// 			} else {
		// 				if sw.a.state {
		// 					maxComps++
		// 				}
		// 			}
		// 			break
		// 		}
		// 		// Если мы не нашли такого переключателя, то
		// 		// не дергаем исходный рычаг
		// 		if i == len(switches[i].a.switches)-1 {
		// 			// возвращаем все как было
		// 			switches[i].b.state = false
		// 			switches[i].a.state = true
		// 		}
		// 	}
		// }
	}
	return maxComps, countSwitches, usedSwitches
}
