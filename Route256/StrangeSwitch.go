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

func dfs(sw *switcher, maxComp int, countSwitch int, usedSwitch []*switcher) (bool, int, int, []*switcher) {
	if sw.used {
		return false, maxComp, countSwitch, usedSwitch
	}
	sw.used = true
	// Смотрим оба компьютера, к которым подключен переключатель
	// Если переключатель включит два компьютера то дергаем рычаг
	if !sw.a.state && !sw.b.state {
		sw.a.state, sw.b.state = true, true
		maxComp += 2
		countSwitch++
		usedSwitch = append(usedSwitch, sw)
		return true, maxComp, countSwitch, usedSwitch
	}
	// Если рычаг выключит 2 компьютера
	if sw.a.state && sw.b.state {
		return false, maxComp, countSwitch, usedSwitch
	}
	// Если рычаг выключит компьютер a, но включит b
	if sw.a.state && !sw.b.state {
		// Пробуем переключить и смотрим, что будет в перспективе
		sw.a.state = false
		sw.b.state = true
		// Нам выгодно в данный момент переключить ключ, если
		// при его переключении мы в перспективе получим больше включенных компьютеров,
		// то есть, если среди рычагов, подключенных к включенному компьютеру, найдется хотя бы 1 еще не использованный,
		// который в будующем включит этот компьютер
		for i := 0; i < len(sw.a.switches); i++ {
			// Если соседний ключ не поcещался
			// Решаем, выгодно ли нам его включить
			// Выгодно, если он включит обратно a и не выключит b
			if !sw.a.switches[i].used {
				res, maxComp, countSwitch, usedSwitch := dfs(sw.a.switches[i], maxComp, countSwitch, usedSwitch)
				countSwitch++
				usedSwitch = append(usedSwitch, sw)
				return res, maxComp, countSwitch, usedSwitch
			}
		}
		// Если не нашлось такого, то возвращаем все как было
		sw.a.state = true
		sw.b.state = false
	}
	if !sw.a.state && sw.b.state {
		sw.a.state = true
		sw.b.state = false
		for i := 0; i < len(sw.b.switches); i++ {
			// Если соседний ключ не почещался
			if !sw.b.switches[i].used {
				res, maxComp, countSwitch, usedSwitch := dfs(sw.b.switches[i], maxComp, countSwitch, usedSwitch)
				countSwitch++
				usedSwitch = append(usedSwitch, sw)
				return res, maxComp, countSwitch, usedSwitch
			}
		}
		// Если не нашлось такого, то возвращаем все как было
		sw.a.state = false
		sw.b.state = true
	}
	return false, maxComp, countSwitch, usedSwitch
}

func StrangeSwitch(switches []switcher) (int, int, []*switcher) {
	maxComps := 0
	countSwitches := 0
	usedSwitches := []*switcher{}
	for i := 0; i < len(switches); i++ {
		_, nowComps, nowcountSwitches, nowusedSwitches := dfs(&switches[i], maxComps, countSwitches, usedSwitches)
		if nowComps >= maxComps {
			maxComps = nowComps
			countSwitches = nowcountSwitches
			usedSwitches = nowusedSwitches
		}
	}
	return maxComps, countSwitches, usedSwitches
}
