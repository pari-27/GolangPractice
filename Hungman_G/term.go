package main

import "fmt"

type HangmanTerm struct {
}

func (h *HangmanTerm) RenderGame(placeholder []string, entries map[string]bool, chances int) error {

	fmt.Println(placeholder)                       // render the placeholder
	fmt.Printf("Chances: %d\n", chances)           // render the chances left
	fmt.Printf("Entries: %v\n", get_keys(entries)) // show the letters or words guessed till now.
	fmt.Printf("Guess a letter or the word: ")

	return nil
}

func (h *HangmanTerm) GetInput() string {
	str := ""
	fmt.Scanln(&str)

	return str
}
