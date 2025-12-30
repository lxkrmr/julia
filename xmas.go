package main

import (
	"fmt"
	"strings"
)

func main() {
	var wahl string

	for wahl != "A" {
		fmt.Println("Willst du für immer mein sein?")
		fmt.Println("A: Ja | B: Nein | C: Vielleicht")

		fmt.Print("Deine Wahl: ")
		fmt.Scanln(&wahl)

		wahl = strings.ToUpper(wahl)
		if wahl == "A" {
			fmt.Println("Yay! 🥰")
		} else {
			fmt.Println("Ungültig, aber sowas von 🙂‍↔️😒😞")
			fmt.Println("Probiere es noch einmal.")
			fmt.Println("Tipp die einzig richtige Antwort ist A")
			fmt.Println("")
		}
	}
}
