package main

import (
	"fmt"
)

func maxValue(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("cписок значений пуст")
	}
	maxV := vals[0]
	for _, v := range vals {
		if v > maxV {
			maxV = v
		}
	}
	return maxV, nil
}

func minValue(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("cписок значений пуст")
	}
	maxV := vals[0]
	for _, v := range vals {
		if v < maxV {
			maxV = v
		}
	}
	return maxV, nil
}
