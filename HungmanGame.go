package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//iterative number constants
const (
	GAME_WIN = iota
	GAME_LOST
	TIMEOUT
	GAME_PLAY
	GAME_ERR
	GAME_HINT
)

type Hungman struct {
	Entries     map[string]bool
	Placeholder []string
	Chances     int
	Clock       *time.Timer
	word        string
	Def         string
	hint        bool
}
type WordStruct struct {
	Id   int
	Word string
}
type Defination struct {
	Text string
}

// command line input for diff ENV
var ENV = flag.String("env", "local", "mode to run the game with")

//get the words from the api
func get_words() string {
	//check if dev enviornment
	if *ENV == "dev" {
		return "elephant"
	}
	/*  for getting a word according to a category
	url := []string{"https://api.datamuse.com/words?topics=", category, "&md=d&max=1"}
	resp, err := http.Get(strings.Join(url, ""))*/

	resp, err := http.Get("https://api.wordnik.com/v4/words.json/randomWord?hasDictionaryDef=true&api_key=a2a73e7b926c924fad7001ca3111acd55af2ffabf50eb4ae5")
	//fmt.Println(resp)
	if err != nil {
		return "elephant"
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "elephant"
	}
	//fmt.Println(body)
	var words WordStruct
	err = json.Unmarshal(body, &words)
	//fmt.Println(words)
	if err != nil {
		return "elephant"
	}

	return words.Word

}

func get_defination(word string) string {
	if *ENV == "dev" {
		return "Any of several very large herbivorous mammals of the family Elephantidae native to Africa and South Asia."
	}
	url := []string{"https://api.wordnik.com/v4/word.json/", word, "/definitions?limit=2&api_key=a2a73e7b926c924fad7001ca3111acd55af2ffabf50eb4ae5"}
	resp, err := http.Get(strings.Join(url, ""))
	if err != nil {
		return "Any of several very large herbivorous mammals of the family Elephantidae native to Africa and South Asia."
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Any of several very large herbivorous mammals of the family Elephantidae native to Africa and South Asia."
	}
	var defination []Defination
	err = json.Unmarshal(body, &defination)
	if err != nil {
		return "Any of several very large herbivorous mammals of the family Elephantidae native to Africa and South Asia."
	}

	return defination[0].Text
}

//to get the keys from the entries map i.e to get a list of all the guesses made by the player
func get_keys(entries map[string]bool) (keys []string) {
	for key, _ := range entries {
		keys = append(keys, key)
	}
	return
}

//to get the user input
func getInput() (input string) {
	fmt.Scanln(&input)
	return
}

func (h *Hungman) Get_Status(Status_name uint) {
	switch Status_name {
	case GAME_WIN:
		fmt.Println("Contgratulations!!!! You won")
	case GAME_LOST:
		fmt.Println("You're out of chances")
		fmt.Println("Word was: ", h.word)
		fmt.Println("Game Over! Try again")
	case GAME_PLAY:
		fmt.Println("-----------------------")
		fmt.Println(h.Placeholder)                       // get the placeholder
		fmt.Printf("Chances: %d\n", h.Chances)           // get the remaining chances
		fmt.Printf("Entries: %v\n", get_keys(h.Entries)) // get the letters or words guessed by the user
		fmt.Printf("Guess a letter or the word or if you want a hint press *: ")

	case TIMEOUT:
		fmt.Println("Timedout... too bad!")
	case GAME_HINT:
		fmt.Println("Hint", h.Def)

	case GAME_ERR:
		fmt.Println("Something wrong with getting input")

	}

}

func play_Game(h *Hungman) (result bool) {

	for h.Chances > 0 {
		userInput := strings.Join(h.Placeholder, "")
		if userInput == h.word {
			h.Get_Status(GAME_WIN)
			return true
		}
		h.Get_Status(GAME_PLAY)
		str := getInput()
		found := false
		if str == h.word {
			h.Get_Status(GAME_WIN)
			return true
		}
		if str == "*" && h.hint == false {
			h.Chances = h.Chances - 1
			h.Get_Status(GAME_HINT)
			h.hint = true

		} else {
			fmt.Println("hint already given!!")
			continue
		}

		if !h.Entries[str] {
			for i, s := range h.word {
				if str == string(s) {
					h.Placeholder[i] = string(s)
					found = true
				}
			}
		}
		h.Entries[str] = true
		if !found {
			h.Chances = h.Chances - 1
		}
		if h.Chances == 0 {
			h.Get_Status(GAME_LOST)
			return false
		}

	}
	return true

}

/*
func get_categories() string {
	category := []string{"1 ---> animal releated", "2---> fauna related"}
	fmt.Println("Select a category to play with :\n")
	for _, v := range category {
		fmt.Println(string(v))
	}
	var catVar int
	fmt.Scanln(&catVar)
	switch catVar {
	case 1:
		return "animals"
	case 2:
		return "flowers"

	}
	return ""
}
*/

func main() {
	flag.Parse()
	h := Hungman{
		Entries:     map[string]bool{},
		Placeholder: []string{},
		Chances:     8,
		Clock:       time.NewTimer(2 * time.Minute),
	}

	//category to select the word from
	//catVar := get_categories()

	h.word = get_words()
	h.Def = get_defination(h.word)

	for i := 0; i < len(h.word); i++ {
		h.Placeholder = append(h.Placeholder, "_")
		if strings.ContainsAny("aeiou", string(h.word[i])) {
			h.Placeholder[i] = string(h.word[i])
		}
	}

	play_Game(&h)
	select {
	case <-h.Clock.C:
		h.Get_Status(TIMEOUT)
	}

}
