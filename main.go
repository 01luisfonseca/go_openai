package main

import (
	"go_openai/tools"
	"log"
)

func main() {
	results := tools.GetCompletion("Give me a random number", "gpt-3.5-turbo")
	log.Println(results)
}
