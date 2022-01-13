package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("A simple calculator")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Value 1: ")
	input1, _ := reader.ReadString('\n')
	fmt.Println("Value 2: ")
	input2, _ := reader.ReadString('\n')
	value1, err1 := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	value2, err2 := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err1 != nil {
		fmt.Println(err1)
		panic("Value 1 must be a number")
	} else if err2 != nil {
		fmt.Println(err2)
		panic("Value 2 must be a number")
	} else {
		floatSum := value1 + value2
		fmt.Println("The sum of", value1, "+", value2, "is", floatSum)
	}
	//panic("This is my error message.")
}
