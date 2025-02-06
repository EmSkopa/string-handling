/*
File Name:  Main.go
Copyright:  2022
Author:     Emir Skopljak
*/

package main

func fillChannel(functions ...func() int) chan int {
	out := make(chan int)
	for _, v := range functions {
		v := v
		go func() {
			out <- v()
		}()
	}
	return out
}

func exampleFunction(counter int) int {
	sum := 0
	for i := 0; i < counter; i++ {
		sum += 1
	}
	return sum
}
