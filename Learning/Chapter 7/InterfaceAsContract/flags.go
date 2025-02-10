package main

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

// FToC преобразует температуру по Фаренгейту в температуру по Цельсию.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }

// KToC преобразует температуру по Кельвину в температуру по Цельсию
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// Реализуем метод String для всех типов
func (c Celsius) String() string { return fmt.Sprintf("%.f°C", c) }

// *celsiusFlag соответствует интерфейсу flag.Value
type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("неверная температура %q", s)
}

// CelsiusFlag определяет флаг Celsius с указанным именем, значением
// по умолчанию и строкой-инструкцией по применению и возвращает адрес
// переменной-флага. Аргумент флага должен содержать числовое значение
// и единицу измерения, например "100С".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{Celsius: value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
