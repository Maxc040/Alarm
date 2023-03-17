package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Controleer of er een argument is meegegeven
	if len(os.Args) == 1 {
		fmt.Println("Geef mij nummer.")
		return
	}

	// Zet het argument om naar een integer
	n, err := strconv.Atoi(os.Args[1])

	// Controleer of er een error is opgetreden tijdens de conversie of als n <= 0 is
	if err != nil || n <= 0 {
		fmt.Println("Wallah jij probeert weg te rennen.")
	}

	// Herhaal de tekst "Politie x !" n keer, waarbij x ophoogt bij elke iteratie
	for i := 1; i <= n; i++ {
		fmt.Printf("Politie %d !\n", i)
		if i >= 10 {
			fmt.Println("Wallah jij probeert weg te rennen.") //wanneer de input niet klopt, dus het is onder de 1 of boven de 10, dan krijgt de gebruiker een eroor message.
			break
		}
	}
}
