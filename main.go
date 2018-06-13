package main

import (
	"fmt"
	"reflect"
)

func checkPrime(x int) []int{
	var primeList []int
	
	for x != 1 {
		if x%2 == 0 {
			x = x/2
			primeList = append(primeList, 2)
		} else {
			primeList = append(primeList, x, 1)
			break
		}
	}
	
	return []int(primeList)
	//fmt.Println(primeList)
	
}

func main() {
	fmt.Println(checkPrime(10))
	TestPrime24()
	TestPrime48()
	TestPrime10()
}

func TestPrime24() {
	var primeListed []int = checkPrime(24)
	var expected []int = []int{2, 2, 2, 3, 1}
	if reflect.DeepEqual(primeListed, expected) {
		fmt.Println("Test case for checkPrime(24) is successful")
	} else {
		fmt.Println("Test case for checkPrime(24) is unsuccessful. Expected ", expected, " but got ", primeListed)
	}
}


func TestPrime10() {
	var primeListed []int = checkPrime(10)
	var expected []int = []int{2, 5, 1}
	if reflect.DeepEqual(primeListed, expected) {
		fmt.Println("Test case for checkPrime(10) is successful")
	} else {
		fmt.Println("Test case for checkPrime(10) is unsuccessful. Expected ", expected, " but got ", primeListed)
	}
}

func TestPrime48() {
	var primeListed []int = checkPrime(48)
	var expected []int = []int{22, 2, 2, 3, 1}
	if reflect.DeepEqual(primeListed, expected) {
		fmt.Println("Test case for checkPrime(48) is successful")
	} else {
		fmt.Println("Test case for checkPrime(48) is unsuccessful. Expected ", expected, " but got ", primeListed)
	}
}



