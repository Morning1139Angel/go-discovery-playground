package main

import "fmt"

func main() {
	//🌼 Arrays
	//=================================================
	intArr32 := [32]int{} //The type of elements and length are

	//the last elem will be the zero value
	firstIntArr5 := [5]int{1, 3, 5, 4} // both part of the array type

	fmt.Printf("intArr32: %v\n", intArr32)
	fmt.Printf("firstIntArr5: %v\n", firstIntArr5)

	//ERROR!
	//cannot use firstIntArr5 (variable of type [5]int) as [32]int value in assignment
	// intArr32 = firstIntArr5

	//ERROR!
	// cannot convert firstIntArr5 (variable of type [5]int) to type [32]int
	//intArr32 = [32]int(firstIntArr5)

	//zeor value
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
	//=================================================
}
