package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	numArabicBM = map[rune]int{48: 0, 49: 1, 50: 2, 51: 3, 52: 4, 53: 5, 54: 6, 55: 7, 56: 8, 57: 9}
	numRomanBM  = map[rune]int{73: 1, 86: 5, 88: 10}
)

type expr interface {
	calc() error
	output()
}

type exprA struct {
	inpS string
	res  int
}

type exprR struct {
	inpS string
	res  int
}

func input() (expr, error) {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.TrimSpace(s)
	if _, ok := numArabicBM[rune(s[0])]; ok {
		return &exprA{inpS: s}, nil
	} else if _, ok := numRomanBM[rune(s[0])]; ok {
		return &exprR{inpS: s}, nil
	}
	return nil, errors.New("expression is incorrect (input)")
}

func (e *exprA) calc() error {
	var (
		a, b int
		oper rune
	)
L:
	for i, v := range e.inpS {
		d, ok := numArabicBM[v]
		switch {
		case ok:
			b = b*10 + d
		case v == '+' && oper == 0:
			oper = v
			a, b = b, 0
		case v == '-' && oper == 0:
			oper = v
			a, b = b, 0
		case v == '*' && oper == 0:
			oper = v
			a, b = b, 0
		case v == '/' && oper == 0:
			oper = v
			a, b = b, 0
		case v == 13:
			break L
		case v == ' ' && (e.inpS[i+1] == 42 || e.inpS[i+1] == 43 || e.inpS[i+1] == 45 || e.inpS[i+1] == 47 ||
			e.inpS[i-1] == 42 || e.inpS[i-1] == 43 || e.inpS[i-1] == 45 || e.inpS[i-1] == 47):
			continue
		case v == ' ':
			return errors.New("expression is incorrect (calc, space trouble)")
		default:
			return errors.New("expression is incorrect (calc, invalid char)")
		}
		if b > 10 {
			return errors.New("exceeding value (calc, over 10)")
		}
	}

	switch oper {
	case '+':
		e.res = a + b
	case '-':
		e.res = a - b
	case '*':
		e.res = a * b
	case '/':
		e.res = a / b
	}
	return nil
}

func (e *exprR) calc() error {
	var (
		a, b int
		oper rune
	)
L:
	for i, v := range e.inpS {
		d1, ok := numRomanBM[v]
		switch {
		case ok:
			if i < len(e.inpS)-1 {
				d2, ok := numRomanBM[rune(e.inpS[i+1])]
				if ok && d1 < d2 {
					b = b - d1
				} else {
					b = b + d1
				}
			} else {
				b = b + d1
			}
		case v == '+' && oper == 0:
			oper = v
			a, b = b, 0
		case v == '-' && oper == 0:
			oper = v
			a, b = b, 0
		case v == '*' && oper == 0:
			oper = v
			a, b = b, 0
		case v == '/' && oper == 0:
			oper = v
			a, b = b, 0
		case v == 13:
			break L
		case v == ' ' && (e.inpS[i+1] == 42 || e.inpS[i+1] == 43 || e.inpS[i+1] == 45 || e.inpS[i+1] == 47 ||
			e.inpS[i-1] == 42 || e.inpS[i-1] == 43 || e.inpS[i-1] == 45 || e.inpS[i-1] == 47):
			continue
		case v == ' ':
			return errors.New("expression is incorrect (calc, space trouble)")
		default:
			return errors.New("expression is incorrect (calc, invalid char)")
		}
		if b > 10 {
			return errors.New("exceeding value (calc, over X)")
		}
	}

	switch oper {
	case '+':
		e.res = a + b
	case '-':
		e.res = a - b
		if e.res < 1 {
			return errors.New("output of the result is not possible (calc, zero or negative numb)")
		}
	case '*':
		e.res = a * b
	case '/':
		e.res = a / b
	}
	return nil
}

func (e *exprA) output() {
	fmt.Println(e.res)
}
func (e *exprR) output() {
	if e.res == 100 {
		fmt.Println("C")
		return
	}

	var (
		r, o int = e.res, 1
		p    string
	)
	SS := map[int]string{1: "I", 4: "IV", 5: "V", 9: "IX", 10: "X", 40: "XL", 50: "L", 90: "XC"}

	for b := r; b/10 > 0; b /= 10 {
		o *= 10
	}
	for r > 0 {
		b := r / o
		if b > 0 {
			switch {
			case b == 9 || b == 5 || b == 4:
				p = p + SS[b*o]
				r = r - b*o
			case b > 5:
				p = p + SS[5*o]
				r = r - 5*o
			case b < 4:
				for b > 0 {
					p = p + SS[1*o]
					r = r - 1*o
					b--
				}
			}
		} else {
			o /= 10
		}
	}
	fmt.Println(p)
}

func main() {
	var ex expr
	ex, err := input()
	if err == nil {
		err = ex.calc()
		if err == nil {
			ex.output()
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}

}
