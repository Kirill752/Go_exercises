package main

import (
	"fmt"
	"slices"
)

type truck struct {
	id       int
	start    int
	end      int
	capacity int
}

func planingOrder(n int, arraivalTable map[int][]int, m int, benzTrucks []truck) []int {
	// // Определяем временной промежуток поступления товаров
	time := make([]int, n)
	iter := 0
	for i := range arraivalTable {
		time[iter] = i
		iter++
	}
	slices.Sort(time)
	// Список машин на пункте погрузки
	onBase := []truck{}
	result := make([]int, n)
	// Начало симуляции
	// Одна итерация - один момент времени
	for _, t := range time {
		for i := 0; i < len(benzTrucks); i++ {
			// Добавляем в список машин на базе приезжающие машины
			if benzTrucks[i].start <= t && benzTrucks[i].end >= t {
				onBase = append(onBase, benzTrucks[i])
			}
		}
		// Сортируем машины в порядке увеличения времени прибытия
		slices.SortStableFunc(onBase, func(a truck, b truck) int {
			if a.start > b.start {
				return 1
			}
			if a.start == b.start {
				return 0
			}
			return -1
		})
		fmt.Println(onBase)
		// for k := range arraivalTable[t] {
		// 	for ind, car := range onBase {
		// 		if car.capacity > 0 {
		// 			currentCar := ind
		// 			choise := car.id
		// 			for j := ind; j < len(onBase) && onBase[j].start == car.start; j++ {
		// 				if onBase[j].id < choise {
		// 					choise = onBase[j].id
		// 					currentCar = j
		// 				}
		// 			}
		// 			onBase[currentCar].capacity--
		// 			result[arraivalTable[t][k]] = choise
		// 			break
		// 		}
		// 		if ind+1 == len(onBase) {
		// 			result[arraivalTable[t][k]] = -1
		// 		}
		// 	}
		// }
		for i := 0; i < len(arraivalTable[t]); i++ {
			// Если есть машины
			if len(onBase) > 0 {
				// Ищем свободную машину
				for j := 0; j < len(onBase); j++ {
					// Грузим товар в первую свободную машину
					if onBase[j].capacity > 0 {
						onBase[j].capacity--
						result[arraivalTable[t][i]] = onBase[j].id
						break
					}
					if j+1 == len(onBase) {
						result[arraivalTable[t][i]] = -1
					}
				}
			} else {
				result[arraivalTable[t][i]] = -1
			}
		}
		onBase = nil
	}
	// for t := Tstart; t <= Tend; t++ {
	// 	// Проверяем приезжающие машины
	// 	for i := 0; i < len(benzTrucks); i++ {
	// 		// Добавляем в список машин на базе приезжающие машины
	// 		if benzTrucks[i].start == t {
	// 			onBase = append(onBase, benzTrucks[i])
	// 		}
	// 	}
	// 	// Загружаем товар
	// 	if _, ok := arraivalTable[t]; ok {
	// 		// Проходимся по каждому вновь прибывшему товару
	// 		for i := 0; i < len(arraivalTable[t]); i++ {
	// 			// Если есть машины
	// 			if len(onBase) > 0 {
	// 				// Ищем свободную машину
	// 				for j := 0; j < len(onBase); j++ {
	// 					// Грузим товар в первую свободную машину
	// 					if onBase[j].capacity > 0 {
	// 						onBase[j].capacity--
	// 						result[arraivalTable[t][i]] = onBase[j].id
	// 						break
	// 					}
	// 					if j+1 == len(onBase) {
	// 						result[arraivalTable[t][i]] = -1
	// 					}
	// 				}
	// 			} else {
	// 				result[arraivalTable[t][i]] = -1
	// 			}
	// 		}
	// 	}
	// 	// Проверяем уезжающие машины
	// 	for i := 0; i < len(onBase); i++ {
	// 		// Если машине время уезжать, то удаляем ее из списка
	// 		if onBase[i].end <= t {
	// 			copy(onBase[i:], onBase[i+1:])
	// 			onBase = onBase[:len(onBase)-1]
	// 		}
	// 	}
	// }
	return result
}
