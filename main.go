package main

import (
	"fmt"
)

func main() {
	//increasePattern(5)
	//decreasePattern(5)
	drawThefuckingTriangle(11)
}

func drawThefuckingTriangle(n int) {
	increaseHollowPattern(n/2 + n%2)
	decreaseHollowPattern(n / 2)
}

func increasePattern(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("* ")
		}
		fmt.Printf("\n")
	}
}
func increaseHollowPattern(n int) {
	for i := 0; i < n; i++ {
		space := ""
		for j := 0; j <= i; j++ {
			if j == i {
				fmt.Printf("%s*", space)
			}
			space += " "
		}
		fmt.Printf("\n")
	}
}

func decreasePattern(n int) {
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			fmt.Printf("* ")
		}
		fmt.Printf("\n")
	}
}
func decreaseHollowPattern(n int) {
	for i := 0; i < n; i++ {
		space := ""
		for j := i; j < n-1; j++ {
			//	fmt.Printf("%d", j)
			space += " "
		}
		fmt.Printf("%s*\n", space)
	}
}
