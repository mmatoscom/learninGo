package Exercises

import "fmt"

// for is Go’s only looping construct. Here are three basic types of for loops.

func main() {

	fmt.Println("Basic type, utilize variable i=i+1 until deadline:")
	//The most basic type, with a single condition.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		if i == 3 {
			fmt.Println("\n")

		}
		i = i + 1
	}

	fmt.Println("Classic for:")
	//A classic initial/condition/after for loop.
	for j := 7; j < 9; j++ {
		fmt.Println(j)
	}

	//for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
	for {
		fmt.Println("loop...")
		break
	}

	//You can also continue to the next iteration of the loop.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
