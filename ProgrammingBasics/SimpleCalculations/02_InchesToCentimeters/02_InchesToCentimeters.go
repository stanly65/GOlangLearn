package main

import "fmt"

func main() {
	var num float64
	fmt.Scan(&num)
	var inch = 2.54
	var cm = num * inch
	println(cm)
}
