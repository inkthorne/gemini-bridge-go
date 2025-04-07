package main

import (
	"fmt"

	greeting "github.com/inkthorne/gemini-bridge-go/bridge"
)

func main() {
	message := greeting.Greet("Gemini Bridge")
	fmt.Println(message)
}
