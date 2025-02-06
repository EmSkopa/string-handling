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
	fireDragon := FireDragon{
		Name:   "Dragon mother",
		Height: 10,
	}
	var egg = fireDragon.Lay()
	var childDragon = egg.Hatch()
	var childDragonTwo = egg.Hatch()
	if childDragon != nil {
		if babyDragon, ok := childDragon.(*FireDragon); ok {
			fmt.Println("Baby dragon's name:", babyDragon.Name)
			fmt.Println("Baby dragon's height:", babyDragon.Height)
		}
	}
	fmt.Println(childDragonTwo)
	fmt.Println(fireDragon)
}

func main() {
	exampleOne()
	exampleTwo()
	exampleThree()
	exampleFour()
}
