package main

import "fmt"

func main() {
	//Insert Code here
	//var a [10]string

	// Using Slice
	//slice := make ([] type, <length>, ,<capacity>)
	slic := make ([] string, 4, 5 ) { 
		"Week 1 - $9.50",
		"Week 2 - $8.00",
		"Week 3 - $10.20",
		"Week 4 - $7.40"
	}

	for i := 0; i < len(slic); i++ {
		fmt.Printf(slic[i]+" = %d\n", len(slic[i]))
	}
}