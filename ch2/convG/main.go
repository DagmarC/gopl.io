package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"gopl.io/ch2/tempconv"
)

var c = flag.Bool("c", false, "C/F converts number into Clesius/Fahrenheit and vice versa")
var m = flag.Bool("m", false, "M/F converts number into the Meters/Feets and vice versa")
var kg = flag.Bool("kg", false, "KG/LBS converts number into the Kilograms/Pounds and vice versa")
var f = flag.Bool("f", false, "C/F converts number into Clesius/Fahrenheit and vice versa")
var ft = flag.Bool("ft", false, "M/F converts number into the Meters/Feets and vice versa")
var l = flag.Bool("l", false, "KG/LBS converts number into the Kilograms/Pounds and vice versa")

func main() {
	flag.Parse()
	var numbers []float64

	for _, a := range flag.Args() {
		n, err := strconv.ParseFloat(a, 64)
		if err != nil {
			log.Fatal("wrong input")
		}
		numbers = append(numbers, n)
	}

	if len(numbers) == 0 {
		readNumbers(&numbers, "Enter the numbers separated via space.")

		readInput(c, "Temperature conversion c/f.")
		readInput(m, "Length conversion? m/f")
		readInput(kg, "Weight conversion? kg/lbs.")

		readInput(f, "Temperature conversion? f/c.")
		readInput(ft, "Length conversion? f/m.")
		readInput(l, "Weight conversion? lbs/kg.")
	}

	for _, n := range numbers {
		if *c {
			fmt.Printf("%v is %.2fF\n", tempconv.Celsius(n), converter(tempconv.Celsius(n)))
		}
		if *m {
			fmt.Printf("%v is %.2fft\n", tempconv.Meters(n), converter(tempconv.Meters(n)))
		}
		if *kg {
			fmt.Printf("%v is %.2flbs\n", tempconv.Kilograms(n), converter(tempconv.Kilograms(n)))
		}
		if *f {
			fmt.Printf("%v is %.2fC\n", tempconv.Fahrenheit(n), converter(tempconv.Fahrenheit(n)))
		}
		if *ft {
			fmt.Printf("%v is %.2fm\n", tempconv.Feets(n), converter(tempconv.Feets(n)))
		}
		if *l {
			fmt.Printf("%v is %.2fkg\n", tempconv.Pounds(n), converter(tempconv.Pounds(n)))
		}
	}
}

func readNumbers(numbers *[]float64, q string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(q)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Wrong input.")
	}
	input = strings.Replace(input, "\n", "", -1)

	for _, s := range strings.Split(input, " ") {
		n, err := strconv.ParseFloat(s, 64)
		if err != nil {
			log.Fatal("Wrong input.")
		}
		*numbers = append(*numbers, n)
	}
}

func readInput(s *bool, q string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(q + " Type yes or no.")
	converts, err := reader.ReadString('\n')
	converts = strings.Replace(converts, "\n", "", -1)

	if err != nil {
		log.Fatal("Wrong input.")
	}
	*s = strings.EqualFold(converts, "yes")
}

func converter(convert tempconv.Units) float64 {
	return convert.Converts()
}
