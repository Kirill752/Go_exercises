package main

import (
	"bufio"
	"fmt"
	"math"
	"net/http"
	"os"
	"strings"
)

// Ехрг представляет арифметическое выражение
type Expr interface {
	// Eval возвращает значение данного Ехрг в среде env.
	Eval(env Env) float64
	// Check сообщает об ошибках в данном Expr и добавляет свои Vars
	Check(vars map[Var]bool) error
	// String выводит арифметическое выражение
	String() string
}

// Var определяет переменную, например x.
type Var string

// literal представляет собой числовую константу, например 3.141.
type literal float64

// unary представляет выражение с унарным оператором, например -х.
type unary struct {
	op rune // '+' или '-'
	x  Expr
}

// binary представляет выражение с бинарным оператором, например х+у.
type binary struct {
	op   rune // '+', или '/'
	x, у Expr
}

// call представляет выражение вызова функции, например sin(x).
type call struct {
	fn   string
	args []Expr
}

func (v Var) String() string {
	return fmt.Sprintf("var: %s", string(v))
}

func (l literal) String() string {
	return fmt.Sprintf("const: %.2f", l)
}

func (u unary) String() string {
	return fmt.Sprintf("unary: %q%s", u.op, u.x.String())
}

func (b binary) String() string {
	return fmt.Sprintf("binary: %s %q %s", b.x.String(), b.op, b.у.String())
}

func (c call) String() string {
	switch c.fn {
	case "pow":
		return fmt.Sprintf("call: %s(%s, %s)", c.fn, c.args[0].String(), c.args[1].String())
	case "sin", "cos", "sqrt":
		return fmt.Sprintf("call: %s(%s)", c.fn, c.args[0].String())
	}
	panic(fmt.Sprintf("неподдерживаемая функция: %s", c.fn))
}

// Отображение переменных на значения
type Env map[Var]float64

func (v Var) Eval(env Env) float64 {
	return env[v]
}
func (v literal) Eval(_ Env) float64 {
	return float64(v)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("неподдерживаемый унарный оператор: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.у.Eval(env)
	case '-':
		return b.x.Eval(env) - b.у.Eval(env)
	case '*':
		return b.x.Eval(env) * b.у.Eval(env)
	case '/':
		return b.x.Eval(env) / b.у.Eval(env)
	}
	panic(fmt.Sprintf("неподдерживаемый бинарный оператор: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "cos":
		return math.Cos(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	case "min":
		min := c.args[0].Eval(env)
		for i := 0; i < len(c.args); i++ {
			if c.args[i].Eval(env) < min {
				min = c.args[i].Eval(env)
			}
		}
		return min
	}
	panic(fmt.Sprintf("неподдерживаемый вызов функции: %s", c.fn))
}

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}
func (literal) Check(vars map[Var]bool) error {
	return nil
}

func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("некоректный унарный оператор %q", u.op)
	}
	return u.x.Check(vars)
}
func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("некоректный бинарный оператор %q", b.op)
	}
	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.у.Check(vars)
}

var numParams = map[string]int{"pow": 2, "sin": 1, "cos": 1, "sqrt": 1}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("неизвестная функция %q", c.fn)
	}
	if c.fn == "min" {
		if len(c.args) <= 0 {
			return fmt.Errorf("вызов %s не имеет аргументов", c.fn)
		} else {
			return nil
		}
	}
	if len(c.args) != arity {
		return fmt.Errorf("вызов %s имеет %d аргументов вместо %d",
			c.fn, len(c.args), arity)
	}
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

func parseAndCheck(s string) (Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}
	expr, err := Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	// for v := range vars {
	// 	if v != "x" && v != "y" && v != "r" {
	// 		return nil, fmt.Errorf("undefined variable: %s", v)
	// 	}
	// }
	return expr, nil
}

func plot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "некоректное выражение: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	plot3D(w, func(x, y float64) float64 {
		r := math.Hypot(x, y)
		return expr.Eval(Env{"x": x, "y": y, "r": r})
	})
}

func main() {
	// http.HandleFunc("/plot", plot)
	// log.Fatal(http.ListenAndServe("localhost:8000", nil))
	in := bufio.NewReader(os.Stdin)
	eq, _, _ := in.ReadLine()
	expr, err := parseAndCheck(string(eq))
	if err != nil {
		fmt.Printf("ошибка парсинга: %v\n", err)
	}
	env := Env{"a": math.Pi / 6, "y": 2, "z": 16}
	fmt.Println(expr.Eval(env))
}
