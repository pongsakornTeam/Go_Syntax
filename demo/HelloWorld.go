package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Golang")
	defineVariable()
	defineList()
	defineSlice()
	defineMap()
	plus(1, 2)
	fmt.Println("outside function : ", plus(1, 2))
	fmt.Println(validate(2, 1))
	fmt.Println(testLoop(5))
}

func defineVariable() {
	var number int = 0
	var text string = "hello"
	float := 12.1
	fmt.Println(number)
	fmt.Println(float)
	fmt.Println(text)
}

func defineList() {
	var listText [4]string
	listText[0] = "1"
	fmt.Println(listText[0])
}

func defineSlice() {
	var sliceText []string
	sliceText = []string{"hello", "slice"}
	fmt.Println(sliceText)
	sliceText = append(sliceText, "newslice")
	fmt.Println(sliceText)
	newSliceText := sliceText[2:]
	fmt.Println(newSliceText)
}

func defineMap() {
	var mapText = make(map[string]string)
	fmt.Println("mapText", mapText)
	mapText["price1"] = "10"
	fmt.Println("mapText", mapText)
}

func plus(a int, b int) int {
	result := a + b
	fmt.Println("inside function :", result)
	return result
}

func validate(a int, b int) bool {
	var result bool = false
	if a > b {
		result = true
	} else {
		result = false
	}
	return result
}

func testLoop(count int) int {
	var result int = 0
	for i := 0; i < count; i++ {
		fmt.Print(i, " ")
		result += i
	}
	fmt.Println()
	return result
}
