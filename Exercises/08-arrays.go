package Exercises

import "fmt"

func main() {

	//Here we create an array a that will hold exactly 5 ints.
	var a[5]int
	//The type of elements and length are both part of the array’s type.
	// By default an array is zero-valued, which for ints means 0s.
	fmt.Println("emp: ",a)

	//We can set a value at an index using the array[index] =
	a[4] = 100
	fmt.Println("set:", a)
	//value syntax, and get a value with array[index].
	fmt.Println("get:", a[4])

	//The builtin len returns the length of an array.
	fmt.Println("lenght:", len(a))

	//Use this syntax to declare and initialize an array in one line.
	b := [5]int{1,2,3,4,5}
	fmt.Println("dcl:",b)


	//Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.
	var twoD[2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2 dimensions array:", twoD)

	// Note that arrays appear in the form [v1 v2 v3 ...] when printed with fmt.Println.
}

