package main

import "fmt"

func main() {
	colors := map[string]string{
		"red": "#FF0000",
	}
	printMap(colors)
	colors1 := make(map[string]string)
	colors1["green"] = "#green"
	printMap(colors1)
	var colors2 map[string]string
	colors2 = make(map[string]string, 20)
	colors2["yellow"] = "#yellow"
	printMap(colors2)
	delete(colors2, "yellow")
	fmt.Println(colors2)
	printMap(colors2)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("key is ", key, " and value is ", value)
	}
}
