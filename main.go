package main

import (
	"fmt"
	"bufio"
	"os"
)	

func main() {
	fmt.Println("Temp Calculator")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter 'f' for Fahrenheit or 'c' for Celsius: ")
	mode, _ := reader.ReadString('\n')
	fmt.Println("Mode: ", mode)

}
