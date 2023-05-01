package main

import (
	"fmt"
)

func add(x float64, y float64) float64 {
	return x + y
}

func subtract(x float64, y float64) float64 {
	return x - y
}

func multiply(x float64, y float64) float64 {
	return x * y
}

func divide(x float64, y float64) (float64, error) {
	if y == 0 {
		return 0, fmt.Errorf("除数不能为零")
	}
	return x / y, nil
}

func main() {
	var x, y float64
	var op string

	for {
		fmt.Printf("请输入第一个数字：")
		if _, err := fmt.Scanln(&x); err != nil {
			fmt.Println("输入不合法，请重新输入！")
			continue
		}
		break
	}

	for {
		fmt.Printf("请输入第二个数字：")
		if _, err := fmt.Scanln(&y); err != nil {
			fmt.Println("输入不合法，请重新输入！")
			continue
		}
		break
	}

	for {
		fmt.Printf("请输入运算符（+、-、*、/）：")
		if _, err := fmt.Scanln(&op); err != nil {
			fmt.Println("输入不合法，请重新输入！")
			continue
		}
		if op != "+" && op != "-" && op != "*" && op != "/" {
			fmt.Println("输入不合法，请重新输入！")
			continue
		}
		break
	}

	var rslt float64
	var err error

	switch op {
	case "+":
		rslt = add(x, y)
	case "-":
		rslt = subtract(x, y)
	case "*":
		rslt = multiply(x, y)
	case "/":
		rslt, err = divide(x, y)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Printf("%.2f %s %.2f = %.2f\n", x, op, y, rslt)
}