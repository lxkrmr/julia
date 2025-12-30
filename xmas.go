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
		_, err := fmt.Scanln(&wahl)

		if err != nil {
			fmt.Println("Irgendwie hast du einen Fehler gemacht, bitte versuche es noch einmal:", err)
		}

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
