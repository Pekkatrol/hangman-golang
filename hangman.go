package main

import (
	"fmt"
	"strings"
)

type hangman struct {
	word []rune
	word_len int
	solved []rune
	lifePoints int
	letters []rune
	running bool
	win bool
}

func is_alpha(current byte) bool {
	return (current >= 'a' && current <= 'z') || (current >= 'A' && current <= 'Z')
}

func is_valid_word(word string) bool {
	for i := 0; i < len(word); i++ {
		if !is_alpha(word[i]) {
			return false
		}
	}
	return true
}

func create_game(word string) (game hangman) {
	if !is_valid_word(word) {
		fmt.Println("Invalid word")
		game.running = false
		return
	}
	game.word = []rune(word)
	game.word_len = len(game.word)
	game.solved = make([]rune, game.word_len)
	for i := range game.solved {
		game.solved[i] = '_'
	}
	game.lifePoints = 10
	game.letters = []rune{}
	game.running = true
	game.win = true
	return
}

func print_game(game hangman) {
	fmt.Print("\033[H\033[2J")
	fmt.Printf("What you currently found:\n%s\n",
	    strings.Join(runeSliceToStringSlice(game.solved), " "))
	fmt.Printf("The letters you already chose: %s\nYou still have %d lives\n",
	    string(game.letters), game.lifePoints)
}

func runeSliceToStringSlice(runes []rune) []string {
	s := make([]string, len(runes))
	for i, r := range runes {
		s[i] = string(r)
	}
	return s
}

func is_already_chose(char string, letters []rune) (bool) {
    var nb_letters int = len(letters)
	
	for index := 0; index < nb_letters; index++ {
        if (char[0] == byte(letters[index])) {
			return true
		}
	}
	return false
}

func request_letter(game *hangman) () {
	var char string
    var nb_letters int = len(game.word)
	var added bool = false

	fmt.Printf("Choose a letter: ")
	fmt.Scanf("%s\n", &char)
	if len(char) != 1 {
		fmt.Println("Please enter a single letter.")
	}
	for (is_already_chose(char, game.letters) == true) {
        fmt.Printf("Choose a different letter: ")
		fmt.Scanf("%s\n", &char)
		if len(char) != 1 {
			fmt.Println("Please enter a single letter.")
		}
	}
	letter := rune(char[0])
	for index := 0; index < nb_letters; index++ {
        if (letter == game.word[index]) {
			game.solved[index] = letter
			added = true
		}
	}
	game.letters = append(game.letters, letter)
	if (added == false) {
        game.lifePoints--
	}
	if (game.lifePoints == 0) {
		game.win = false
	}
}

func equalRuneSlices(a, b []rune) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}