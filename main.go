package main

import "fmt"

func main() {
	// var nameOne string = "Kevin"
	// var nameTwo = "Emma"
	// nameThree := "Mbawalla"
	// fmt.Println(nameOne, nameTwo, nameThree)

	// var numOne int = 19
	// var numTwo int8 = 127
	// var numThree int64 = -756
	// var numFour uint64 = 69420

	// a := 10
	// b := 12
	// add(a, b)
	// fmt.Println("Hello World!")

	var ages = [3]int{10, 9, 19}
	names := [4]string{"Kevin", "Luffy", "Dune", "Josh"}

	fmt.Println(ages, names)

	// Slice
	mcs := []string{"Luffy", "Ichigo", "Naruto", "Deku", "Yuji"}

	rangeOne := mcs[:3]
	rangeTwo := mcs[1:5]
	rangeThree := mcs[3:]

	bigThree := rangeOne
	rangeTwo = append(rangeTwo, "Goku")
	rangeThree = append(rangeThree, "Kamui")

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	fmt.Println(bigThree, rangeTwo, rangeThree)
}

func add(numOne int, numTwo int) {
	sum := numOne + numTwo
	fmt.Println(sum)
}