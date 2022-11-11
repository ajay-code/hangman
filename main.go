package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func containsAny(s string, chars string) bool {
	for _, ch := range s {
		if chars == string(ch) {
			return true
		}
	}
	return false
}

func join(strings []string, seperator string) string {
	if len(strings) == 0 {
		return ""
	}
	s := ""
	for _, v := range strings {
		s += v + seperator
	}
	return s
}

func getLetter(found []string) (string, error) {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for {
		letter, err := prompt("Pick a letter:", join(found, " "))
		if err != nil {
			panic(err)
		}
		if len(letter) == 1 && containsAny(alphabet, letter) {
			return letter, nil
		}
		fmt.Println("Invalid input: must enter a single lowercase letter")
	}
}

func prompt(vals ...interface{}) (string, error) {
	if len(vals) != 0 {
		fmt.Println(vals...)
	}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		return "", err
	}
	return scanner.Text(), nil

}

func updateFound(found []string, word string, letter string) bool {
	complete := true
	for i, r := range word {
		if letter == string(r) {
			found[i] = letter
		}
		if found[i] == "_" {
			complete = false
		}
	}
	return complete
}

func main() {
	words := []string{
		"zebra", "moose", "alligator", "elephant", "ibex", "jerboa", "cat",
	} // list of words to pick from
	rand.Seed(time.Now().UnixNano())     // seed math/rand with current time
	word := words[rand.Intn(len(words))] // get random word with the help of math/rand
	nGusses := len(word)                 // for N-letter word, player has N gusses
	found := make([]string, 0, len(word))
	for i := 0; i < nGusses; i++ {
		found = append(found, "_")
	}

	// game begins here
	for nGusses > 0 {
		fmt.Println("you have", nGusses, "remaining gusses.")
		letter, err := getLetter(found)
		if err != nil {
			panic(err)
		}
		if !containsAny(word, letter) {
			nGusses--
		}
		if updateFound(found, word, letter) {
			fmt.Println("You win! The word was:", word)
			os.Exit(0)
		}
	}
	fmt.Println("You lose! The word was:", word)
}
