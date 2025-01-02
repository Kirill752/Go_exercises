package main

// 121
func maxProfit(prices []int) int {
	i := 0
	n := len(prices)
	profit := 0
	best_buy := prices[0]
	for i < n {
		if prices[i] < best_buy {
			best_buy = prices[i]
		}
		profit = max(profit, prices[i]-best_buy)
		i++
	}
	return profit
}

// 122
func maxProfit_II(prices []int) int {
	// Маркер того, купоена акция или нет
	stock := false
	n := len(prices)
	profit := 0
	// Первый день
	if n > 1 && prices[1]-prices[0] > 0 { // покупка
		profit -= prices[0]
		stock = true
	}
	// После последнего дня цена не меняется
	prices = append(prices, prices[n-1])
	n++
	// Дни между первым и последним
	for i := 1; i < n-1; i++ {
		if (prices[i]-prices[i-1] <= 0) && (prices[i+1]-prices[i] >= 0) && i < n-2 && stock == false { // покупка
			profit -= prices[i]
			stock = true
		} else {
			if (prices[i]-prices[i-1] >= 0) && (prices[i+1]-prices[i] <= 0) && stock == true { // продажа
				profit += prices[i]
				stock = false
			}
		}
		// fmt.Printf("day: %d, profit = %d\n", i, profit)
	}
	// fmt.Println(prices)
	return profit
}
