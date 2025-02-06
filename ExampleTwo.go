/*
File Name:  Main.go
Copyright:  2022
Author:     Emir Skopljak
*/

package main

func findMaxSum(numbers []int) int {
	var numberOne = 0
	var numberTwo = 0
	var sum = 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > numberOne && numberOne < numberTwo {
			sum -= numberOne
			numberOne = numbers[i]
			sum += numberOne
		} else if numbers[i] > numberTwo {
			sum -= numberTwo
			numberTwo = numbers[i]
			sum += numberTwo
		}
	}
	return sum
}
