package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/urfave/negroni"
)

type Hangman interface {
	RenderGame([]string, map[string]bool, int) error
	GetInput() string
}

var MAX_CHANCES int = 8
var dev = flag.Bool("dev", false, "dev mode")

func get_keys(entries map[string]bool) (keys []string) {
	for k, _ := range entries {
		keys = append(keys, k)
	}
	return
}

func get_word() string {
	if *dev {
		return "elephant"
	}
	resp, err := http.Get("https://random-word-api.herokuapp.com/word?number=5")
	if err != nil {
		return "elephant"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var words []string
	err = json.Unmarshal(body, &words)
	if err != nil {
		// handle error
		return "elephant"
	}

	//fmt.Println(words)
	for _, word := range words {
		if len(word) > 4 && len(word) < 9 {
			return word
		}
	}

	return words[0]
}

func main() {
	flag.Parse()

	// start http server when not in development mode
	go web_game()

	// start term game
	term_game()
}

func term_game() {
	word := "elephant"
	if *dev == false {
		word = get_word()
	}

	game := &HangmanTerm{}
	if play(game, word) == true {
		fmt.Println("You win! You've saved yourself from a hanging")
	} else {
		fmt.Println("Damn! You're hanged!!")
		fmt.Println("Word was: ", word)
	}
}

func web_game() {
	router := InitRouter()
	server := negroni.Classic()
	server.UseHandler(router)

	server.Run(":12345")
}

// TODO: Return an error with error-code if the game timed out
func play(h Hangman, word string) bool {
	// lookup for entries made by the user.
	entries := map[string]bool{}

	// list of "_" corrosponding to the number of letters in the word. [ _ _ _ _ _ ]
	placeholder := []string{}
	//placeholder := make([]string, len(word), len(word))

	// get length of the word len(word)
	// initialize slice with each element as "_"
	for i := 0; i < len(word); i++ {
		placeholder = append(placeholder, "_")
		//placeholder[i] = "_"
	}

	t := time.NewTimer(2 * time.Minute)
	//t := time.NewTimer(5 * time.Second)

	chances := MAX_CHANCES

	// true: win
	// false: lose
	result := make(chan bool)
	go func() {
		for {
			// evaluate a loss! If user guesses a wrong letter or the wrong word, they lose a chance.
			userInput := strings.Join(placeholder, "")
			if chances == 0 && userInput != word {
				result <- false
				return
			}

			// evaluate a win!
			if userInput == word {
				result <- true
				return
			}

			// Console display
			h.RenderGame(placeholder, entries, chances)

			// Addon validation: Allow only alpha!
			// Addon validation: manage case!

			// take the input
			str := h.GetInput()

			// if len(str) > 2, compare the word with the str
			if len(str) > 2 {
				if str == word {
					result <- true
					return
				} else {
					// you lose a chance
					entries[str] = true
					chances -= 1
					continue
				}
			}

			// compare and update entries, placeholder and chances.
			_, ok := entries[str]
			if ok {
				// key exists already; duplicate
				continue
			}

			entries[str] = true
			// check if letter exists in the word!
			found := false
			for i, v := range word {
				if str == string(v) {
					placeholder[i] = string(v)
					found = true
				}
			}

			if !found {
				chances -= 1
			}
		}
	}()

	select {
	case r := <-result:
		if r { // win
			return true

		} else { //loss
			return false
		}
	case <-t.C:
		return false
	}
}
