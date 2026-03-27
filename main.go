package main

import (
	"fmt"
	"os"
)

func main() {
	argv := os.Args
	var nb_arg int = len(argv)
    var game hangman

	if (nb_arg != 2) {
	    os.Exit(2)
	}
	game = create_game(argv[1])
	for (game.running == true) {
		print_game(game)
		request_letter(&game)
		game.running = !equalRuneSlices(game.word, game.solved) && game.lifePoints > 0
	}
	if (game.win == true) {
	    fmt.Printf("You win with %d lives left.\n", game.lifePoints)
	} else {
		fmt.Printf("You lose. The word was: %s\n", string(game.word))
	}
}