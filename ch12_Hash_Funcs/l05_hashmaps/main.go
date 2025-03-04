package main

import "fmt"

func main() {
	var gpas map[string]float64
    gpas = make(map[string]float64, 4)

	// don't touch below this line

	names := []string{"Lane", "Yoojin", "Jonny", "Christine"}

	gpas[names[0]] = 3.8
	gpas[names[1]] = 3.5
	gpas[names[2]] = 2.0
	gpas[names[3]] = 4.5

	for _, name := range names {
		gpa := gpas[name]
		fmt.Printf("%v has a GPA of %v\n", name, gpa)
	}
}
