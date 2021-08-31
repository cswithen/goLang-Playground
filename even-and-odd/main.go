package main

import "fmt"

// my solution
func main() {

	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(n)

	for _, num := range n {
		if num%2 == 0 {
			fmt.Println(fmt.Sprint(num) + " is even")
			// the solution is fmt.Println(number, "is even")
		} else {
			fmt.Println(fmt.Sprint(num) + " is odd")
			// the solution is fmt.Println(number, "is odd")
		}
	}

}
