package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	var in int
	color.Magenta("TI-84 Emulator")
	fmt.Println("1. Graph")
	fmt.Println("2. Math")
	fmt.Println("3. Settings")
	fmt.Println("4. Exit")
	fmt.Scanln(&in)
	switch in {
	case 1:
		graphingFunction()
	case 2:
		mathFunction()
	case 3:
		var vnum = "0.1"
		color.Red("TI-84 Emulator")
		color.Yellow("Version: " + vnum)
		color.Yellow("Created by Evan Carter 2019")
	case 4:
		os.Exit(0)
	}
}
func graphingFunction() {

}
func mathFunction() {

}
