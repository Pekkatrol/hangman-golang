# Hangman Game in Golang

## Goal

The goal of this project is to create a command-line Hangman game in Golang. The script lets you play Hangman by guessing letters of a hidden word, with a limited number of attempts.

## What I used

- `fmt` for displaying output and reading user input
- `os` for handling command-line arguments and exiting on error
- `strings` for formatting output
- Slices and runes for word and letter management

### You need Golang

If you don't have Golang installed, run:
```bash
sudo apt update
sudo apt install golang-go
```

## How to use it?

Run the script with the following argument:
```bash
go run . <word>
```
- `<word>`: the word to guess (letters only, no spaces or symbols)

### Example

```bash
go run . banane
```
You should see something like:
```bash
What you currently found:
_ _ _ _ _ _
The letters you already chose: 
You still have 10 lives
Choose a letter: 
```
Then, follow the prompts to guess letters. The game ends when you find the word or run out of lives.

If you provide an invalid word (with non-letter characters), the script will print an error message and exit.

---

Enjoy playing Hangman in your terminal!
