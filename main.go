package main

import (
	"fmt"
	"os"
)

func main() {
	os.Exit(calculate())

}

func calculate() int {
	width, height := float64(0), float64(0)
	costPerTile := float64(0)
	fmt.Print("enter cost per tile: ")
	fmt.Scanln(&costPerTile)
	fmt.Print("enter width: ")
	fmt.Scanln(&width)
	fmt.Print("enter height: ")
	fmt.Scanln(&height)
	fmt.Println(costPerTile * height * width)
	return 0
}
