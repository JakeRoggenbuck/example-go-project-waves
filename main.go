package main

import (
	"fmt"
	"strings"
)

var SPEED_OF_LIGHT float32 = 300

func find_wavelength(freq float32) float32 {
	return SPEED_OF_LIGHT / freq
}

func find_frequency(wavelength float32) float32 {
	return SPEED_OF_LIGHT / wavelength
}

func main() {
	fmt.Println("Calculate waves!")
	fmt.Println("λ = C/f")

	fmt.Println("What do you want to try to find?")
	fmt.Print("(F) Frequency or (W) Wavelength: ")

	var option string
    fmt.Scanf("%s", &option)

	new_opt := strings.ToUpper(option)
	if new_opt == "F" {
		fmt.Print("Wavelength (M): ")

		var inwave float32
		fmt.Scanf("%f", &inwave)

		f := find_frequency(inwave)
		fmt.Println(f)
	} else if new_opt == "W" {
		fmt.Print("Frequency (MHz): ")

		var infreq float32
		fmt.Scanf("%f", &infreq)

		w := find_wavelength(infreq)
		fmt.Println(w)
	}
}
