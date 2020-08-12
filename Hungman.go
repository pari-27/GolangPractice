package main

import (
	"fmt"
	"strings"
)

func getKeys(keys map[string]bool) []string {
	k := []string{}

	for i, _ := range keys {
		k = append(k, i)
	}
	return k
}

func main() {
	word := "elephant"
	entries := map[string]bool{}
	placeholder := []string{}
	var userword string

	for i := 0; i < len(word); i++ {
		placeholder = append(placeholder, "_")

	}

	var chances int = len(word)
	for chances > 0 {
		userword = strings.Join(placeholder, "")
		fmt.Println(placeholder)
		fmt.Println("letters guessed:", getKeys(entries))
		fmt.Println("chances left:", chances)
		if userword == word {
			fmt.Println("Congratulations !!!! You Won ")
			break
		} else {
			userchar := ""
			fmt.Println("guess a letter or word:")
			fmt.Scanln(&userchar)
			if !entries[userchar] {
				entries[userchar] = true
				if strings.Contains(word, userchar) {
					if userchar == word {
						fmt.Println("Congratulations !!!! You Won ")
						break
					}
					for i, v := range word {
						if string(v) == userchar {
							placeholder[i] = string(v)
						}
					}
				} else {
					chances = chances - 1
					if chances == 0 {
						fmt.Println("Game Over !!")
					}
				}
			}
		}
	}

}
