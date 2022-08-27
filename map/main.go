package main

import (
	"fmt"
)

// map like objects in js , dict in python
func main() {
	// var colors map[string]string
	// colors := make(map[string]string)
	// map[type of key]type of value

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "ffffff",
	}

	// colors["white"] = "#fffffff"
	// fmt.Println(colors)
	// delete(colors, "white")
	// fmt.Println(colors)
	printMap(colors)

}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}
