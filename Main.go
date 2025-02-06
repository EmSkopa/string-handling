/*
File Name:  Main.go
Copyright:  2022
Author:     Emir Skopljak
*/

package main

import (
	"fmt"
)

func exampleOne() {
	fmt.Println(generateRandomString(false))
	fmt.Println(generateRandomString(true))

	if result, err := storyStats(generateRandomString(false)); err == nil {
		fmt.Println(result)
	}
	if result, err := storyStats(generateRandomString(true)); err == nil {
		fmt.Println(result)
	}
}

func exampleTwo() {
	list := []int{5, 9, 7, 11}
	fmt.Println(findMaxSum(list))
}

func exampleThree() {
	expensiveFunction := func() int { return exampleFunction(200000000) }
	cheapFunction := func() int { return exampleFunction(10000000) }

	ch := fillChannel(expensiveFunction, cheapFunction)

	if ch != nil {
		for i := 0; i < 2; i++ {
			fmt.Printf("Result: %d\n", <-ch)
		}
	}
}

func exampleFour() {
	var fireDragon FireDragon
	var egg = fireDragon.Lay()
	var childDragon = egg.Hatch()
	var childDragonTwo = egg.Hatch()
	fmt.Println(childDragon)
	fmt.Println(childDragonTwo)
}

func main() {
	exampleOne()
	exampleTwo()
	exampleThree()
	exampleFour()
}
