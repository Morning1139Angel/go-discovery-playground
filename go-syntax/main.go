package main

import "fmt"

//notes added form : https://go.dev/blog/slices-intro

func main() {
	//🌼 Arrays
	//=================================================
	//An array type definition specifies a length and an element type
	intArr32 := [32]int{}              //The type of elements and length are
	firstIntArr5 := [5]int{1, 3, 5, 4} // both part of the array type
	//👆 the last elem will be the zero value

	//u can tell the compiler to infer the count as well
	lenghtInferedArr := [...]int{1, 33, 444, 98}
	println("lenght of the array 'lenghtInferedArr' is: ", len(lenghtInferedArr))

	fmt.Printf("intArr32: %v\n", intArr32)
	fmt.Printf("firstIntArr5: %v\n", firstIntArr5)

	/*An array's size is fixed;
	its length is part of its type
	([4]int and [5]int are distinct, incompatible types)*/

	//ERROR!
	//cannot use firstIntArr5 (variable of type [5]int) as [32]int value in assignment
	// intArr32 = firstIntArr5

	//ERROR!
	// cannot convert firstIntArr5 (variable of type [5]int) to type [32]int
	//intArr32 = [32]int(firstIntArr5)

	/*
		Arrays do not need to be initialized explicitly;
		the zero value of an array is a ready-to-use array
		whose elements are themselves zeroed
	*/
	var secondIntArr5 [5]int
	fmt.Printf("before assignment secondIntArr5: %v\n", secondIntArr5)

	secondIntArr5 = firstIntArr5 //valid assignment... both same type
	fmt.Printf("after assignment secondIntArr5: %v\n", secondIntArr5)

	//2d array
	var twoD [2][3]int
	fmt.Println("count of 2d array rows:", len(twoD))
	fmt.Println("count of 2d array rows:", len(twoD))

	for i := 0; i < len(twoD); i++ {
		twoD[i] = [3]int{i + 1, i + 2, i + 3}

	}
	fmt.Println("lenght of 2d array :", len(twoD))

	//array itteration with range
	for rowIndx, row := range twoD {
		fmt.Println("============================")
		fmt.Printf("Im iterating over the \"twoD\" array's row number %d\n", rowIndx)
		for columnIndx, value := range row {
			fmt.Printf("element at index [%d][%d] is: %d\n", rowIndx, columnIndx, value)
		}
	}

	/*note1:
	Go arrays are values. An array variable denotes the entire array;
	it is not a pointer to the first array element (as would be the case in C).
	This means that when you assign or pass around an array value you will make
	a copy of its contents. One way to think about arrays
	is as a sort of struct but with indexed rather than named fields:
	a fixed-size composite value.*/
	arr := [5]int{100, 100, 100, 100, 100}
	fmt.Println("value of arr before calling f1: ", arr)
	f1(arr)
	fmt.Println("value of arr after calling f1: ", arr)

	//=================================================
}

func f1(arr [5]int) {
	arr = [5]int{-1, -1, -1, -1, -1}
	fmt.Println("value of arr inside f1 call: ", arr)
}
