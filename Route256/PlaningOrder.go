package main

import (
	"slices"
)

type truck struct {
	id       int
	start    int
	end      int
	capacity int
}

type product struct {
	id   int
	time int
	car  int
}

func planingOrder(arraival []product, benzTrucks []truck) {
	// Сортируем продукты по времени прибытия
	slices.SortFunc(arraival, func(a product, b product) int {
		if a.time > b.time {
			return 1
		}
		if a.time == b.time {
			return 0
		}
		return -1
	})
	// Сортируем машины по времени прибытия и по индексу
	slices.SortStableFunc(benzTrucks, func(a truck, b truck) int {
		if a.start > b.start {
			return 1
		}
		if a.start == b.start {
			if a.id > b.id {
				return 1
			}
			if a.id == b.id {
				return 0
			}
			return -1
		}
		return -1
	})
	nowProdIndex := 0
	// Смотрим в какую машину какой заказ положить
	for _, car := range benzTrucks {
		// Проходимся по временам прибытия товаров
		for i := nowProdIndex; i < len(arraival) && arraival[i].time <= car.end && car.capacity > 0; i++ {
			// Если в машину еще можно положить товар
			if arraival[i].time >= car.start {
				// Кладем его в машину
				car.capacity--
				arraival[i].car = car.id
				// Сдвигаем индекс времени, так как уже погрузили товар
				nowProdIndex = i + 1
			}
		}
	}
	slices.SortStableFunc(arraival, func(a product, b product) int {
		if a.id > b.id {
			return 1
		}
		if a.id == b.id {
			return 0
		}
		return -1
	})
}
