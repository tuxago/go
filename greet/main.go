package main

import "os"

func main() {
	var input string = os.Args[1]
	var a = greeter(input, "french")
	println(a)
}

func greeter(input string, lang string) string {
	switch lang {
	case "french":
		return "Bonjour " + input
	case "spanish":
		return "Hola " + input
	default:
		return "Hello " + input
	}

}
