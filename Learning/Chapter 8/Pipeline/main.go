package main

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// 3 вариант
	go counter(naturals)
	go squarer(naturals, squares)
	printer(squares)

	// // Генерация
	// // 1 и 2 варианты
	// go func() {
	// 	for x := 0; x < 5; x++ {
	// 		naturals <- x
	// 	}
	// 	close(naturals)
	// }()

	// // Возведение в квадрат
	// go func() {
	// 	// 1 вариант
	// 	// for {
	// 	// 	x, ok := <-naturals
	// 	// 	if !ok {
	// 	// 		break // Канал закрыт и опусташен
	// 	// 	}
	// 	// 	squars <- x * x
	// 	// }

	// 	// 2 вариант
	// 	for x := range naturals {
	// 		squares <- x * x
	// 	}
	// 	close(squares)
	// }()

	// // Печать
	// // 1 вариант
	// // for {
	// // 	x, ok := <-squares
	// // 	if !ok {
	// // 		break // Канал закрыт и опусташен
	// // 	}
	// // 	fmt.Printf("%v ", x)
	// // 	time.Sleep(1 * time.Second)
	// // }
	// // 2 вариант
	// for x := range squares {
	// 	fmt.Printf("%v ", x)
	// 	time.Sleep(1 * time.Second)
	// }
	// fmt.Println()
}
