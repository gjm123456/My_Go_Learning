package main

import (
	"fmt"
	"math"
)

//包内变量
var (
	p1 = 1
	p2 = "kkk"
	p3 = false
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q \n", a, s)
}

func varilableInitialValue() {
	var a, c int = 3, 4
	var b = "abc"
	fmt.Println(a, b, c)
}

func variableTypeDeduction() {
	var a, c, b = 3, true, "abc"
	fmt.Println(a, b, c)
}

func variableShorter() {
	a, c, b := 3, true, "abc"
	b = "qw"
	fmt.Println(a, b, c)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	//c = math.Sqrt(a*a + b*b) Sqrt需要传float
	//c = math.Sqrt(float64(a*a + b*b)) Sqrt返回float, c必为int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {

}

func main() {
	fmt.Println("Hello World")
	variableZeroValue()
	varilableInitialValue()
	variableTypeDeduction()
	variableShorter()
}
