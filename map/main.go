// maps can be used representing a collection of very closely related properties

package main

import (
	"fmt"
)

func main() {
	// first way to declare a map
	// create a map[key-type] value-type
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// second way to declare a map, specifically used when you would want to declare keys later
	var colors2nd map[string]string

	//third way
	colors3rd := make(map[int]string)
	colors3rd[10] = "#ffffff"
	//how to delete keys off of a map based on the key
	delete(colors3rd, 10)

	printMap(colors)
	fmt.Println(colors2nd)
	fmt.Println(colors3rd)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
