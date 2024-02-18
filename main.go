package main

import (
	"fmt"
	"log"
)

func celsius (temp_f int) int {
	temp_c := (temp_f -32) * 5/9
	return temp_c
}

func fahrenheit (temp_c int) int {
	temp_f := (temp_c * 9/5) + 32 
	return temp_f
}

func main() {
	fmt.Println("Temp Calculator")
	fmt.Println("Enter 'f' for Fahrenheit or 'c' for Celsius: ")
	var mode string
	_, err := fmt.Scan(&mode)
	if err != nil {
		log.Fatal(err)
	}
	switch mode {
	case "f":
		fmt.Println("Type in the temperature in Celsius: ")
		var temp_c int
		fmt.Scan(&temp_c)
		temp_f := fahrenheit(temp_c)
		fmt.Println("Temperature in Fahrenheit: ", temp_f)
	case "c":
		fmt.Println("Type in the temperature in Fahrenheit: ")
		var temp_f int
		fmt.Scan(&temp_f)
		temp_c := celsius(temp_f)
		fmt.Println("Temperatur in Celsius: ", temp_c)
	default:
		fmt.Println("Undefined input")
	}
}
